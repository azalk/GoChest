package GoChest

import (
	"encoding/binary"
	"math"
	"sort"
)
import "github.com/BobuSumisu/aho-corasick"

// A couple words on the language used in the comments of this code:
// the input sequence can consist of a number of vectors (each vector can have length 1 in the non multivariate case)
// the input sequence then consists of the vectors written directly after each other (as opposed to a 2D array)
// wordLength is equal to the length of the vectors and is being set in the ListEstimator function
// When saying "element" I will refer to a single vector, there are len(sequence) / wordLength elements in any sequence
// When saying "character" I will refer to one parameter of a vector, there are len(sequence) different characters in the sequence
// (and wordLength different characters make up a vector)
// only in the non multivariate case does a single character equal an element which is the same as a vector
//
// Finally during discretization each character which is a float by default, might be represented by multiple different bytes (digits)
// The digitCount variable below shows how many digits make up a single element
// Thus one element is represented by digitCount different digits in the discreteSequence
// and an character is represented by digitCount / wordLength different digits in the discreteSequence
const maxDiscreteLevel = 5
var wordLength = 1

var discreteLevel int
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
	// The boundaries give the index of the element that starts and ends each bound
	// not the character/digit in the sequence/discreteSequence respectively

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
	// Since we dont really care about the underlying element/vector structure when discretizing
	// we sort the sequence by characters, destroying that underlying structure when determining the discreteLevel

	sortedSequence := make([]float64, len(sequence))
	copy(sortedSequence, sequence)
	sort.Float64s(sortedSequence)

	smallestDifference := math.MaxFloat64

	for i := range sortedSequence {
		if i != len(sortedSequence)-1 {
			difference := sortedSequence[i+1] - sortedSequence[i]
			if difference != 0 && difference < smallestDifference {
				smallestDifference = difference
			}
		}
	}

	return max(-int(math.Log2(smallestDifference)), 1)
}

func discretizeSequence(sequence []float64, level int) ([]byte, int) {
	// we discretize the sequence by characters, not by elements

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
	// How many digits each character of the original sequence gets split up into
	digitsPerCharacter := int(logBaseByte) + 1

	output := make([]byte, digitsPerCharacter*len(sequence))
	byteSequence := make([]byte, 8)
	for i := range scaledSequence {
		discreteValue := sort.SearchInts(bins, int(scaledSequence[i]))
		binary.LittleEndian.PutUint32(byteSequence, uint32(discreteValue))
		copy(output[i*digitsPerCharacter:(i+1)*digitsPerCharacter], byteSequence[:digitsPerCharacter+1])
	}

	digitCount := digitsPerCharacter * wordLength
	return output, digitCount
}

func buildTrie(discreteSequence []byte, digitCount int) *ahocorasick.Trie {
	trieBuilder := ahocorasick.NewTrieBuilder()

	// we divide the discreteSequence we are getting by the digitCount to get the amount of elements
	maxPatternLen := int(math.Log(float64(len(discreteSequence)) / float64(digitCount)))

	elementsInSequence := len(discreteSequence)/digitCount

	for m := 1; m < min(elementsInSequence, maxPatternLen)+1; m++ {
		for j := 0; j < elementsInSequence - (m + 1); j += 1 {
			trieBuilder.AddPattern(discreteSequence[j * digitCount : (j + m) * digitCount])
		}
	}

	return trieBuilder.Build()
}

// gets the distance between two segments of the discreteSequence
// the segment indices are element indices
func discreteDistance(level int, segment1, segment2 [2]int, trie *ahocorasick.Trie, discreteSequence []byte, digitCount int) float64 {

	matches := make([][]*ahocorasick.Match, 2)

	matches[0] = trie.Match(discreteSequence[segment1[0] * digitCount:segment1[1] * digitCount])
	matches[1] = trie.Match(discreteSequence[segment2[0] * digitCount:segment2[1] * digitCount])

	counts := make(map[string][]int)
	for t := 0; t < 2; t++ {
		for _, match := range matches[t] {
			if _, ok := counts[match.MatchString()]; !ok {
				counts[match.MatchString()] = make([]int, 2)
			}
			counts[match.MatchString()][t] += 1
		}
	}

	// the sequenceLength is in elements
	sequenceLength := []int{segment1[1] - segment1[0], segment2[1] - segment2[0]}

	output := 0.0
	frequencies := make([]float64, 2)
	for key, count := range counts {
		keywordLength := len(key) / digitCount

		for i := 0; i < 2; i++ {
			frequencies[i] = float64(count[i]) / float64(sequenceLength[i]-keywordLength+1)
		}

		adjustedCountDifference := math.Abs(frequencies[0] - frequencies[1])
		output += adjustedCountDifference * math.Pow(2, -float64(keywordLength))
	}

	return output * math.Pow(2, -float64(level+1))
}

func discreteDistanceMidpoint(level, leftBoundary, midPoint, rightBoundary int, trie *ahocorasick.Trie, discreteSequence []byte, digitCount int) float64 {
	return discreteDistance(level, [2]int{leftBoundary, midPoint}, [2]int{midPoint, rightBoundary}, trie, discreteSequence, digitCount)
}
