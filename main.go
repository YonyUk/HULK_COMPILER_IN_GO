package main

import (
	. "hulk.com/app/filesystem"
	. "hulk.com/app/hulk"
)

func main() {
	interpreter := NewHulkInterpreter()
	reader, _ := NewFileReader("code.hulk")
	code, _ := reader.ReadFile()
	interpreter.Execute(code)
}
