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

	david := pathify.New().
		Set("firstnmae", "David").Map()

	return pathify.New().
		Set("salary", 120).
		Set("dad", john).
		Set("mom.firstname", "Jane").
		Set("mom.lastName", "Dane").
		Set("children", []map[string]interface{}{}).
		Set("children[0]", david).
		Set("children[0].age", 20).
		Map()

}

func main() {
	//out:=complexMapStructure()
	out := pathify.New().
		Set("group.members[0].skills.speed", 50).
		Set("group.members[1].skills.strength", 20).
		Fill().
		Map()

	b, err := json.MarshalIndent(out, " ", "  ")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(string(b))
}
