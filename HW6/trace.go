package main

import (
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	trace.Start(f)
	// trace.Start(os.Stderr)
	defer trace.Stop()
	var mtx sync.Mutex
	for i := 0; i < 5; i++ {
		go func() {
			mtx.Lock()
			defer mtx.Unlock()
			for i := 0; i < 1000; i += 1 {
			}
		}()
	}
}
