package sintax

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/grammar"
)

var TokenGrammar IAttributedGrammar

func init() {
	TOKEN := NewGrammarSymbol("TOKEN", NonTerminal, false)
	LIST := NewGrammarSymbol("LIST", NonTerminal, false)
	SEQUENCE := NewGrammarSymbol("SEQUENCE", NonTerminal, false)
	TokenDeclaration := NewGrammarSymbol("TokenDecl", NonTerminal, false)

	equal := NewGrammarSymbol("=", Terminal, false)
	comma := NewGrammarSymbol(",", Terminal, false)
	OpenList := NewGrammarSymbol("[", Terminal, false)
	CloseList := NewGrammarSymbol("]", Terminal, false)
	text := NewGrammarSymbol("text", Terminal, false)
	variable := NewGrammarSymbol("variable", Terminal, false)
	token := NewGrammarSymbol("token", Terminal, false)

	TokenGrammar = NewAttributedGrammar(TOKEN)
	TokenGrammar.AddProduction(TOKEN, []IGrammarSymbol{TokenDeclaration, equal, LIST}, func(i []IAST, s string) IAST {
		return i[0]
	})
	TokenGrammar.AddProduction(TokenDeclaration, []IGrammarSymbol{token, variable}, func(i []IAST, s string) IAST {
		return i[0]
	})
	TokenGrammar.AddProduction(LIST, []IGrammarSymbol{OpenList, SEQUENCE, CloseList}, func(i []IAST, s string) IAST {
		return i[1]
	})
	TokenGrammar.AddProduction(SEQUENCE, []IGrammarSymbol{text}, func(i []IAST, s string) IAST {
		return i[0]
	})
	TokenGrammar.AddProduction(SEQUENCE, []IGrammarSymbol{text, comma, SEQUENCE}, func(i []IAST, s string) IAST {
		return i[0]
	})
}
