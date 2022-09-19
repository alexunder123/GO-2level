package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ExitFunc() {
	fmt.Println("Завершение процессов")
	os.Exit(0)
}

func main() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range sigChan {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("Начинаем выход из программы")
				ExitFunc()
			}
		}
	}()
	fmt.Println("Программа запущена")
	time.Sleep(20 * time.Second)
}