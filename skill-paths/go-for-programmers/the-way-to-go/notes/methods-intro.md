## what is a method?
- a go method is a function that acts on a var of a certain type, called the receiver.
- therefore a method in go is just a special type of function.

- a method acting on a var in go is like an object of a class calling its function in other OO languages using a `.` selector.

- receiver can be of any type except an interface as an interface is an abstract definition while a method is an implementation.

- a method cannot be a pointer type - but it can be a pointer to any of the allowed types. 
- the combo of a struct type and its methods is equivalent to a class in OO.
- the difference is that in Go, the code for the type and the methods dont have to be grouper together and can be found in different source files within the same package

- the collection of all the methods on a given type ```T``` is called the method set of ```T```