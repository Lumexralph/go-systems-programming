package main

import (
	"time"
	"fmt"
	"runtime"
)

func main() {
	var memRunTime runtime.MemStats
	printStats(memRunTime)

	for i := 0; i < 40; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}

	printStats(memRunTime)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}

		// Pause the current process or thread
		time.Sleep(5 * time.Second)
	}
	printStats(memRunTime)
}

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Allocation: ", mem.Alloc)
	fmt.Println("mem.TotalAllocation: ", mem.TotalAlloc)
	fmt.Println("mem.HeapAllocation: ", mem.HeapAlloc)
	fmt.Println("mem.NumGarbageCollector: ", mem.NumGC)
	fmt.Println("-------")
}
