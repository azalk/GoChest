package GoChest

type Changepoint struct {
	segment       []int
	boundaries    [][]int
	exactPosition int
}

// indexOffset here is an offset of boundaries which are element indexed
func (changepoint Changepoint) getBoundary(indexOffset int) int {
	return changepoint.boundaries[changepoint.segment[0]][changepoint.segment[1]+indexOffset]
}

// gets the segment length in elements
func (changepoint Changepoint) getSegmentLength() int {
	return changepoint.getBoundary(1) - changepoint.getBoundary(0)
}

// returns the left and right boundaries as well as the midpoint as element indices
func (changepoint Changepoint) getLeftMidpointRight(index int) (int, int, int) {
	leftBoundary := changepoint.getBoundary(-1)
	rightBoundary := changepoint.getBoundary(2)
	midpoint := changepoint.getBoundary(0) + index

	return leftBoundary, midpoint, rightBoundary
}

func (changepoint *Changepoint) findExactChangepoint(scoresArr [][]float64) {
	maxScore := 0.0
	for i, scores := range scoresArr {
		score := 0.0
		for _, subScore := range scores {
			score += subScore
		}

		if score > maxScore {
			maxScore = score
			(*changepoint).exactPosition = changepoint.getBoundary(0) + i
		}
	}
}
