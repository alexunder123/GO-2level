package main

import (
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	f, err := os.Create("sched.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	var mtx sync.Mutex
	for i := 0; i < 8; i++ {
		go func() {
			mtx.Lock()
			defer mtx.Unlock()
			for i := 0; i < 10000000; i += 1 {
				if i%1000000 == 0 {
					runtime.Gosched()
				}
			}
		}()
	}
}
