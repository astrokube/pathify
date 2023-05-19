package main

import "encoding/json"

/*
*

	func buildNode() pathifyv2.Node[map[string]any] {
		john := pathifyv2.New().
			SetValue("firstname", "John").
			SetValue("lastname", "Doe").Out()

		wendy := pathifyv2.New().
			SetValue("firstname", "Wendy").Out()

		return pathifyv2.New().
			SetValue("salary", 120).
			SetValue("dad", john).
			SetValue("mom.firstname", "Jane").
			SetValue("mom.lastName", "Dane").
			SetValue("children[0].firstname", "David").
			SetValue("children[0].age", 20).
			SetValue("children[0].tutor", wendy)
	}

*
*/
func main() {
	/**
	out := complexMapStructure()

	b, err := json.MarshalIndent(out, " ", "  ")
	if err != nil {
		println(err.Error())
		//os.Exit(1)
	}
	println(string(b))
	prettyPrint()
	println("------------ Pathify Load ------------")
	**/
	//exampleLoad()
	printJSON(loadArray())
	// printJSON(loadMap())
}

func printJSON(content any) {
	b, _ := json.Marshal(content)
	println(string(b))
}
