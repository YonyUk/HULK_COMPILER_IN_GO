package hulk

import (
	. "hulk.com/app/grammar"
)

var NumberGrammar IGrammar

// Init the grammar for the number tokens of hulk language
func init() {
	SignedNumber := NewGrammarSymbol("SignedNumber", NonTerminal, false)
	Number := NewGrammarSymbol("Number", NonTerminal, false)
	before_tail := NewGrammarSymbol("before_tail", NonTerminal, false)
	after_tail := NewGrammarSymbol("after_tail", NonTerminal, false)
	e_notation := NewGrammarSymbol("e_notation", NonTerminal, false)
	dot_notation := NewGrammarSymbol("dot_notation", NonTerminal, false)
	d0 := NewGrammarSymbol("0", Terminal, false)
	d1 := NewGrammarSymbol("1", Terminal, false)
	d2 := NewGrammarSymbol("2", Terminal, false)
	d3 := NewGrammarSymbol("3", Terminal, false)
	d4 := NewGrammarSymbol("4", Terminal, false)
	d5 := NewGrammarSymbol("5", Terminal, false)
	d6 := NewGrammarSymbol("6", Terminal, false)
	d7 := NewGrammarSymbol("7", Terminal, false)
	d8 := NewGrammarSymbol("8", Terminal, false)
	d9 := NewGrammarSymbol("9", Terminal, false)
	dot := NewGrammarSymbol(".", Terminal, false)
	e := NewGrammarSymbol("e", Terminal, false)
	e_plus := NewGrammarSymbol("+", Terminal, false)
	e_minus := NewGrammarSymbol("-", Terminal, false)
	plus := NewGrammarSymbol("+", Terminal, false)
	minus := NewGrammarSymbol("-", Terminal, false)
	epsilon := NewGrammarSymbol("epsilon", Terminal, true)

	// Numbers grammar
	NumberGrammar = NewGrammar(SignedNumber)

	// SignedNumber productions
	NumberGrammar.AddProduction(SignedNumber, []IGrammarSymbol{plus, Number})
	NumberGrammar.AddProduction(SignedNumber, []IGrammarSymbol{minus, Number})
	NumberGrammar.AddProduction(SignedNumber, []IGrammarSymbol{d1, before_tail})
	NumberGrammar.AddProduction(SignedNumber, []IGrammarSymbol{d2, before_tail})
	NumberGrammar.AddProduction(SignedNumber, []IGrammarSymbol{d3, before_tail})
	NumberGrammar.AddProduction(SignedNumber, []IGrammarSymbol{d4, before_tail})
	NumberGrammar.AddProduction(SignedNumber, []IGrammarSymbol{d5, before_tail})
	NumberGrammar.AddProduction(SignedNumber, []IGrammarSymbol{d6, before_tail})
	NumberGrammar.AddProduction(SignedNumber, []IGrammarSymbol{d7, before_tail})
	NumberGrammar.AddProduction(SignedNumber, []IGrammarSymbol{d8, before_tail})
	NumberGrammar.AddProduction(SignedNumber, []IGrammarSymbol{d9, before_tail})
	NumberGrammar.AddProduction(SignedNumber, []IGrammarSymbol{d0})

	// Number productions
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d1, before_tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d2, before_tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d3, before_tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d4, before_tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d5, before_tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d6, before_tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d7, before_tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d8, before_tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d9, before_tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d0})

	// before_tail productions
	NumberGrammar.AddProduction(before_tail, []IGrammarSymbol{d0, before_tail})
	NumberGrammar.AddProduction(before_tail, []IGrammarSymbol{d1, before_tail})
	NumberGrammar.AddProduction(before_tail, []IGrammarSymbol{d2, before_tail})
	NumberGrammar.AddProduction(before_tail, []IGrammarSymbol{d3, before_tail})
	NumberGrammar.AddProduction(before_tail, []IGrammarSymbol{d4, before_tail})
	NumberGrammar.AddProduction(before_tail, []IGrammarSymbol{d5, before_tail})
	NumberGrammar.AddProduction(before_tail, []IGrammarSymbol{d6, before_tail})
	NumberGrammar.AddProduction(before_tail, []IGrammarSymbol{d7, before_tail})
	NumberGrammar.AddProduction(before_tail, []IGrammarSymbol{d8, before_tail})
	NumberGrammar.AddProduction(before_tail, []IGrammarSymbol{d9, before_tail})
	NumberGrammar.AddProduction(before_tail, []IGrammarSymbol{epsilon})
	NumberGrammar.AddProduction(before_tail, []IGrammarSymbol{dot, after_tail})
	NumberGrammar.AddProduction(before_tail, []IGrammarSymbol{e, e_notation})

	// dot productions
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d0, dot_notation})
	NumberGrammar.AddProduction(SignedNumber, []IGrammarSymbol{d0, dot_notation})
	NumberGrammar.AddProduction(dot_notation, []IGrammarSymbol{dot, after_tail})

	// after_tail productions
	NumberGrammar.AddProduction(after_tail, []IGrammarSymbol{d0, after_tail})
	NumberGrammar.AddProduction(after_tail, []IGrammarSymbol{d1, after_tail})
	NumberGrammar.AddProduction(after_tail, []IGrammarSymbol{d2, after_tail})
	NumberGrammar.AddProduction(after_tail, []IGrammarSymbol{d3, after_tail})
	NumberGrammar.AddProduction(after_tail, []IGrammarSymbol{d4, after_tail})
	NumberGrammar.AddProduction(after_tail, []IGrammarSymbol{d5, after_tail})
	NumberGrammar.AddProduction(after_tail, []IGrammarSymbol{d6, after_tail})
	NumberGrammar.AddProduction(after_tail, []IGrammarSymbol{d7, after_tail})
	NumberGrammar.AddProduction(after_tail, []IGrammarSymbol{d8, after_tail})
	NumberGrammar.AddProduction(after_tail, []IGrammarSymbol{d9, after_tail})
	NumberGrammar.AddProduction(after_tail, []IGrammarSymbol{epsilon})

	// e_notation productions
	NumberGrammar.AddProduction(e_notation, []IGrammarSymbol{e_plus, after_tail})
	NumberGrammar.AddProduction(e_notation, []IGrammarSymbol{e_minus, after_tail})
}
