package hulk

import (
	. "hulk.com/app/grammar"
)

var OperatorGrammar IGrammar

func init() {
	operators := []string{
		"+", //Arithmetic operators
		"-",
		"*",
		"/",
		"^",
		"%",
		"as",
		"@",
		"@@",
		"&", // boolean operators
		"|",
		"!",
	}
	OperatorGrammar = GetWordsGrammar(operators)
}
