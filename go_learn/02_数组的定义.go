package main

import (
	"fmt"
)

func main() {
	a := [...]int{1: 1, 2, 3}
	for i, v := range a {
		fmt.Println(i, v)

	}
}
