package main

import (
	"Asynchronous-HTTPClient/synchronous"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	synchronous.ComputeSynchronous()
	period := time.Since(start).Nanoseconds()
	fmt.Println("period: ", period, "ms")
}
