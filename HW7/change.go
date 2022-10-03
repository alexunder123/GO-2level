package main

import (
	"errors"
	"fmt"
	"reflect"
)

var emptyValue = errors.New("value not found")
var wrongValue = errors.New("wrong map value")

func ChangeIn(in struct{ testStr string }, values map[string]interface{}) error {
	ind := reflect.ValueOf(&in)
	val0 := ind.Elem()
	fmt.Println(val0.Field(0))
	val1 := reflect.ValueOf(values)
	if val1.Kind() != reflect.Map {
		return wrongValue
	}

	iter := reflect.ValueOf(values).MapRange()
	for iter.Next() {
		k := iter.Key()
		if k == val0.Field(0) {
			v := iter.Value()
			val0.SetString(v.Elem().String())
		}
	}
	fmt.Println(val0)

	return nil
}

func main() {
	var in struct{ testStr string }
	in.testStr = "strstr"
	var values = map[string]interface{}{}
	values["mnt"] = "mmm"
	values["strstr"] = "lalala"
	var err error
	err = ChangeIn(in, values)
	if err != nil {
		fmt.Println("Что-то пошло не так:", err)
	}
	fmt.Println("End")
}
