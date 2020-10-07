package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Decoder struct {
	files   []*os.File
	members []*gob.Decoder
}

func newDecoder(slice int) *Decoder {
	// create temp files and corresponded decoder
	tempfiles := make([]*os.File, slice)
	tempdecoders := make([]*gob.Decoder, slice)
	for i := 0; i < slice; i++ {
		fh, err := os.Open(fmt.Sprintf("f%d", i))
		if err != nil {
			log.Fatalf("fail io on ./f%d, err: %v", i, err)
		}
		tempfiles[i] = fh
		tempdecoders[i] = gob.NewDecoder(bufio.NewReader(fh))
	}

	return &Decoder{files: tempfiles, members: tempdecoders}
}

func (d *Decoder) evictAll() error {
	for _, fh := range d.files {
		if err := fh.Close(); err != nil {
			return err
		}
	}

	return nil
}
