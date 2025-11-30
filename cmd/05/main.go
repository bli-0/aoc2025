package main

import (
	_ "embed"
	"fmt"
)

//go:embed inputs
var input string

func main() {
	fmt.Println(input)
}
