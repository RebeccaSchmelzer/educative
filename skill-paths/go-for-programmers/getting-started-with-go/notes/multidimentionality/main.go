package main

import "fmt"

/* Multidimentional arrays notes
- arrays are always 1d but can be composed to form multi dimen arrays
- multi arrays are always rectangular
*/

/*Multidimentional slices
- slices are always 1d but can be composed to handle 2d and more
- with slices of slices - that makes an array of slices
*/

const (
	WIDTH  = 1930 //columns of 2d array
	HEIGHT = 1080 //rows of 2d array
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func main() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
	//fmt.Printf("The screen :%d", screen)
	vals := [][]int{} //2d slice

	//init the first 2 rows
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}

	//append each row to the 2d slice
	vals = append(vals, row1)
	vals = append(vals, row2)

	//print first and second row
	fmt.Printf("row 1: %d\n", vals[0])
	fmt.Printf("row 2: %d\n", vals[1])

	//first element
	fmt.Printf("first element: %d\n", vals[0][0])

	//print the entire slice
	fmt.Printf("the slice: %d\n", vals)
}
