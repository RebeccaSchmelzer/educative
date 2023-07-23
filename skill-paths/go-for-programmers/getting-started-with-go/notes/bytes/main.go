package main

import "bytes"
import "fmt"

/*bytes and slices
- go has a package spec for functions using that type
- contains the type buffer and similar to the strings package
- buffer is a var sized buffer of bytes with read and write methods bc reading and writing an unknown num of bytes is best done buffered
*/

//create a buffer
var buffer bytes.Buffer

//or create as a pointer with new
var r *bytes.Buffer = new(bytes.Buffer)

func main() {
	//create a new buffer
	var b bytes.Buffer

	//write strings to the buffer
	b.WriteString("ABC")
	b.WriteString("DEF")

	//use fprintf with the buffer
	fmt.Fprintf(&b, "A num: %d, a string: %v\n", 10, "bird")

	fmt.Println(b.String())
}
