package grammar

import (
	"errors"

	. "hulk.com/app/ast"
	. "hulk.com/app/tools"
)

type AttributedGrammar struct {
	IGrammar
	rules map[string]func([]IAST, string) IAST
}

func (g *AttributedGrammar) GetProductionId(head string, production []string) string {
	id := head + "--->"
	for _, symbol_ := range production {
		id += symbol_
	}
	return id
}

func NewAttributedGrammar(start IGrammarSymbol) *AttributedGrammar {
	return &AttributedGrammar{
		rules:    make(map[string]func([]IAST, string) IAST),
		IGrammar: NewGrammar(start),
	}
}

func (g *AttributedGrammar) Rules() map[string]func([]IAST, string) IAST {
	return g.rules
}

func (g *AttributedGrammar) AddProduction(symbol IGrammarSymbol, production []IGrammarSymbol, rule func([]IAST, string) IAST) error {
	err := g.IGrammar.AddProduction(symbol, production)
	g.rules[g.GetProductionId(symbol.Symbol(), Map(production, func(s IGrammarSymbol) string { return s.Symbol() }))] = rule
	return err
}

func (g *AttributedGrammar) GetProductionRule(production_id string) (func([]IAST, string) IAST, error) {
	if reduction, ok := g.rules[production_id]; ok {
		return reduction, nil
	}
	return nil, errors.New("Not production found")
}
