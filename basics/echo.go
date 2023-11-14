package main

import (
	"fmt"
	"os"
) // multiple imports go in paranthesies

func main() {
	var s, sep string //if variable is not explicitly defined it will start from the zero value of it's type
	sep = " "
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Print(s)
}
