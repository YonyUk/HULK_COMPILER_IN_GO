package main

import (
	. "hulk.com/app/generator"
)

func main() {
	input := NewProjectInput{
		TokenTypes: []string{
			"LiteralNumber",
			"LiteralString",
		},
		Name: "LGEN",
	}
	path := "../LGEN"
	BuildProject(path, input)
}
