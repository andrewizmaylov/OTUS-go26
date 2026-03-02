package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func Top10(s string) []string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")
	words := strings.Split(s, " ")

	freq := make(map[string]int)
	for _, word := range words {
		if utf8.RuneCountInString(word) > 0 {
			freq[word]++
		}
	}

	type WordCount struct {
		Word  string
		Count int
	}

	wordCounts := make([]WordCount, 0, len(freq))
	for word, count := range freq {
		wordCounts = append(wordCounts, WordCount{Word: word, Count: count})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		if wordCounts[i].Count != wordCounts[j].Count {
			return wordCounts[i].Count > wordCounts[j].Count
		}

		return wordCounts[i].Word < wordCounts[j].Word
	})

	for i, wc := range wordCounts {
		if i > 9 {
			break
		}
		fmt.Printf("word: %#v, count: %#v\n", wc.Word, wc.Count)
	}

	output := make([]string, 0, len(freq))

	for i, wc := range wordCounts {
		if i >= 10 {
			break
		}
		output = append(output, wc.Word)
	}

	return output
}
