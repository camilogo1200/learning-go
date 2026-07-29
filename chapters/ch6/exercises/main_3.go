package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	// test allocations
	TestAllocationForLoop()
	TestAllocationExponentialCopyPattern()
}

func TestAllocationForLoop() {
	nItems := 10_000_000

	person := Person{
		FirstName: "Cris",
		LastName:  "Ross",
		Age:       150,
	}

	slPerson := make([]Person, nItems)

	for i := 0; i < nItems; i++ {
		slPerson[i] = person
	}

	//fmt.Println(slPerson[9_999_999].FirstName)
}

func TestAllocationExponentialCopyPattern() {
	nItems := 10_000_000

	person := Person{
		FirstName: "Cris",
		LastName:  "Ross",
		Age:       150,
	}

	slPerson := make([]Person, nItems)

	slPerson[0] = person

	for i := 1; i < nItems; {
		i += copy(slPerson[i:], slPerson[:i])
	}

	//fmt.Println(slPerson[9_999_999].FirstName)
}
