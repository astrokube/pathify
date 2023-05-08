package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/astrokube/pathify"
)

func buildNode() pathify.Node {
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
		Set("children[0].tutor", wendy)
}

func complexMapStructure() map[string]interface{} {
	return buildNode().Map()
}

func prettyPrint() {
	fmt.Println(buildNode().PrettyPrint())
}

func main() {
	out := complexMapStructure()

	b, err := json.MarshalIndent(out, " ", "  ")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(string(b))
	prettyPrint()
}
