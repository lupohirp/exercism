package etl

import (
	"maps"
	"slices"
	"strings"
)

func Transform(in map[int][]string) map[string]int {

	transfomedLexiconia := map[string]int{}
	sortedTransformedLexiconia := map[string]int{}

	for k, v := range in {

		for _, vIn := range v {

			transfomedLexiconia[strings.ToLower(vIn)] = k
		}

	}

	theSortedSliceOfKeys := slices.Sorted(maps.Keys(transfomedLexiconia))

	for _, v := range theSortedSliceOfKeys {
		sortedTransformedLexiconia[v] = transfomedLexiconia[v]
	}

	return sortedTransformedLexiconia

}
