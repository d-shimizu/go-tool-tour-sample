package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Goog morning!")
	case t.Hour() <17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Goog evening.")
	}
}
