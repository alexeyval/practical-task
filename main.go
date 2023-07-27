package main

import (
	"context"
	"time"
)

func main() {
	parent := context.Background()
	var checker = NewChecker(parent)
	go checker.Add(NewGoMetrClient("1", 1))
	go checker.Add(NewGoMetrClient("4", 1))
	go checker.Add(NewGoMetrClient("5", 1))
	go checker.Add(NewGoMetrClient("666", 1))
	go checker.Add(NewGoMetrClient("7", 1))

	go func() {
		for {
			go checker.Add(NewGoMetrClient("7", 1))
			time.Sleep(1 * time.Second)
		}
	}()

	go checker.Run()
	time.Sleep(17 * time.Second)
	checker.Stop()
}
