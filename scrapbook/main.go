package main

import (
	_ "embed"
	"fmt"
)

//go:embed data.txt
var s string

func main() {

	fmt.Println(s)

}
