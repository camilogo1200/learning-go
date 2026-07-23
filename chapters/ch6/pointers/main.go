package main

import "fmt"

type person struct {
	Name       string
	MiddleName *string
	Lastname   string
	Age        int32
}

func main() {
	var x int32 = 30
	pointer := &x

	//Go is a language that manages parameters as pass-by-value so inmutability is reflected with the pointer parameters

	//Go does not have pointer arithmetics like C++
	//&address operator
	// * the indirection operator, next to a pointer variable returns the pointed-to value
	fmt.Printf(" Pointer memory address = {%v} - value ={%v}\n", pointer, *pointer)

	//zero value for pointers are nil
	var pointertoInt *int32 = nil
	pointertoInt = &x
	fmt.Printf(" Pointer memory address = {%v} - value ={%v}\n", pointertoInt, *pointertoInt)

	//The built-in function `new` creates a pointer variable. it returns a pointer to a zero value instance of the provided type

	var booleanPtr = new(bool)
	var floatPtr = new(float32)

	fmt.Printf("booleanPtr is null ={%v} - memory address ={%v} initial value = {%v}\n", booleanPtr == nil, booleanPtr, *booleanPtr)
	fmt.Printf("booleanPtr is null ={%v} - memory address ={%v} - initial value = {%v}\n", floatPtr == nil, floatPtr, *floatPtr)

	//for creating a pointer, you can't use the & and the type,
	//however you can create a pointer to a struct using the & followed by the struct literal
	personPtr := &person{}
	fmt.Printf("personPtr is null ={%v} - memory address ={%v} - initial value = {%v}\n", personPtr == nil, personPtr, *personPtr)

	person2Ptr := &person{
		Name:       "Cris",
		MiddleName: makePointer("Johnson"),
		Lastname:   "Specter",
		Age:        33,
	}
	fmt.Printf("person2Ptr is null ={%v} - memory address ={%v} - initial value = {%v}\n", person2Ptr == nil, person2Ptr, *person2Ptr)

	person3Ptr := &person{
		Name:       "Katty",
		MiddleName: new("Johnson"),
		Lastname:   "Newark",
		Age:        33,
	}
	fmt.Printf("person2Ptr is null ={%v} - memory address ={%v} - initial value = {%v}\n", person3Ptr == nil, person3Ptr, *person3Ptr)

	//to change the value on a pointer function parameter you need to dereference if not the values is not changed
	demo := "hello"
	failedToUpdateParameter(&demo)
	fmt.Printf("failed to update the value of Ptr = {%v}\n", demo)
	updatePointer(&demo)
	fmt.Printf("update the value of Ptr = {%v}\n", demo)

}

func updatePointer(s *string) {
	*s = "The updated Hello World!"
}

func failedToUpdateParameter(s *string) {
	x2 := "My newest hello world failed!"
	s = &x2
	//s = new("Hello world!")

}

/*Just a helper function to convert a literal to a pointer of the provided type*/
func makePointer[T any](t T) *T {
	return &t
}
