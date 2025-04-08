package grammar

// Grammar interface
type IGrammar interface {
	// Terminals of this grammar
	Terminals() []IGrammarSymbol
	// Nonterminals for this grammar
	NonTerminals() []IGrammarSymbol
	// The start symbol of this grammar
	StartSymbol() IGrammarSymbol
	// Adds a new production to this grammar
	AddProduction(symbol IGrammarSymbol, symbols []IGrammarSymbol)
}
