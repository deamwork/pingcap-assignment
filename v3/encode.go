package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Encoder struct {
	files   []*os.File
	writers []*bufio.Writer
	members []*gob.Encoder
	slice   int
}

func newEncoder(slice int) *Encoder {
	// Initialize all required resources
	files := make([]*os.File, slice)
	writers := make([]*bufio.Writer, slice)
	members := make([]*gob.Encoder, slice)
	for i := 0; i < slice; i++ {
		f, err := os.Create(fmt.Sprintf("f%d", i))
		if err != nil {
			log.Fatal(err)
		}
		files[i] = f
		writers[i] = bufio.NewWriter(f)
		members[i] = gob.NewEncoder(writers[i])
	}
	return &Encoder{
		files:   files,
		writers: writers,
		members: members,
		slice:   slice,
	}
}

// SaveMap will save the current wordMap as file slices
func (e *Encoder) SaveWordsMap(wordMap *WordMap) error {
	for word, countIndex := range *wordMap {
		idx := fnv32hash(word) % e.slice
		wordDict := Dictionary{Word: word, CountIndex: countIndex}
		// try encode
		if err := e.members[idx].Encode(wordDict); err != nil {
			return fmt.Errorf("Failed to encode WordDict. err:%v\n", err)
		}
	}
	return nil
}

func (e *Encoder) ensureLifecycle() {
	if err := e.flush(); err != nil {
		log.Fatalf("flush writer err: %v", err)
	}
	if err := e.closeAll(); err != nil {
		log.Fatalf("close file err: %v", err)
	}
}

func (e *Encoder) flush() error {
	for _, w := range e.writers {
		if err := w.Flush(); err != nil {
			return err
		}
	}
	return nil
}

func (e *Encoder) closeAll() error {
	for _, f := range e.files {
		err := f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
