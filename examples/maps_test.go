package main

import (
	"fmt"
	"strings"

	"github.com/astrokube/pathify"
	"github.com/astrokube/pathify/pathifier"
)

func ExampleMapWithAttributes() {
	value := pathify.Map().
		With(
			pathifier.WithStringPrefix("person."),
		)("name", "David", "lastname", "doe").
		Set("firstname", "John", "parent.firstname", "David").
		With(pathifier.WithFuncPrefix(strings.ToUpper))("greeting", "hello").
		String(pathifier.AsJSON)
	fmt.Println(value)
	// Output:
	// {"GREETING":"hello","firstname":"John","parent":{"firstname":"David"},"person":{"lastname":"doe","name":"David"}}
}
