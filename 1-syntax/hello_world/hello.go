// 這一行代表該檔案所以程式碼都屬於main 套件
// main 是特殊套件 代表要在終端機直接執行
package main

//從fmt套件使用讓文字格式化的程式碼
import (
	"fmt"
)

//函式

func main() {

	var helloWorld = "Hello World"
	//調用fmt套件中的println
	//fmt -> format
	//println -> print line
	fmt.Println(helloWorld)

	a := 10
	fmt.Println(a)
}
