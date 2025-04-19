package sintax

import (
	. "hulk.com/app/grammar"
)

var HulkGrammar IGrammar

func init() {
	ArithmeticExpr := NewGrammarSymbol("ArithmeticExpr", NonTerminal, false)
	PlusMinusTerm := NewGrammarSymbol("PlusMinusTerm", NonTerminal, false)
	MulDivTerm := NewGrammarSymbol("MulDivTerm", NonTerminal, false)
	LP := NewGrammarSymbol("(", Terminal, false)
	RP := NewGrammarSymbol(")", Terminal, false)
	Number := NewGrammarSymbol("number", Terminal, false)
	Plus := NewGrammarSymbol("+", Terminal, false)
	Minus := NewGrammarSymbol("-", Terminal, false)
	Mul := NewGrammarSymbol("*", Terminal, false)
	Div := NewGrammarSymbol("/", Terminal, false)

	HulkGrammar = NewGrammar(ArithmeticExpr)

	// E -> E + T | E - T | T
	HulkGrammar.AddProduction(ArithmeticExpr, []IGrammarSymbol{ArithmeticExpr, Plus, PlusMinusTerm})
	HulkGrammar.AddProduction(ArithmeticExpr, []IGrammarSymbol{ArithmeticExpr, Minus, PlusMinusTerm})
	HulkGrammar.AddProduction(ArithmeticExpr, []IGrammarSymbol{PlusMinusTerm})
	// T -> T * F | T / F | F
	HulkGrammar.AddProduction(PlusMinusTerm, []IGrammarSymbol{PlusMinusTerm, Mul, MulDivTerm})
	HulkGrammar.AddProduction(PlusMinusTerm, []IGrammarSymbol{PlusMinusTerm, Div, MulDivTerm})
	HulkGrammar.AddProduction(PlusMinusTerm, []IGrammarSymbol{MulDivTerm})
	// F -> ( E ) | number
	HulkGrammar.AddProduction(MulDivTerm, []IGrammarSymbol{LP, ArithmeticExpr, RP})
	HulkGrammar.AddProduction(MulDivTerm, []IGrammarSymbol{Number})
}
