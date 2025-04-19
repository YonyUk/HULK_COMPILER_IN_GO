package grammar

func GrammarUnion(g1 IGrammar, g2 IGrammar) IGrammar {
	g_result := NewGrammar(g1.StartSymbol())
	for _, nt := range g1.NonTerminals() {
		for _, production := range g1.GetProductions(nt) {
			g_result.AddProduction(nt, production)
		}
	}
	for _, nt := range g2.NonTerminals() {
		for _, production := range g2.GetProductions(nt) {
			g_result.AddProduction(nt, production)
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
