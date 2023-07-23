package main

import "fmt"

func main() {
	val := 0
	screen := [2][2]int{}

	//assign vals to screen
	for row := range screen {
		for column := range screen[row] {
			screen[row][column] = val
			val = val + 1
		}
	}

	for row := range screen {
		for column := range screen[0] {
			fmt.Print(screen[row][column], " ")
		}
		fmt.Print("\n")
	}
}
