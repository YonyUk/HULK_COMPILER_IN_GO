package sintax

import (
	. "hulk.com/app/grammar"
)

var ArithMeticGrammar IGrammar

func init() {
	ArithmeticExpr := NewGrammarSymbol("ArithmeticExpr", NonTerminal, false)
	PlusMinusTerm := NewGrammarSymbol("PlusMinusTerm", NonTerminal, false)
	MulDivTerm := NewGrammarSymbol("MulDivTerm", NonTerminal, false)
	ExpTerm := NewGrammarSymbol("ExpTerm", NonTerminal, false)
	LP := NewGrammarSymbol("(", Terminal, false)
	RP := NewGrammarSymbol(")", Terminal, false)
	Number := NewGrammarSymbol("number", Terminal, false)
	Plus := NewGrammarSymbol("+", Terminal, false)
	Minus := NewGrammarSymbol("-", Terminal, false)
	Mod := NewGrammarSymbol("%", Terminal, false)
	Mul := NewGrammarSymbol("*", Terminal, false)
	Div := NewGrammarSymbol("/", Terminal, false)
	Exp := NewGrammarSymbol("^", Terminal, false)

	ArithMeticGrammar = NewGrammar(ArithmeticExpr)

	// E -> E + T | E - T | E % T | T
	ArithMeticGrammar.AddProduction(ArithmeticExpr, []IGrammarSymbol{ArithmeticExpr, Plus, PlusMinusTerm})
	ArithMeticGrammar.AddProduction(ArithmeticExpr, []IGrammarSymbol{ArithmeticExpr, Minus, PlusMinusTerm})
	ArithMeticGrammar.AddProduction(ArithmeticExpr, []IGrammarSymbol{ArithmeticExpr, Mod, PlusMinusTerm})
	ArithMeticGrammar.AddProduction(ArithmeticExpr, []IGrammarSymbol{PlusMinusTerm})
	// T -> T * F | T / F | F
	ArithMeticGrammar.AddProduction(PlusMinusTerm, []IGrammarSymbol{PlusMinusTerm, Mul, MulDivTerm})
	ArithMeticGrammar.AddProduction(PlusMinusTerm, []IGrammarSymbol{PlusMinusTerm, Div, MulDivTerm})
	ArithMeticGrammar.AddProduction(PlusMinusTerm, []IGrammarSymbol{MulDivTerm})
	// // F -> F ^ Q | Q
	ArithMeticGrammar.AddProduction(MulDivTerm, []IGrammarSymbol{MulDivTerm, Exp, ExpTerm})
	ArithMeticGrammar.AddProduction(MulDivTerm, []IGrammarSymbol{ExpTerm})
	// Q -> ( E ) | number
	ArithMeticGrammar.AddProduction(ExpTerm, []IGrammarSymbol{LP, ArithmeticExpr, RP})
	ArithMeticGrammar.AddProduction(ExpTerm, []IGrammarSymbol{Number})
}
