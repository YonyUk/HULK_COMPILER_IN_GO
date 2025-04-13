package grammar

import (
	"errors"

	. "hulk.com/app/tools"
)

type Grammar struct {
	terminals     []IGrammarSymbol
	non_terminals []IGrammarSymbol
	productions   map[string][][]IGrammarSymbol
}

func NewGrammar(start IGrammarSymbol) *Grammar {
	return &Grammar{
		terminals:     []IGrammarSymbol{},
		non_terminals: []IGrammarSymbol{start},
		productions:   make(map[string][][]IGrammarSymbol),
	}
}

func (grammar *Grammar) Terminals() []IGrammarSymbol {
	result := make([]IGrammarSymbol, len(grammar.terminals))
	copy(result, grammar.terminals)
	return result
}

func (grammar *Grammar) NonTerminals() []IGrammarSymbol {
	result := make([]IGrammarSymbol, len(grammar.non_terminals))
	copy(result, grammar.non_terminals)
	return result
}

func (grammar *Grammar) StartSymbol() IGrammarSymbol {
	return NewGrammarSymbol(grammar.non_terminals[0].Symbol(), grammar.non_terminals[0].Type(), grammar.non_terminals[0].Epsilon())
}

func (grammar *Grammar) AddProduction(symbol IGrammarSymbol, symbols []IGrammarSymbol) error {
	if symbol.Type() == Terminal {
		return errors.New("The symbols most be a non terminal")
	}
	if _, ok := grammar.productions[symbol.Symbol()]; !ok {
		grammar.productions[symbol.Symbol()] = [][]IGrammarSymbol{}
	}
	if _, err := IndexOf(grammar.productions[symbol.Symbol()], func(syms []IGrammarSymbol) bool { return CompareArrays(syms, symbols) }); err == nil {
		return errors.New("The given production already exists")
	}
	if _, err := IndexOf(grammar.non_terminals, func(s IGrammarSymbol) bool { return symbol.Symbol() == s.Symbol() }); err != nil {
		grammar.non_terminals = append(grammar.non_terminals, symbol)
	}
	for i := 0; i < len(symbols); i++ {
		switch symbols[i].Type() {
		case NonTerminal:
			if _, err := IndexOf(grammar.non_terminals, func(s IGrammarSymbol) bool { return s.Symbol() == symbols[i].Symbol() }); err != nil {
				grammar.non_terminals = append(grammar.non_terminals, symbols[i])
			}
		case Terminal:
			if _, err := IndexOf(grammar.terminals, func(s IGrammarSymbol) bool { return s.Symbol() == symbols[i].Symbol() }); err != nil {
				grammar.terminals = append(grammar.terminals, symbols[i])
			}
		default:
			continue
		}
	}
	grammar.productions[symbol.Symbol()] = append(grammar.productions[symbol.Symbol()], symbols)
	return nil
}

func (grammar *Grammar) GetProductions(symbol IGrammarSymbol) [][]IGrammarSymbol {
	if symbol.Type() != NonTerminal {
		return [][]IGrammarSymbol{}
	}
	return grammar.productions[symbol.Symbol()]
}
