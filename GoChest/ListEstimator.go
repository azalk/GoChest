package GoChest

import (
	ahocorasick "github.com/BobuSumisu/aho-corasick"
	"github.com/cheggaaa/pb"
	"math"
	"sync"
)

func getSegmentScores() [][]float64 {
	segmentCount := discreteLevel*(len(boundaries[0])-3) + discreteLevel*(len(boundaries[1])-3)
	scores := make([][][]float64, 2)
	for t := 0; t < 2; t++ {
		scores[t] = make([][]float64, discreteLevel)
		for level := 0; level < discreteLevel; level++ {
			scores[t][level] = make([]float64, len(boundaries[t])-3)
		}
	}

	bar := pb.New(segmentCount).Prefix("Computing Segment Scores: ")
	bar.AlwaysUpdate = true
	var waitGroup sync.WaitGroup
	waitGroup.Add(segmentCount)
	for t := 0; t < 2; t++ {
		for level := 0; level < discreteLevel; level++ {
			for boundary := 1; boundary < len(boundaries[t])-2; boundary++ {
				func(t, level, boundary int) {
					defer waitGroup.Done()
					leftBoundary := boundaries[t][boundary] * digitCount[level]
					rightBoundary := boundaries[t][boundary+1] * digitCount[level]
					midpoint := leftBoundary + int(float64(boundaries[t][boundary+1]-boundaries[t][boundary])*0.5)*digitCount[level]

					scores[t][level][boundary-1] = discreteDistanceMidpoint(level, leftBoundary, midpoint, rightBoundary)
					bar.Increment()
				}(t, level, boundary)
			}
		}
	}

	output := make([][]float64, 2)
	output[0] = make([]float64, len(boundaries[0])-1)
	output[1] = make([]float64, len(boundaries[1])-1)

	waitGroup.Wait()
	bar.Set(segmentCount)
	bar.Finish()

	for t := 0; t < 2; t++ {
		for level := 0; level < discreteLevel; level++ {
			for boundary := 0; boundary < len(boundaries[t])-3; boundary++ {
				output[t][boundary+1] += scores[t][level][boundary]
			}
		}
	}

	return output
}

func ListEstimator(sequence []float64, minimumDistance float64) []int {
	discreteLevel = min(getDiscreteLevel(sequence), maxDiscreteLevel)

	trie = make([]*ahocorasick.Trie, discreteLevel)
	digitCount = make([]int, discreteLevel)
	discreteSequence = make([][]byte, discreteLevel)
	boundaries = make([][]int, 2)

	boundaries[0] = getBoundaries(len(sequence), minimumDistance/3.0, 0)
	boundaries[1] = getBoundaries(len(sequence), minimumDistance/3.0, 1)

	bar := pb.New(discreteLevel).Prefix("Generating Tries: ")
	bar.AlwaysUpdate = true
	var waitGroup sync.WaitGroup
	waitGroup.Add(discreteLevel)
	for level := 0; level < discreteLevel; level++ {
		go func(level int) {
			defer waitGroup.Done()
			discreteSequence[level] = discretizeSequence(sequence, level)
			trie[level] = buildTrie(discreteSequence[level], level)
			bar.Increment()
		}(level)
	}
	waitGroup.Wait()
	bar.Set(discreteLevel)
	bar.Finish()

	segmentScores := getSegmentScores()
	changepoints := make([]Changepoint, 0)

	for true {
		maxScore := 0.0
		index := make([]int, 2)
		for t := 0; t < 2; t++ {
			for i, score := range segmentScores[t] {
				if score > maxScore {
					maxScore = score
					index = []int{t, i}
				}
			}
		}

		if maxScore == 0.0 {
			break
		}

		changepoints = append(changepoints, Changepoint{segment: index, boundaries: boundaries})

		// Wherever the Changepoint is in this segment, it cannot be in the two segment immediately left or right to it so we dont consider those
		// This follows as every segment is minimumDistance/3 long
		for offset := -2; offset < 3; offset++ {
			if index[1]+offset >= 0 && index[1]+offset < len(boundaries[index[0]])-1 {
				segmentScores[index[0]][index[1]+offset] = 0
			}
		}
	}

	segmentLength := changepoints[0].getSegmentLength()
	bar = pb.New(segmentLength * discreteLevel * len(changepoints)).Prefix("Finding Exact Changepoints: ")
	bar.AlwaysUpdate = true

	exactChangepointsScores := make([][]float64, segmentLength)
	for changepointIndex, changepoint := range changepoints {
		waitGroup.Add(segmentLength * discreteLevel)

		for i := 0; i < segmentLength; i++ {
			exactChangepointsScores[i] = make([]float64, discreteLevel)

			for level := 0; level < discreteLevel; level++ {

				leftBoundary, midpoint, rightBoundary := changepoint.getLeftMidpointRight(i, digitCount[level])

				go func(i, level, leftBoundary, midpoint, rightBoundary int) {
					defer waitGroup.Done()
					exactChangepointsScores[i][level] = discreteDistanceMidpoint(level, leftBoundary, midpoint, rightBoundary)
					bar.Increment()
				}(i, level, leftBoundary, midpoint, rightBoundary)
			}
		}

		waitGroup.Wait()
		(&changepoints[changepointIndex]).findExactChangepoint(exactChangepointsScores)
	}

	bar.Set(segmentLength * discreteLevel * len(changepoints))
	bar.Finish()

	output := make([]int, 1)
	output[0] = changepoints[0].exactPosition

	for _, changepoint := range changepoints {
		tooClose := false
		for _, alreadyPresent := range output {
			if math.Abs(float64(changepoint.exactPosition-alreadyPresent)) < (minimumDistance * float64(len(sequence))) {
				tooClose = true
				break
			}
		}

		if !tooClose {
			output = append(output, changepoint.exactPosition)
		}
	}

	return output
}
