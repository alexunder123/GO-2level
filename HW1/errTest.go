package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

func calc(a, b int, op string) int {
	var res int
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	return res
}

type ErrorWithTrace struct {
	text  string
	trace string
}

func New(text string) error {
	return &ErrorWithTrace{
		text:  text,
		trace: string(debug.Stack()),
	}
}

func (e *ErrorWithTrace) Error() string {
	return fmt.Sprintf("error: %s\ntrace:\n%s", e.text, e.trace)
}

func main() {
	var err error
	defer func() {
		if v := recover(); v != nil {
			dt := time.Now()
			fmt.Println("Время появления ошибки: ", dt.Format("15:04:05"))
			err = New("my error")
			fmt.Println(err)
		}
	}()

	var a, b, res int
	var op string
	a = 5
	op = "/"
	res = calc(a, b, op)
	fmt.Println("Результат операции: ", res)
}
