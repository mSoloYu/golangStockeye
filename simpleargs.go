package main

import (
	"fmt"
	"os"
)

func main() {
	size := len(os.Args)
	fmt.Println(size)
	for size > 1 {
		fmt.Println(os.Args[size-1])
		size--
	}
}
