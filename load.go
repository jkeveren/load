package main

import (
	"fmt"
	"runtime"
)

func main() {
	CPUCount := runtime.NumCPU()
	fmt.Printf("Loading all %d logical processors\n", CPUCount)
	for i := 0; i < CPUCount-1; i++ {
		go func() {
			for {
			}
		}()
	}
	for {
	}
}
