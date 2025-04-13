package hulk

import (
	. "hulk.com/app/grammar"
)

var KeywordsGrammar IGrammar

func init() {
	keywords := []string{
		"function",
		"if",
		"else",
	}
	KeywordsGrammar = GetWordsGrammar(keywords)
}
