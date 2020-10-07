package main

import (
	"log"
	"runtime"
)

const megabytes = uint64(1024 * 1024)

var lastFree uint64

func statMemoryProfile() {
	var memory runtime.MemStats
	runtime.ReadMemStats(&memory)

	log.Printf("gcRound: %v, alloc/total: %v / %v, sys: %v, lastGCFree: %v",
		memory.NumGC, memory.Alloc/megabytes, memory.TotalAlloc/megabytes,
		memory.Sys/megabytes,
		((memory.TotalAlloc-memory.Alloc)-lastFree)/megabytes,
	)

	lastFree = memory.TotalAlloc - memory.Alloc
}
