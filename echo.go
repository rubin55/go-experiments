package main

import (
	"fmt"
	"os"
	"strings"
)

// A comment about the main function.
func main() {
    //Make it an iterator with index and print.
	fmt.Println(strings.Join(os.Args[1:], " "))
}

