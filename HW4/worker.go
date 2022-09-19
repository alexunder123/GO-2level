package main

import (
	"fmt"
	"time"
)

func main() {
	number := 0

	var work = make(chan struct{}, 100)

	for i := 0; i < 1000; i++ {
		work <- struct{}{}

		go func() {
			defer func() {
				<-work
			}()
			number++
		}()
		time.Sleep(1 * time.Microsecond)
	}
	fmt.Println(number)
}
