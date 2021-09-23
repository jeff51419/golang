package main

import (
    "fmt"
)

func demo(s string, x int) string {
    ret := fmt.Sprintf("This is a demo, your input is: %s %d", s, x)
    return ret
}

func main() {
    s := "string"
    i := 1111
    fmt.Println(demo(s, i))
}