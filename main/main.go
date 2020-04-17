package main

import (
	"github.com/rubin55/go-experiments/rabbit"
	"time"
)

func main() {
	rabbit.Publish()
	time.Sleep(2 * time.Second)
	rabbit.Consume()
}
