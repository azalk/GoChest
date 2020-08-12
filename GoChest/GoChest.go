package GoChest

import (
	"encoding/binary"
	"fmt"
	"math"
	"sort"
	"sync"
)
import "github.com/BobuSumisu/aho-corasick"
import "github.com/cheggaaa/pb"

const maxDiscreteLevel = 5

var discreteLevel int

var trie []*ahocorasick.Trie
var discreteSequence [][]byte
var digitCount []int
var boundaries [][]int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getBoundaries(sequenceLength int, alpha float64, t int) []int {
	output := make([]int, 0)

	nextBound := math.Pow(float64(t+2), -1) * float64(sequenceLength) * alpha
	for true {
		output = append(output, int(nextBound))
		nextBound += float64(sequenceLength) * alpha
		if int(nextBound) >= sequenceLength {
			break
		}
	}

	return output
}

func getDiscreteLevel(sequence []float64) int {
	deltas := make([]float64, len(sequence))
	copy(deltas, sequence)
	sort.Float64s(deltas)

	minDif := math.MaxFloat64

	for i := range deltas {
		if i != len(deltas)-1 {
			dif := deltas[i+1] - deltas[i]
			if dif != 0 && dif < minDif {
				minDif = dif
			}
		}
	}

	return max(-int(math.Log2(minDif)), 1)
}

func discretizeSequence(sequence []float64, level int) []byte {
	binMap := make(map[int]struct{})
	scaledSequence := make([]float64, len(sequence))
	scalar := math.Pow(2, float64(level+1))
	for i := range sequence {
		scaledSequence[i] = sequence[i] * scalar
		binMap[int(scaledSequence[i])] = struct{}{}
	}

	bins := make([]int, 0)
	for key := range binMap {
		bins = append(bins, key)
	}
	sort.Ints(bins)

	logBaseByte := math.Log(float64(len(bins))) / math.Log(256)
	digitCount[level] = int(logBaseByte) + 1

	output := make([]byte, digitCount[level]*len(sequence))
	byteSequence := make([]byte, 8)
	for i := range scaledSequence {
		discreteValue := sort.SearchInts(bins, int(scaledSequence[i]))
		binary.LittleEndian.PutUint32(byteSequence, uint32(discreteValue))
		copy(output[i*digitCount[level]:(i+1)*digitCount[level]], byteSequence[:digitCount[level]+1])
	}

	return output
}

func buildTrie(sequence []byte, level int) *ahocorasick.Trie {
	trieBuilder := ahocorasick.NewTrieBuilder()

	maxPatternLen := int(math.Log(float64(len(sequence)) / float64(digitCount[level])))

	for m := 1; m < min(len(sequence)/digitCount[level], maxPatternLen)+1; m++ {
		for j := 0; j < len(sequence)-m*digitCount[level]+1; j += digitCount[level] {
			trieBuilder.AddPattern(sequence[j : j+m*digitCount[level]])
		}
	}

	return trieBuilder.Build()
}

func discreteDistance(level, leftBoundary, midPoint, rightBoundary int) float64 {
	matches := make([][]*ahocorasick.Match, 2)
	matches[0] = trie[level].Match(discreteSequence[level][leftBoundary:midPoint])
	matches[1] = trie[level].Match(discreteSequence[level][midPoint:rightBoundary])

	counts := make(map[string][]int)
	for t := 0; t < 2; t++ {
		for _, match := range matches[t] {
			if _, ok := counts[match.MatchString()]; !ok {
				counts[match.MatchString()] = make([]int, 2)
			}
			counts[match.MatchString()][t] += 1
		}
	}

	sequenceLength := []int{midPoint - leftBoundary, rightBoundary - midPoint}

	output := 0.0
	frequencies := make([]float64, 2)
	for key, count := range counts {
		keywordLength := len(key) / digitCount[level]

		for i := 0; i < 2; i++ {
			frequencies[i] = float64(count[i]) / float64(sequenceLength[i]-keywordLength+1)
		}

		adjustedCountDifference := math.Abs(frequencies[0] - frequencies[1])
		output += adjustedCountDifference * math.Pow(2, -float64(keywordLength))
	}

	return output * math.Pow(2, -float64(level+1))
}

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

					scores[t][level][boundary-1] = discreteDistance(level, leftBoundary, midpoint, rightBoundary)
					bar.Increment()
				}(t, level, boundary)
			}
		}
	}
	bar.Set(segmentCount)
	bar.Finish()

	output := make([][]float64, 2)
	output[0] = make([]float64, len(boundaries[0])-1)
	output[1] = make([]float64, len(boundaries[1])-1)
	waitGroup.Wait()

	for t := 0; t < 2; t++ {
		for level := 0; level < discreteLevel; level++ {
			for boundary := 0; boundary < len(boundaries[t])-3; boundary++ {
				output[t][boundary+1] += scores[t][level][boundary]
			}
		}
	}

	return output
}

func FindChangePoints(sequence []float64, minimumDistance float64) []int {
	discreteLevel = min(getDiscreteLevel(sequence), maxDiscreteLevel)

	trie = make([]*ahocorasick.Trie, discreteLevel)
	digitCount = make([]int, discreteLevel)
	discreteSequence = make([][]byte, discreteLevel)
	boundaries = make([][]int, 2)

	boundaries[0] = getBoundaries(len(sequence), minimumDistance/3.0, 0)
	boundaries[1] = getBoundaries(len(sequence), minimumDistance/3.0, 1)

	bar := pb.New(discreteLevel).Prefix("Generating Tries: ")
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
	bar.Set(discreteLevel)
	bar.Finish()
	waitGroup.Wait()

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
	waitGroup.Add(segmentLength * discreteLevel * len(changepoints))

	exactChangepointsScores := make([][][]float64, len(changepoints))
	for i, changepoint := range changepoints {
		exactChangepointsScores[i] = make([][]float64, segmentLength)

		for j := 0; j < segmentLength; j++ {
			exactChangepointsScores[i][j] = make([]float64, discreteLevel)

			for level := 0; level < discreteLevel; level++ {
				leftBoundary, midpoint, rightBoundary := changepoint.getLeftMidpointRight(j, digitCount[level])
				go func(i, j, level, leftBoundary, midpoint, rightBoundary int) {
					defer waitGroup.Done()
					exactChangepointsScores[i][j][level] = discreteDistance(level, leftBoundary, midpoint, rightBoundary)
					bar.Increment()
				}(i, j, level, leftBoundary, midpoint, rightBoundary)
			}
		}
	}
	bar.Set(segmentLength * discreteLevel * len(changepoints))
	bar.Finish()
	waitGroup.Wait()

	for i := range changepoints {
		(&changepoints[i]).findExactChangepoint(exactChangepointsScores[i])
	}

	output := make([]int, 1)
	output[0] = changepoints[0].exactPosition

	for _, chpt := range changepoints {
		fmt.Println(chpt.segment, chpt.exactPosition)
	}

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
