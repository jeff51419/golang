package main

import "fmt"

func main() {
	// 宣告變數
	var quantity int = 4
	var length, width float64 = 1.2, 2.4
	var customerName string = "Damon Cole"

	// 指派值給變數
	// quantity = 4
	// length, width = 1.2, 2.4
	// customerName = "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}
