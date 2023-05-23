package main

import (
	"strings"

	"github.com/astrokube/pathify/pathifier"

	"github.com/astrokube/pathify"
)

var peopleArray = []any{
	map[string]any{
		"firstname": "John",
		"lastname":  "Doe",
		"age":       29,
	},
	map[string]any{
		"firstname": "Jane",
		"lastname":  "Moe",
		"age":       30,
	},
}

func main() {
	p := pathify.Load[[]any](peopleArray).Set(
		"[1].lastname", "Doe",
		"[0].firstname", "Wendy",
		"[2].firstname", "Cindy",
		"[1].firstname", strings.ToUpper,
	)
	println(p.String(pathifier.WithOutputFormat(pathifier.YAML)))
}
