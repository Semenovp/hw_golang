package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type sortItem struct {
	key   string
	value int
}

var re = regexp.MustCompile(`[.,()!?^&*'";:/\\|-]+`)

func Top10(text string) []string {
	data := strings.Fields(text)

	if len(data) == 0 {
		return []string{}
	}

	wordFrequency := make(map[string]int)
	for _, word := range data {
		// Remove unwanted characters like .,()- from the word
		word = re.ReplaceAllString(word, "")
		if len(word) == 0 {
			continue
		}
		wordFrequency[strings.ToLower(word)]++
	}

	items := make([]sortItem, 0, len(wordFrequency))
	for key, value := range wordFrequency {
		items = append(items, sortItem{key, value})
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].value == items[j].value {
			return items[i].key < items[j].key
		}
		return items[i].value > items[j].value
	})

	var result []string

	lenSize := 10
	if len(items) < 10 {
		lenSize = len(items)
	}

	for i := 0; i < lenSize; i++ {
		result = append(result, items[i].key)
	}

	return result
}
