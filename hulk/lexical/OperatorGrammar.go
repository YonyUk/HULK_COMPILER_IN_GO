package hulk

import (
	. "hulk.com/app/grammar"
)

var OperatorGrammar IGrammar

func init() {
	operators := []string{
		"+",
		"-",
		"*",
		"/",
		"^",
		"%",
		"as",
		"@",
		"@@",
	}
	OperatorGrammar = GetWordsGrammar(operators)
}
