package main

import (
	"fmt"
	"net/http"
)

var version float64 = 1.01
var author string = "Bryan Blackburn"
var PORT int = 8080

func main() {

	http.ListenAndServe(":8080", nil)

	fmt.Printf("Server Started on Port: %v \n", PORT)
	fmt.Printf("%v, Version:%v", author, version)

}
