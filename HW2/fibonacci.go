package main

import "fmt"

func fib_calc(a, b, c, d int) int {
	if c == d {
		return b
	}
	d++
	return fib_calc(b, a+b, c, d)
}

func main() {
	var a, b, c int
	a, b = 0, 1
	fmt.Println("Введите порядковый номер числа Фибоначчи: ")
	fmt.Scanln(&c)

	if c <= 0 {
		fmt.Println("Введено неверное значение")
	} else if c == 1 {
		fmt.Println("Значение числа Фибоначчи равно 0")
	} else if c == 2 {
		fmt.Println("Значение числа Фибоначчи равно 1")
	} else {
		d := 2
		fmt.Println("Значение числа Фибоначчи равно ", fib_calc(a, b, c, d))
	}

}
