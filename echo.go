package main

import (
	"fmt"
	"os"
	"strings"
)

// A comment about the main function.
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

