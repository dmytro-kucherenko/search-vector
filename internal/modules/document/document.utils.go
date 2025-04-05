package document

import (
	"errors"
	"maps"
	"math"
	"search-vector/internal/common/utils"
	"slices"
	"strings"
)

type Vector = map[string]float64

func GetVectorIDF(contents []string) Vector {
	vector := make(Vector)

	for _, content := range contents {
		for _, word := range strings.Fields(content) {
			_, ok := vector[word]
			if !ok {
				vector[word] = 1
			}
		}
	}

	return vector
}

func GetVectorTF(words []string, terms []string) Vector {
	counts := make(map[string]uint)
	for _, word := range words {
		counts[word]++
	}

	max := float64(slices.Max(slices.Collect(maps.Values(counts))))
	vector := make(Vector)

	for _, term := range terms {
		vector[term] = 0.5 + 0.5*float64(counts[term])/max
	}

	return vector
}

func GetVector(query string, vectorIDF Vector) (Vector, error) {
	vector := make(Vector)
	words := strings.Fields(query)

	for _, word := range words {
		_, ok := vectorIDF[word]
		if !ok {
			return nil, errors.New("query contains invalid term")
		}
	}

	vectorTF := GetVectorTF(words, slices.Collect(maps.Keys(vectorIDF)))
	for term, idf := range vectorIDF {
		vector[term] = vectorTF[term] * idf
	}

	return vector, nil
}

func CompareVectors(left Vector, right Vector) float64 {
	keys := utils.Union(slices.Collect(maps.Keys(left)), slices.Collect(maps.Keys(right)))

	var multiply float64
	var leftSquare float64
	var rightSquare float64

	for _, key := range keys {
		leftValue := left[key]
		rightValue := right[key]

		multiply += leftValue * rightValue
		leftSquare += math.Pow(leftValue, 2)
		rightSquare += math.Pow(rightValue, 2)
	}

	return multiply / (math.Sqrt(leftSquare) * math.Sqrt(rightSquare))
}
