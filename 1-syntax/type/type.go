package main

import (
	"fmt"
	"reflect"
	// "math"
	// "strings"
)

func main() {
	// fmt.Println(math.Floor("Head first go"))
	// fmt.Println(strings.Title(2.75))
	// go 屬於靜態型別 statically typed
	// 函式預期他們的引數是特定的型別

	// 回傳引數的型別
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello"))
}
