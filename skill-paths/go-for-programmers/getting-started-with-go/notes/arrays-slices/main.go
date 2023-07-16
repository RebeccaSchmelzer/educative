package main

import "fmt"

func f(a [3]int)   { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }

func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("items at index %d is %d\n", i, arr[i])
	}
	var arr2 [3]int
	f(arr2)
	fp(&arr2)
}
