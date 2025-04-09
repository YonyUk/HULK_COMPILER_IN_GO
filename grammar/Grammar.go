package grammar

import (
	"errors"

	. "hulk.com/app/tools"
)

type Grammar struct {
	terminals     []IGrammarSymbol
	non_terminals []IGrammarSymbol
	productions   map[IGrammarSymbol][][]IGrammarSymbol
}

func NewGrammar(start IGrammarSymbol) Grammar {
	return Grammar{
		terminals:     []IGrammarSymbol{},
		non_terminals: []IGrammarSymbol{start},
		productions:   make(map[IGrammarSymbol][][]IGrammarSymbol),
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
	return NewGrammarSymbol(grammar.non_terminals[0].Symbol(), grammar.non_terminals[0].Type())
}

func (grammar *Grammar) AddProduction(symbol IGrammarSymbol, symbols []IGrammarSymbol) error {
	if symbol.Type() == Terminal {
		return errors.New("The symbols most be a non terminal")
	}
	if _, ok := grammar.productions[symbol]; !ok {
		grammar.productions[symbol] = [][]IGrammarSymbol{}
	}
	if _, err := IndexOf(grammar.productions[symbol], func(syms []IGrammarSymbol) bool { return CompareArrays(syms, symbols) }); err == nil {
		return errors.New("The given production already exists")
	}
	if _, err := IndexOf(grammar.non_terminals, func(s IGrammarSymbol) bool { return symbol == s }); err != nil {
		grammar.non_terminals = append(grammar.non_terminals, symbol)
	}
	for i := 0; i < len(symbols); i++ {
		switch symbols[i].Type() {
		case NonTerminal:
			if _, err := IndexOf(grammar.non_terminals, func(s IGrammarSymbol) bool { return s == symbols[i] }); err != nil {
				grammar.non_terminals = append(grammar.non_terminals, symbols[i])
			}
		case Terminal:
			if _, err := IndexOf(grammar.terminals, func(s IGrammarSymbol) bool { return s == symbols[i] }); err != nil {
				grammar.terminals = append(grammar.terminals, symbols[i])
			}
		default:
			continue
		}
	}
	grammar.productions[symbol] = append(grammar.productions[symbol], symbols)
	return nil
}
