package main

import (
	"fmt"
	"strings"
)

func main222() {
	fmt.Println("a" > "b")
	fmt.Println("z" > "b")

	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("b", "b"))
	fmt.Println(strings.Compare("c", "b"))
}
