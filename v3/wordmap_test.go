package main

import (
	"reflect"
	"testing"
)

func TestWordMap_add(t *testing.T) {
	t.Run("add key", func(t *testing.T) {
		wordMap := WordMap{}
		wordMap.add("alex", 1)

		want := WordMap{}
		want["alex"] = CountIndex{Count: 1, Sequence: 1}

		if !reflect.DeepEqual(wordMap, want) {
			t.Errorf("got %+v, but want: %+v\n", wordMap, want)
		}
	})

	t.Run("update key", func(t *testing.T) {
		wordMap := WordMap{}
		wordMap["alex"] = CountIndex{Count: 1, Sequence: 7}
		wordMap.add("alex", 9)

		want := WordMap{}
		want["alex"] = CountIndex{Count: 2, Sequence: 7}

		if !reflect.DeepEqual(wordMap, want) {
			t.Errorf("got %+v, but want: %+v\n", wordMap, want)
		}
	})
}

func TestWordMap_firstWord(t *testing.T) {
	wordMap := WordMap{}
	wordMap.add("alex", 1)
	wordMap.add("chi", 2)
	wordMap.add("alex", 3)
	wordMap.add("cai", 4)
	wordMap.add("cai", 5)
	wordMap.add("ruo", 6)
	targetWord := wordMap.firstWord(6)
	want := Dictionary{
		Word:       "chi",
		CountIndex: CountIndex{Count: 1, Sequence: 2},
	}

	if !reflect.DeepEqual(targetWord, want) {
		t.Errorf("got %+v, but want: %+v\n", targetWord, want)
	}
}
