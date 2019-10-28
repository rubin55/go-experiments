package main

import (
	"fmt"
	"github.com/rubin55/go-experiments/echo"
	"github.com/rubin55/go-experiments/hello"
	"github.com/rubin55/go-experiments/unique"
	"os"
)

func main() {
	fmt.Println(hello.Hello() + " " + echo.Echo(os.Args))
	fmt.Println(unique.NewUnique(os.Args[1:]))
}
