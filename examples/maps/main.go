package main

import (
	"github.com/astrokube/pathify"
	"github.com/astrokube/pathify/pathifier"
)

func main() {
	println(pathify.New().Set("firstname", "John", "parent.firstname", "David").String(pathifier.AsJSON))
}
