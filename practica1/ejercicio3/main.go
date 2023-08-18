package main

import (
	"fmt"
)

func rotate(slice []string, rotations int, direction int) {
	length := len(slice)
	if length == 0 {
		return
	}
	rotations %= length

	if direction == 1 {
		copy(slice, append(slice[length-rotations:], slice[:length-rotations]...))
	} else if direction == 0 {
		copy(slice, append(slice[rotations:], slice[:rotations]...))
	}
}

func main() {
	slice := []string{"a", "b", "c", "d","e","f","g","h"}
	var rotations, direction int

	fmt.Println("Ingrese el número de rotaciones: ")
	fmt.Scan(&rotations)

	fmt.Println("Ingrese el lado de rotación ( 0 = left / 1 = right ): ")
	fmt.Scan(&direction)

	fmt.Println("Antes de la rotación:", slice)
	rotate(slice, rotations, direction)
	fmt.Println("Después de la rotación:", slice)
}
