package main

import (
	"encoding/json"
	"os"

	"github.com/astrokube/pathify"
)

func complexMapStructure() map[string]interface{} {

	john := pathify.New().
		Set("firstname", "John").
		Set("lastname", "Doe").Map()

	wendy := pathify.New().
		Set("firstname", "Wendy").Map()

	return pathify.New().
		Set("salary", 120).
		Set("dad", john).
		Set("mom.firstname", "Jane").
		Set("mom.lastName", "Dane").
		Set("children[0].firstname", "David").
		Set("children[0].age", 20).
		Set("children[0].tutor", wendy).
		Map()
}

func main() {
	out := complexMapStructure()

	b, err := json.MarshalIndent(out, " ", "  ")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(string(b))
}
