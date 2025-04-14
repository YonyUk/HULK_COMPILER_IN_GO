package hulk

import (
	. "hulk.com/app/grammar"
)

var KeywordsGrammar IGrammar

func init() {
	keywords := []string{
		"print",
		"sqrt",
		"sin",
		"cos",
		"exp",
		"log",
		"rand",
		"function",
		"let",
		"in",
		"return",
		"if",
		"else",
		"elif",
		"while",
		"for",
		"range",
		// "type",
		// "inherits",
		// "Number",
		// "String",
		// "Boolean",
		// "protocol",
		// "extends",
		// "Object",
		// "def",
		// "match",
		// "case",
		// "default",
	}
	KeywordsGrammar = GetWordsGrammar(keywords)
}
