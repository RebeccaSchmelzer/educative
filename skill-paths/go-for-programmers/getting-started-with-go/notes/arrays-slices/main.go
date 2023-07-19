package main

import "fmt"

func f(a [3]int)   { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }

//declaring array literals - kinda like contructors when some values are known beforehand. use {} to only init a couple items but also works when you wanna init everything
/**1st variant*/
var arrAge = [5]int{1, 2, 3, 4, 5}
var arrVar = [5]int{1, 2, 4} //only 3 nums declared here - others init to 0

//2nd variant
var arrLazy = [...]int{1, 2, 4}

//[...] indicates the compiler has to count the num of items to get the len of the array
//[...]int is not a type tho so you cant declare var arrLazy [...]int = blah
//if [...] is omitted then a slice is made

//3rd variant
var arrKey = [5]string{3: "me"}

/** Passing an array into a function
- dont use up all your memory in copying arrays
- pass a pointer to the array or use a slice of the array
**/
func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}

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

	array := [3]float64{7.0, 24.4, 9.1}
	x := Sum(&array)
	fmt.Printf("yo this is my sum: %f", x)
}
