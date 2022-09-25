package main

import (
	"fmt"
	"sync"
)

func main() {
	number, cnt := 0, 1000
	var wg sync.WaitGroup
	wg.Add(cnt)
	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				number++
			}
			wg.Done()
		}()
	}
	for i := 0; i < 100; i++ {
		number++
	}
	wg.Wait()
	fmt.Println(number)
}
