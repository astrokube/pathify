package main

import (
	"github.com/astrokube/pathify"
	"github.com/astrokube/pathify/pathifier"
	"strings"
)

func main() {
	println(pathify.New().
		With(
			pathifier.WithStringPrefix("person."),
		)("name", "Jon", "lastname", "doe").
		Set("firstname", "John", "parent.firstname", "David").
		With(pathifier.WithFuncPrefix(strings.ToUpper))("greeting", "hello").
		String(pathifier.AsJSON))
}
