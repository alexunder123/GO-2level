package main

import (
	"fmt"
	"sync"
)

func main() {
	number, cnt := 0, 1000

	var wg sync.WaitGroup
	var mtx sync.Mutex
	wg.Add(cnt)
	for i := 0; i < 1000; i++ {
		go func() {
			mtx.Lock()
			defer mtx.Unlock()
			number++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(number)
}
