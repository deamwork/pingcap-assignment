package main

import (
	"bufio"
	"encoding/gob"
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

var fileFlag = flag.String("file", "foo.bar", "file input for search")

const sliceLen = 30
const mapLen = sliceLen * 1000

func main() {
	flag.Parse()
	log.Println("split file..")

	sequenceTotal, err := splitFile(*fileFlag, sliceLen)
	if err != nil {
		log.Fatalf("split fail, err: %v", err)
	}

	log.Println("search word..")

	decoder := newDecoder(sliceLen)
	defer decoder.evictAll()

	// init final target
	targetWord := Dictionary{
		Word: "",
		CountIndex: CountIndex{
			Count:    0,
			Sequence: sequenceTotal,
		},
	}

	// main logic here
	for _, decoderInstance := range decoder.members {
		UniqDictMap := NewUniqWordMap(decoderInstance)

		// get first non-repeated word from the slice
		stageTargetWord := UniqDictMap.firstWord(sequenceTotal)

		// compare with other slice file
		if stageTargetWord.Sequence < targetWord.Sequence {
			targetWord = stageTargetWord
		}

		// evict UniqDictMap
		UniqDictMap = nil

		statMemoryProfile()
	}

	log.Printf("target word: [%s], sequence: %d, tmploc:%d", targetWord.Word, targetWord.Sequence, targetWord.Count)
}

func splitFile(path string, slice int) (int, error) {
	encoder := newEncoder(slice)
	defer encoder.ensureLifecycle()

	f, err := os.Open(path)
	if err != nil {
		return -1, errors.New("cannot open file handler")
	}
	defer ensureFileClose(f)

	// then, use bufio to streaming procress target file
	reader := bufio.NewReader(f)

	// stat file and get size
	stat, err := f.Stat()
	if err != nil {
		return -1, errors.New("cannot stat file")
	}

	size := int(stat.Size())
	sliceSize := size / slice
	filePointer := 0

	wordMap := make(WordMap, mapLen)

	// set 64 bytes if case of long word
	buffer := make([]byte, 0, 32)

	// use sequence to store each word position, a.k.a. sequence
	sequence := 0

	for {
		b, err := reader.ReadByte()
		if err != nil {
			// break if reach end of file
			if err == io.EOF {
				break
			}

			return -1, err
		}

		// check if byte is a letter
		if isLetter(b) {
			buffer = append(buffer, b)
		} else {
			// save buffer (word)
			if len(buffer) != 0 {
				word := string(buffer)
				sequence++
				wordMap.add(word, sequence)

				// flush buffer
				buffer = buffer[:0]
			}
		}

		// avoid OOM, write to disk if needed
		if filePointer > sliceSize {
			if err := encoder.SaveWordsMap(&wordMap); err != nil {
				log.Fatalln("fail to save word map")
			}
			statMemoryProfile()

			// reset
			wordMap = make(WordMap, mapLen)
			filePointer = 0
		}
		filePointer++
	}

	// save all if loop ended
	if err := encoder.SaveWordsMap(&wordMap); err != nil {
		log.Fatalln("fail to save word map")
	}
	statMemoryProfile()

	return sequence, nil
}

// NewUniqWordMap will load temp file, merge reconstruct wordmap, and return it.
func NewUniqWordMap(decoder *gob.Decoder) *WordMap {
	uniqueWordMap := make(WordMap, mapLen)

	for {
		dict := Dictionary{}
		if err := decoder.Decode(&dict); err != nil {
			// deal with eof first
			if err == io.EOF {
				break
			}

			log.Fatalf("io on temp file err: %v", err)
		}

		if index, exist := uniqueWordMap[dict.Word]; exist {
			uniqueWordMap[dict.Word] = CountIndex{
				Count:    dict.Count + index.Count,
				Sequence: index.Sequence,
			}
		} else {
			uniqueWordMap[dict.Word] = CountIndex{
				Count:    dict.Count,
				Sequence: dict.Sequence,
			}
		}
	}

	return &uniqueWordMap
}
