### code copied from online comppiler

```go
package main
import "fmt"

func Filter(s []int, fn func(int) bool) []int {
	var p []int // this is nil for now
	for _, i := range s {
	    if fn(i) {
	        p = append(p,i)
	    }
	}
	return p
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
    s := []int{0,1,2,3,4,5,6,7,8}
    s = Filter(s,even)
    fmt.Println(s)
}
```

### notes on how the code works

- even takes in a num and returns a bool
- filter takes a slice of ints s
  - declare a new slice p to put the result nums
  - then loop thru s to find if they're even fn function even
  - if fn(i) is true than it gets appended to the p slice and returned after the loop finishes iterating
