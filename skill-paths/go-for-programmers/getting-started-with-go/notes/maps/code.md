## init a map

````go
package main
import "fmt"

func main() {
  mapLit := map[string]int{"one": 1, "two": 2}   // making map & adding key-value pair
  var mapAssigned map[string]int
  mapCreated := make(map[string]float32)        // making map with make()
  mapAssigned = mapLit
  mapCreated["key1"] = 4.5      // creating key-value pair for map
  mapCreated["key2"] = 3.14159
  mapAssigned["two"] = 3        // changing value of already existing key
  fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
  fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
  fmt.Printf("Map assigned at \"two\" is: %d\n", mapAssigned["two"])
  fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
}```
````

## key value item

```go
package main
import "fmt"

func main() {
  var value int
  var isPresent bool
  map1 := make(map[string]int)
  map1["New Delhi"] = 55
  map1["Beijing"] = 20
  map1["Washington"] = 25
  value, isPresent = map1["Beijing"]  // checking existence of a key

  if isPresent {
    fmt.Printf("The value of \"Beijing\" in map1 is: %d\n", value)
  } else {
    fmt.Println("map1 does not contain Beijing")
  }
  value, isPresent = map1["Paris"]   // chekcing existence of a key
  fmt.Printf("Is \"Paris\" in map1 ?: %t\n", isPresent)
  fmt.Printf("Value is: %d\n", value)

  // delete an item:
  delete(map1, "Washington")
  value, isPresent = map1["Washington"] // checking existence of a key

  if isPresent {
    fmt.Printf("The value of \"Washington\" in map1 is: %d\n", value)
  } else {
    fmt.Println("map1 does not contain Washington")
  }
}
```
