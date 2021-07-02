package main

import (
	"fmt"
	"reflect"
	"strings"
)
func main() {
	var arr []interface{}
	var data string
	fmt.Println("enter your charcater")
	fmt.Scanln(&data)
	if strings.Contains(data, ".") {
		arr = append(arr, 0.0)
	}
	if strings.Contains(data, "true") || strings.Contains(data, "false") {
		arr = append(arr, true)
	}
	arr = append(arr, 0, "")
	var value interface{}
	var err error
	for _, v := range arr {
		value, err = convertType(data, v)
		if err == nil {
			break
		}
	}
	fmt.Printf("Data Type of %s is :%T\n", data, value)

}

func convertType(data string, v interface{}) (interface{}, error) {
	a := reflect.TypeOf(v)
	value := reflect.New(a)

	_, err := fmt.Sscan(data, value.Interface())
	if err != nil {
		fmt.Print("error", err)
	}
	return value.Elem().Interface(), err
}
