package grammar

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/tools"
)

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

func AugmentAttributedGrammar(g IAttributedGrammar) IAttributedGrammar {
	new_start_symbol := NewGrammarSymbol(g.StartSymbol().Symbol()+"_new_start", NonTerminal, false)
	g_result := NewAttributedGrammar(new_start_symbol)
	g_result.AddProduction(new_start_symbol, []IGrammarSymbol{g.StartSymbol()}, func(asts []IAST, new_symbol string) IAST {
		asts[0].UpdateSymbol(new_symbol)
		return asts[0]
	})
	for _, nt := range g.NonTerminals() {
		for _, production := range g.GetProductions(nt) {
			production_id := g.GetProductionId(nt.Symbol(), Map(production, func(s IGrammarSymbol) string { return s.Symbol() }))
			rule, _ := g.GetProductionRule(production_id)
			g_result.AddProduction(nt, production, rule)
		}
	}
	return g_result
}
