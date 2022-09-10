//Package fibonacci
//
//Функция вычисляет число фибоначчи рекурсивным методом
//Функция принимает 4 значения: предыдущее и текущее вычисленные значения, порядковый номер вычисленного значения и порядковый номер требуемого значения.
//
package doc

//fib_calc returns an int
func fib_calc(a, b, c, d int) int {
	if c == d {
		return b
	}
	d++
	return fib_calc(b, a+b, c, d)
}
