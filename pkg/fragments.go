package dictionary

import (
	"math"
)

type fragment struct {
	source interface{} // arbitrary value to associate with this fragment
	wordMap map[string]int // key<word> value<position in tf map>
	tf map[int]float64 // key<vector index> value<qty of term occurance>
	magnitude float64 // sum of each tf value ^2
}

func makeFragment(words []string, source interface{}) fragment {
	keys := 0
	dict := fragment {
		source,
		make(map[string]int),
		make(map[int]float64),
		0,
	}

	for _, word := range words {
		_, ok := dict.wordMap[word]
		if !ok {
			dict.wordMap[word] = keys
			keys += 1
		}

		dict.tf[dict.wordMap[word]] += 1
	}

	for _, val := range dict.tf {
		dict.magnitude += val * val
	}

	return dict
}

// compute cosine similiartiy between the new word vector and the
// original fragment.
func (savedFragment *fragment) ComputeSimilarity(sampleFragment *fragment) float64 {
	var dotproduct float64 = 0
	
	for word, sampleTfIndex := range sampleFragment.wordMap {
		tfIndex, ok := savedFragment.wordMap[word]
		if ok {
			dotproduct += sampleFragment.tf[sampleTfIndex] / savedFragment.tf[tfIndex]
		}
	}

	// Return similiarty
	return dotproduct / (math.Sqrt(savedFragment.magnitude) * math.Sqrt(sampleFragment.magnitude))
}
