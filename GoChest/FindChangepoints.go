package GoChest

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// The matrix containing all the distances between the segments
// Distances will be computed as needed and then cached for later use
var distanceMap = make(map[string]float64)

var changepoints []int
var sequenceLength int

func getLeftRightOfSegment(segment int) [2]int {
	left := 0
	if segment != 0 {
		left = changepoints[segment-1]
	}

	right := sequenceLength
	if segment != len(changepoints) {
		right = changepoints[segment]
	}

	return [2]int{left, right}
}

func getKeyString(segment1, segment2 int) string {
	if segment1 < segment2 {
		return strconv.Itoa(segment1) + "/" + strconv.Itoa(segment2)
	}
	return strconv.Itoa(segment2) + "/" + strconv.Itoa(segment1)
}

func getDistance(segment1, segment2 int) float64 {
	if _, ok := distanceMap[getKeyString(segment1, segment2)]; !ok {
		distance := 0.0

		// The distance is not cached, we need to compute it
		for level := 0; level < discreteLevel; level++ {
			distance += discreteDistance(level, getLeftRightOfSegment(segment1), getLeftRightOfSegment(segment2))
		}

		// Add the computed distance to the map
		distanceMap[getKeyString(segment1, segment2)] = distance
		return distance
	}

	return distanceMap[getKeyString(segment1, segment2)]

}

func FindChangepoints(sequence []float64, minimumDistance float64, processCount int, wordLength int) []int {
	fmt.Println("Debug Version!")
	sequenceLength = len(sequence) / wordLength

	changepoints = ListEstimator(sequence, minimumDistance, wordLength)
	sort.Ints(changepoints)

	if processCount > len(changepoints) {
		return changepoints
	}

	clusterCenter := make([]int, 1)
	// While the next line is redundant because of the 0 initialisation of Go,
	// it is here to emphasize that the clustering algorithm always starts with the first segment as first cluster
	clusterCenter[0] = 0

	for len(clusterCenter) < processCount {
		maxDistance := 0.0
		index := 0

		// We iterate over all segments
		for i := 0; i < len(changepoints)+1; i++ {

			// We look for the smallest distance of that segment to any cluster center
			minDistance := math.MaxFloat64
			for _, centerPosition := range clusterCenter {
				distance := getDistance(i, centerPosition)

				if distance < minDistance {
					minDistance = distance
				}
			}

			// If the minimal distance to all existing cluster centers is bigger than our current biggest distance, replace it
			if minDistance > maxDistance {
				maxDistance = minDistance
				index = i
			}

		}

		// Add the one with the biggest distance to the existing centers to our cluster centers
		clusterCenter = append(clusterCenter, index)
	}

	// Determining the cluster each segment belongs to
	clusters := make([]int, len(changepoints)+1)
	for i := 0; i < len(changepoints)+1; i++ {
		minDistance := math.MaxFloat64

		for _, centerPosition := range clusterCenter {
			distance := getDistance(i, centerPosition)

			if distance < minDistance {
				minDistance = distance
				clusters[i] = centerPosition
			}
		}
	}

	// Only adding those to the output where the process generating them differs
	output := make([]int, 0)
	for i := 1; i < len(changepoints)+1; i++ {
		if clusters[i-1] != clusters[i] {
			output = append(output, changepoints[i-1])
		}
	}

	return output
}
