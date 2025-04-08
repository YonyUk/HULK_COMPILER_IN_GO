package grammar

import (
	. "hulk.com/app/ast"
)

// Attributed grammar interface
type IAttributedGrammar interface {
	// Terminals of this grammar
	Terminals() []IGrammarSymbol
	// Nonterminals for this grammar
	NonTerminals() []IGrammarSymbol
	// The start symbol of this grammar
	StartSymbol() IGrammarSymbol
	// Adds a new production to this grammar
	AddProduction(symbol IGrammarSymbol, symbols []IGrammarSymbol, rule func(asts []IAST) IAST)
}
