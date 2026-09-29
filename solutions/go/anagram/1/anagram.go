package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {

	anagrams := make([]string, 0)

	for _, v := range candidates {

		if len(subject) != len(v) || strings.EqualFold(v, subject) {
			continue
		}

		isAnagram := true

		for _, c := range subject {

			cCount := strings.Count(strings.ToLower(subject), strings.ToLower(string(c)))
			vCount := strings.Count(strings.ToLower(v), strings.ToLower(string(c)))

			if cCount != vCount {
				isAnagram = false
				break
			}

			if !strings.Contains(strings.ToLower(v), strings.ToLower(string(c))) {
				isAnagram = false
				break
			}
		}

		if isAnagram {
			anagrams = append(anagrams, v)
		}

	}

	return anagrams
}
