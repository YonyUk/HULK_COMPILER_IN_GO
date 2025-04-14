package hulk

import (
	. "hulk.com/app/grammar"
)

var BooleanGrammar IGrammar

func init() {
	literals := []string{
		"true",
		"false",
	}
	BooleanGrammar = GetWordsGrammar(literals)
}
