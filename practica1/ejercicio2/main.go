package main

import (
	"fmt"
)

func main() {
  maxAsterisks := 5

	height := (maxAsterisks + 1) / 2

	for i := 1; i <= height; i++ {
		for j := 1; j <= height-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := height - 1; i >= 1; i-- {
		for j := 1; j <= height-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

