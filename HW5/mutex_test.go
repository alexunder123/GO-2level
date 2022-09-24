package main

import (
	"runtime"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkReadMtx11(b *testing.B) {
	var (
		number float64
		mutex  sync.Mutex
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				go func() {
					mutex.Lock()
					defer mutex.Unlock()
					number = 23545897.12586
				}()
				go func() {
					mutex.Lock()
					defer mutex.Unlock()
					_ = number
				}()
			}
		})
	})
}

func BenchmarkReadRWMtx11(b *testing.B) {
	var (
		number float64
		mutex  sync.RWMutex
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				go func() {
					mutex.Lock()
					defer mutex.Unlock()
					number = 23545897.12586
				}()
				go func() {
					mutex.RLock()
					defer mutex.RUnlock()
					_ = number
				}()
			}
		})
	})
}

func BenchmarkReadMtx19(b *testing.B) {
	var (
		number float64
		mutex  sync.Mutex
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				go func() {
					mutex.Lock()
					defer mutex.Unlock()
					number = 23545897.12586
				}()
				for i := 0; i < 9; i++ {
					go func() {
						mutex.Lock()
						defer mutex.Unlock()
						_ = number
					}()

				}
			}
		})
	})
}

func BenchmarkReadRWMtx19(b *testing.B) {
	var (
		number float64
		mutex  sync.RWMutex
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				go func() {
					mutex.Lock()
					defer mutex.Unlock()
					number = 23545897.12586
				}()
				for i := 0; i < 9; i++ {
					go func() {
						mutex.RLock()
						defer mutex.RUnlock()
						_ = number
					}()

				}
			}
		})
	})
}

func BenchmarkReadMtx91(b *testing.B) {
	var (
		number float64
		mutex  sync.Mutex
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 9; i++ {
					go func() {
						mutex.Lock()
						defer mutex.Unlock()
						number = 23545897.12586
					}()
				}
				go func() {
					mutex.Lock()
					defer mutex.Unlock()
					_ = number
				}()

			}
		})
	})
}

func BenchmarkReadRWMtx91(b *testing.B) {
	var (
		number float64
		mutex  sync.RWMutex
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 9; i++ {
					go func() {
						mutex.Lock()
						defer mutex.Unlock()
						number = 23545897.12586
					}()
				}
				go func() {
					mutex.RLock()
					defer mutex.RUnlock()
					_ = number
				}()
			}
		})
	})
}
