package main

import (
	"Asynchronous-HTTPClient/asynchronous"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	asynchronous.ComputeAsynchronous()
	period := time.Since(start).Nanoseconds()
	fmt.Println("period: ", period, "ms")
}
