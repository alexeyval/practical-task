package main

import (
	"fmt"
)

func main() {
	checker := NewChecker()
	checker.Add(NewGoMetrClient("1", 1))
	checker.Add(NewGoMetrClient("4", 1))
	checker.Add(NewGoMetrClient("5", 1))
	checker.Add(NewGoMetrClient("7", 1))

	fmt.Println(checker)

	checker.Check()
}
