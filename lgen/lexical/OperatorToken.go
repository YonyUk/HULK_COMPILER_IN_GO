package lexical

import (
	. "hulk.com/app/grammar"
)

var OperatorTokenGrammar IGrammar

func init() {
	operators := []string{
		"=",
		"<",
		">",
		"|",
	}
	OperatorTokenGrammar = GetWordsGrammar(operators)
}
