package main

import (
	"fmt"

	. "hulk.com/app/filesystem"
	. "hulk.com/app/lgen"
)

func main() {
	file_reader, _ := NewFileReader("code.hulk")
	code, _ := file_reader.ReadFile()
	lexer := LGENInterpreter
	lexer.LoadCode(code)
	for lexer.Next() {
		fmt.Println(lexer.Current())
	}
	// for r := rune(0); r < rune(256); r++ {
	// 	fmt.Println(r, ":", string(r))
	// }
}
