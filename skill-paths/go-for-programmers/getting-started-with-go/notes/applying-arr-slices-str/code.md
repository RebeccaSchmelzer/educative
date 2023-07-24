### code copied and commented from lesson

```go
package main
import "fmt"

/* making a slice of bytes/runes from a string
- remember that a string in mem is the pointer to the data and the len
- string is a val type - an array of characters
- if s is a str - a slice of bytes c can be made with with
*/
// c := []byte(s)

// //can also use the copy function
// copy(dst []byte, src string)

var arr []byte = []byte{'a','b','a','a','a','c','d','e','f','g'}

func main() {
    //make a global array of type byte and init it to have length of the arr
	arru := make([]byte,len(arr)) // this will contain consecutive non-repeating items
	ixu := 0 // index in arru
	//keep track of previously visited element of arr
	tmp := byte(0)
	for _, val := range arr {
		if val!=tmp {
			arru[ixu] = val
			fmt.Printf("%c ", arru[ixu])
			ixu++
		}
		tmp = val
	}
}
```
