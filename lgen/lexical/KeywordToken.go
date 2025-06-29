package lexical

import (
	. "hulk.com/app/grammar"
)

var KeywordTokenGrammar IGrammar

func init() {
	keywords := []string{
		"token",
	}
	KeywordTokenGrammar = GetWordsGrammar(keywords)
}
