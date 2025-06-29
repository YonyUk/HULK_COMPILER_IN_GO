package grammarpackage

func IAttributedGrammarCode() string {
	return `package grammar

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
		AddProduction(symbol IGrammarSymbol, symbols []IGrammarSymbol, rule func([]IAST, string) IAST) error
		// Gets the productions of one symbol
		GetProductions(symbol IGrammarSymbol) [][]IGrammarSymbol
		// Gets the FIRST of the given symbol
		FIRST(symbols []IGrammarSymbol) []IGrammarSymbol
		// Gets the FOLLOW of the given symbol
		FOLLOW(symbol IGrammarSymbol) []IGrammarSymbol
		// Computes the firsts and follows for the grammar
		MakeFirstsAndFollows(endmarker IGrammarSymbol)
		// gets the reduction rule
		GetProductionRule(production_id string) (func([]IAST, string) IAST, error)
		// gets the production id
		GetProductionId(head string, production []string) string
		// Gets the productions rules
		Rules() map[string]func([]IAST, string) IAST
	}
	`
}
