package grammar

func GrammarUnion(grammars []IGrammar, start_symbol_id string) IGrammar {
	if len(grammars) < 2 {
		panic("grammars most have at least two elements")
	}
	start_symbol := NewGrammarSymbol(start_symbol_id, NonTerminal, false)
	g_result := NewGrammar(start_symbol)
	for _, g := range grammars {
		g_result.AddProduction(start_symbol, []IGrammarSymbol{g.StartSymbol()})
		for _, nt := range g.NonTerminals() {
			for _, production := range g.GetProductions(nt) {
				g_result.AddProduction(nt, production)
			}
		}
	}
	return g_result
}

func AugmentGrammar(g IGrammar) IGrammar {
	new_start_symbol := NewGrammarSymbol(g.StartSymbol().Symbol()+"_new_start", NonTerminal, false)
	g_result := NewGrammar(new_start_symbol)
	g_result.AddProduction(new_start_symbol, []IGrammarSymbol{g.StartSymbol()})
	for _, nt := range g.NonTerminals() {
		for _, production := range g.GetProductions(nt) {
			g_result.AddProduction(nt, production)
		}
	}
	return g_result
}
