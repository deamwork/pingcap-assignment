package main

import (
	"hash/fnv"
	"log"
	"os"
)

const int32Max = 0x7FFFFFFF

func isLetter(b byte) bool {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		return true
	}

	return false
}

// using fnv mapReduce (FNV-1a) to hash a string, produce int
func fnv32hash(str string) int {
	hash := fnv.New32a()
	_, _ = hash.Write([]byte(str))

	return int(hash.Sum32() & int32Max)
}

func ensureFileClose(file *os.File) {
	if err := file.Close(); err != nil {
		log.Fatalf("close file err: %v", err)
	}
}
