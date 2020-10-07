package main

import "strings"

// Dictionary contain a certain word and it's occurrences
type Dictionary struct {
	Word string
	CountIndex
}

// CountIndex: total occurrences and sequence of the word
type CountIndex struct {
	Count    int
	Sequence int
}

// WordsMap: map of word occurrences
type WordMap map[string]CountIndex

// add: word to wordmap instance
func (wm *WordMap) add(word string, seq int) {
	word = strings.ToLower(word)
	if counterIndex, found := (*wm)[word]; found {
		(*wm)[word] = CountIndex{counterIndex.Count + 1, counterIndex.Sequence}
	} else {
		(*wm)[word] = CountIndex{1, seq}
	}
}

// firstWord: return a Dictionary with minimum sequence number
func (wm *WordMap) firstWord(sequenceTotal int) Dictionary {
	// init
	firstWord := Dictionary{"", CountIndex{0, sequenceTotal}}

	for word, counterIndex := range *wm {
		if counterIndex.Count == 1 && counterIndex.Sequence < firstWord.Sequence {
			firstWord.Word = word
			firstWord.Count = 1
			firstWord.Sequence = counterIndex.Sequence
		}
	}

	return firstWord
}
