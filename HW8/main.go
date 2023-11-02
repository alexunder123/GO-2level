package main

import (
	"bufio"
	"crypto"
	_ "crypto/MD5"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	NameStack  []string
	HashStack  []string
	AddrStack  []string
	Dublicates []int
	QDubl      []int
)

func main() {

	var fl, adr string

	flag.StringVar(&adr, "a", "", "адрес каталога для считывания")
	flag.StringVar(&fl, "r", "", "удалить файлы после поиска Y/y")
	flag.Parse()

	fmt.Println("Программа поиска дубликатов")

	if adr == "" {
		fmt.Println("Введите путь к папке:")
		cat := bufio.NewScanner(os.Stdin)
		cat.Scan()
		adr = cat.Text()
	}

	if !strings.HasSuffix(adr, "\\") {
		adr += "\\"
	}
	addr, err := os.Stat(adr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if !addr.IsDir() {
		fmt.Println("Путь не является папкой")
		os.Exit(0)
	}

	ReadDirectory(adr)
	SearcDubl()
	PrintDubl()

	if fl == "" {
		fmt.Print("Удалить дубликаты? Y/y: ")
		_, err = fmt.Scanln(&fl)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}

	if fl == "Y" || fl == "y" {
		fmt.Print("Вы точно хотите удалить дубликаты? Y/y: ")
		_, err = fmt.Scanln(&fl)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if fl == "Y" || fl == "y" {
			DeleteDubl()
		}
	}

	fmt.Println("Программа выполнена")
}

func ReadDirectory(adr string) {
	catalogue, err := os.ReadDir(adr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	for _, file := range catalogue {
		if !file.IsDir() {
			checksum := FileMD5(adr + file.Name())
			NameStack = append(NameStack, file.Name())
			HashStack = append(HashStack, checksum)
			AddrStack = append(AddrStack, adr+file.Name())
		} else {
			ReadDirectory(adr + file.Name() + "\\")
		}
	}
	return
}

func FileMD5(path string) string {
	h := crypto.MD5.New()
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func SearcDubl() {
	for i := 0; i < len(NameStack)-1; i++ {
		for _, number := range Dublicates {
			if number == i {
				i++
			}
		}
		d := false
		q := 0
		for j := i + 1; j < len(NameStack); j++ {
			if NameStack[i] == NameStack[j] && HashStack[i] == HashStack[j] {
				if !d {
					Dublicates = append(Dublicates, i, j)
					d = true
					q++
				} else {
					Dublicates = append(Dublicates, j)
					q++
				}
			}
		}
		if q > 0 {
			QDubl = append(QDubl, q)
		}
	}
}

func PrintDubl() {
	j := 0
	for _, number := range QDubl {
		fmt.Println("Найдена группа дубликатов")
		for i := 0; i <= number; i++ {
			fmt.Println(AddrStack[Dublicates[j]])
			j++
		}
	}
}

func DeleteDubl() {
	j := 0
	for _, number := range QDubl {
		for i := 0; i <= number; i++ {
			if i > 0 {
				os.Remove(AddrStack[Dublicates[j]])
			}
			j++
		}
	}
}
