package convert

import (
	"fmt"
)

func ExampleToPtr() {
	// Convert string to *string
	str := "hello"
	strPtr := ToPtr(str)
	fmt.Println(*strPtr)

	// Convert int to *int
	num := 42
	numPtr := ToPtr(num)
	fmt.Println(*numPtr)

	// Convert float64 to *float64
	pi := 3.14
	piPtr := ToPtr(pi)
	fmt.Println(*piPtr)

	// Convert bool to *bool
	flag := true
	flagPtr := ToPtr(flag)
	fmt.Println(*flagPtr)

	// Output:
	// hello
	// 42
	// 3.14
	// true
}

func ExampleToPtr_struct() {
	type Person struct {
		Name string
		Age  int
	}

	// Convert struct to *struct
	person := Person{Name: "John", Age: 30}
	personPtr := ToPtr(person)
	fmt.Println(personPtr.Name, personPtr.Age)

	// Output:
	// John 30
}

func ExampleToPtr_slice() {
	// Convert slice to *slice
	numbers := []int{1, 2, 3}
	numbersPtr := ToPtr(numbers)
	fmt.Println((*numbersPtr)[0], (*numbersPtr)[1], (*numbersPtr)[2])

	// Output:
	// 1 2 3
}
