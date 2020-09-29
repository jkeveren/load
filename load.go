package main

import (
	"fmt"
	"runtime"
)

func main() {
	CPUCount := runtime.NumCPU()
	fmt.Printf("Loading %d logical processors\n", CPUCount)
	for i := 0; i < CPUCount-1; i++ {
		go func() {
			for {
			}
		}()
	}
	for {
	}
}
