package main

import (
	"os"
	"testing"
)

func TestDecoder(t *testing.T) {
	wordMap := WordMap{}
	wordMap.add("alex", 1)
	wordMap.add("chi", 2)
	wordMap.add("cai", 3)
	wordMap.add("ruo", 4)

	encoder := newEncoder(2)
	if err := encoder.SaveWordsMap(&wordMap); err != nil {
		t.Error(err)
	}
	if err := encoder.flush(); err != nil {
		t.Error(err)
	}
}

func Test_newDecoder(t *testing.T) {
	os.Create("f0")
	decoder := newDecoder(1)
	if decoderCount := len(decoder.members); decoderCount != 1 {
		t.Errorf("unexpected decode worker count, got %d, want %d.\n", decoderCount, 1)
	}
	os.Remove("f0")
}
