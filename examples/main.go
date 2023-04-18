package main

import (
	"encoding/json"
	"os"

	"github.com/astrokube/pathify"
)

func main() {

	var john = pathify.New().
		Set("firstname", "John").
		Set("lastname", "Doe").Map()

	var david = pathify.New().
		Set("firstnmae", "David").Map()

	var out = pathify.New().
		Set("salary", 120).
		Set("dad", john).
		Set("mom.firstname", "Jane").
		Set("mom.lastName", "Dane").
		Set("children", []map[string]interface{}{}).
		Set("children[0]", david).
		Set("children[0].age", 20).
		Map()

	b, err := json.Marshal(out)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(string(b))

}
