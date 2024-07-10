package main

import (
	"fmt"
)

var version float64 = 1.01
var author string = "Bryan Blackburn"

func main() {
	fmt.Printf("%v, Version:%v", author, version)
}
