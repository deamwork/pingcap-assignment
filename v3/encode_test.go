package main

import (
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	wordMap := WordMap{}
	wordMap.add("alex", 1)
	wordMap.add("chi", 2)
	wordMap.add("cai", 3)
	wordMap.add("ruo", 4)

	encoder := newEncoder(2)
	t.Run("save map", func(t *testing.T) {
		if err := encoder.SaveWordsMap(&wordMap); err != nil {
			t.Error(err)
		}
	})

	t.Run("flush & close", func(t *testing.T) {
		if err := encoder.flush(); err != nil {
			t.Error(err)
		}

		if err := encoder.closeAll(); err != nil {
			t.Error(err)
		}
	})
}

func Test_newEncoder(t *testing.T) {
	os.Create("foo.bar")
	newEncoder(1)
	if _, err := os.Stat("foo.bar"); os.IsNotExist(err) {
		t.Error("cannot create temp file")
	}
}
