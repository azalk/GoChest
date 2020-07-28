package GoChest

type Changepoint struct {
	segment []int
	boundaries [][]int
	exactPosition int
}

func (changepoint Changepoint) getBoundary(indexOffset int) int {
	return changepoint.boundaries[changepoint.segment[0]][changepoint.segment[1] + indexOffset]
}

func (changepoint Changepoint) getSegmentLength() int {
	return changepoint.getBoundary(1) - changepoint.getBoundary(0)
}

func (changepoint Changepoint) getLeftMidpointRight(index int, digitCount int) (int, int, int) {
	leftBoundary := changepoint.getBoundary(-1) * digitCount
	rightBoundary := changepoint.getBoundary(2) * digitCount
	midpoint := leftBoundary + (index * digitCount)

	return leftBoundary, midpoint, rightBoundary
}

func (changepoint *Changepoint) findExactChangepoint(scoresArr [][]float64) {
	maxScore := 0.0
	for i, scores := range scoresArr {
		score := 0.0
		for _, subscore := range scores {
			score += subscore
		}

		if score > maxScore {
			maxScore = score
			(*changepoint).exactPosition = changepoint.getBoundary(0) + i
		}
	}
}
