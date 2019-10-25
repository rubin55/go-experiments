package main

import (
	"fmt"
	"github.com/rubin55/go-experiments/echo"
	"github.com/rubin55/go-experiments/hello"
)

func main() {
	fmt.Println(hello.Hello() + " " + echo.Echo())
}
