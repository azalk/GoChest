package GoChest

import (
	"encoding/binary"
	"math"
	"sort"
)
import "github.com/BobuSumisu/aho-corasick"

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

func discreteDistance(level int, segment1, segment2 [2]int) float64 {
	matches := make([][]*ahocorasick.Match, 2)
	matches[0] = trie[level].Match(discreteSequence[level][segment1[0]:segment1[1]])
	matches[1] = trie[level].Match(discreteSequence[level][segment2[0]:segment2[1]])

	counts := make(map[string][]int)
	for t := 0; t < 2; t++ {
		for _, match := range matches[t] {
			if _, ok := counts[match.MatchString()]; !ok {
				counts[match.MatchString()] = make([]int, 2)
			}
			counts[match.MatchString()][t] += 1
		}
	}

	sequenceLength := []int{(segment1[1] - segment1[0]) / digitCount[level], (segment2[1] - segment2[0]) / digitCount[level]}

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

func discreteDistanceMidpoint(level, leftBoundary, midPoint, rightBoundary int) float64 {
	return discreteDistance(level, [2]int{leftBoundary, midPoint}, [2]int{midPoint, rightBoundary})
}
