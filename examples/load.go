package main

import (
	"github.com/astrokube/pathify"
)

func loadArray() any {
	return pathify.Load[[]any]([]any{
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
	}).
		Set(
			"[1].lastname", "Doe",
			"[0].firstname", "Wendy",
			"[2]", map[string]any{
				"firstname": "David",
				"lastname":  "Loe",
				"age":       23,
			},
		).
		/**
		SetValue("[0].age", func(x int) int {
			return x + 10
		}).**/Out()
}

func loadMap() any {
	return pathify.Load[map[string]any](map[string]any{
		"firstname": "Jane",
		"lastname":  "Doe",
		"age":       29,
	}).Set(
		"lastname", "Moe").Out()
}
