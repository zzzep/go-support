package setter_test

import (
	"fmt"
	"github.com/zzzep/go-support/setter"
)

func ExampleSetIfNotEmptyString() {
	var name string
	setter.SetIfNotEmptyString(&name, "John")
	fmt.Println(name)

	var email string
	setter.SetIfNotEmptyString(&email, "")
	fmt.Println(email)

	// Output:
	// John
	//
}

func ExampleSetIfNotEmptySlice() {
	var tags []string
	setter.SetIfNotEmptySlice(&tags, []string{"go", "golang", "library"})
	fmt.Println(tags)

	var empty []string
	setter.SetIfNotEmptySlice(&empty, []string{})
	fmt.Println(empty)

	// Output:
	// [go golang library]
	// []
}

func ExampleSetIfNotEmptyMap() {
	var config map[string]string
	setter.SetIfNotEmptyMap(&config, map[string]string{"key": "value"})
	fmt.Println(config)

	var empty map[string]string
	setter.SetIfNotEmptyMap(&empty, map[string]string{})
	fmt.Println(empty)

	// Output:
	// map[key:value]
	// map[]
}

func ExampleSetIfNotEmpty() {
	var count int
	setter.SetIfNotEmpty(&count, 42)
	fmt.Println(count)

	var zero int
	setter.SetIfNotEmpty(&zero, 0)
	fmt.Println(zero)

	var message string
	setter.SetIfNotEmpty(&message, "Hello, World!")
	fmt.Println(message)

	// Output:
	// 42
	// 0
	// Hello, World!
}
