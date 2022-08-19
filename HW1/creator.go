package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func createFiles(i, n *int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()
	err := os.Mkdir("tmp", 0750)
	if err != nil && !os.IsExist(err) {
		fmt.Errorf("ошибка создания папки %w", err)
		return
	}
	for *i < *n {
		file := "./tmp/" + strconv.Itoa(*i) + ".txt"
		f, err := os.Create(file)
		if err != nil {
			fmt.Errorf("ошибка создания файла %w", err)
		}
		defer f.Close()
		*i++
	}
}

func main() {
	defer func() {
		if v := recover(); v != nil {
			dt := time.Now()
			fmt.Println("Время появления ошибки: ", dt.Format("15:04:05"))
			fmt.Println("recovered", v)
		}
	}()
	i, n := 0, 10
	for i < n {
		createFiles(&i, &n)
	}
}
