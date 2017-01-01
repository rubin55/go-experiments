package main

import (
	"fmt"
	"os"
	"strings"
)

// A comment about the main function.
func main() {
    //Todo 1: Make it an iterator with index and print.
    //Todo 2: Implement metrics for performance.
	fmt.Println(strings.Join(os.Args[1:], " "))
}

