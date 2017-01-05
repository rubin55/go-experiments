package main

import "fmt"
import "os"
import "strconv"

// A comment about the main function.
func main() {
	s := ""
	for idx, arg := range os.Args[1:] {
		s += strconv.Itoa(idx) + " " + arg + "\n"
	}
	fmt.Println(s)
	//Todo: Implement metrics for performance.
}
