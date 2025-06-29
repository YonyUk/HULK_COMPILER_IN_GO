package grammarpackage

func GrammarCode() string {
	return `package grammar

import (
	"errors"

	. "hulk.com/app/tools"
)

type Grammar struct {
	terminals     []IGrammarSymbol
	non_terminals []IGrammarSymbol
	productions   map[string][][]IGrammarSymbol
	firsts        map[string][]IGrammarSymbol
	follows       map[string][]IGrammarSymbol
}

func NewGrammar(start IGrammarSymbol) *Grammar {
	return &Grammar{
		terminals:     []IGrammarSymbol{},
		non_terminals: []IGrammarSymbol{start},
		productions:   make(map[string][][]IGrammarSymbol),
		firsts:        make(map[string][]IGrammarSymbol),
		follows:       make(map[string][]IGrammarSymbol),
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

func (grammar *Grammar) MakeFirstsAndFollows(endmarker IGrammarSymbol) {
	grammar.make_firsts()
	grammar.make_follows(endmarker)
}

func (grammar *Grammar) FIRST(symbols []IGrammarSymbol) []IGrammarSymbol {
	epsilon := len(symbols) > 0
	position := 0
	result := []IGrammarSymbol{}
	for epsilon && position < len(symbols) {
		for _, symbol := range grammar.firsts[symbols[position].Symbol()] {
			_, err := IndexOf(result, func(s IGrammarSymbol) bool { return s.Symbol() == symbol.Symbol() })
			if err != nil && !symbol.Epsilon() {
				result = append(result, symbol)
			}
		}
		epsilon = grammar.derive_in_epsilon(symbols[position])
		position++
	}
	if epsilon {
		index, _ := IndexOf(grammar.terminals, func(s IGrammarSymbol) bool { return s.Epsilon() })
		result = append(result, grammar.terminals[index])
	}
	return result
}

func (grammar *Grammar) FOLLOW(symbol IGrammarSymbol) []IGrammarSymbol {
	return grammar.follows[symbol.Symbol()]
}

func (grammar *Grammar) init_firsts_sets() {
	symbols := []IGrammarSymbol{}
	for _, terminal := range grammar.terminals {
		symbols = append(symbols, terminal)
	}
	for _, non_terminal := range grammar.non_terminals {
		symbols = append(symbols, non_terminal)
	}
	for _, symbol := range symbols {
		if symbol.Type() == Terminal {
			grammar.firsts[symbol.Symbol()] = []IGrammarSymbol{symbol}
			continue
		}
		if symbol.Type() == NonTerminal && grammar.derive_in_epsilon(symbol) {
			index, _ := IndexOf(grammar.terminals, func(s IGrammarSymbol) bool { return s.Epsilon() })
			grammar.firsts[symbol.Symbol()] = append(grammar.firsts[symbol.Symbol()], grammar.terminals[index])
			continue
		}
		grammar.firsts[symbol.Symbol()] = []IGrammarSymbol{}
	}
}

func (grammar *Grammar) make_firsts() {
	grammar.init_firsts_sets()
	change := true
	for change {
		change = false
		for _, symbol := range grammar.non_terminals {
			for _, production := range grammar.productions[symbol.Symbol()] {
				epsilon := true
				position := 0
				for epsilon && position < len(production) {
					for _, sym := range grammar.firsts[production[position].Symbol()] {
						_, err := IndexOf(grammar.firsts[symbol.Symbol()], func(s IGrammarSymbol) bool { return sym.Symbol() == s.Symbol() })
						if err != nil && !sym.Epsilon() {
							grammar.firsts[symbol.Symbol()] = append(grammar.firsts[symbol.Symbol()], sym)
							change = true
						}
					}
					epsilon = grammar.derive_in_epsilon(production[position])
					position++
				}
				if epsilon {
					ep_index, _ := IndexOf(grammar.terminals, func(s IGrammarSymbol) bool { return s.Epsilon() })
					grammar.firsts[symbol.Symbol()] = append(grammar.firsts[symbol.Symbol()], grammar.terminals[ep_index])
					change = true
				}
			}
			if change {
				break
			}
		}
	}
}

func (grammar *Grammar) init_follows_sets(endmarker IGrammarSymbol) {
	for _, symbol := range grammar.non_terminals {
		grammar.follows[symbol.Symbol()] = []IGrammarSymbol{}
	}
	grammar.follows[grammar.StartSymbol().Symbol()] = append(grammar.follows[grammar.StartSymbol().Symbol()], endmarker)
}

func (grammar *Grammar) make_follow_for(symbol IGrammarSymbol) bool {
	result := false
	for _, non_terminal := range grammar.non_terminals {
		for _, production := range grammar.productions[non_terminal.Symbol()] {
			index, err := IndexOf(production, func(s IGrammarSymbol) bool { return s.Symbol() == symbol.Symbol() })
			if err == nil {
				first := grammar.FIRST(production[index+1:])
				_, ep_err := IndexOf(first, func(s IGrammarSymbol) bool { return s.Epsilon() })
				if index == len(production)-1 || ep_err == nil {
					for _, sym := range grammar.follows[non_terminal.Symbol()] {
						_, err1 := IndexOf(grammar.follows[symbol.Symbol()], func(s IGrammarSymbol) bool { return s.Symbol() == sym.Symbol() })
						if err1 != nil && !sym.Epsilon() {
							grammar.follows[symbol.Symbol()] = append(grammar.follows[symbol.Symbol()], sym)
							result = true
						}
					}
				}
				for _, sym := range first {
					_, err1 := IndexOf(grammar.follows[symbol.Symbol()], func(s IGrammarSymbol) bool { return s.Symbol() == sym.Symbol() })
					if err1 != nil && !sym.Epsilon() {
						grammar.follows[symbol.Symbol()] = append(grammar.follows[symbol.Symbol()], sym)
						result = true
					}
				}
			}
		}
	}
	return result
}

func (grammar *Grammar) make_follows(endmarker IGrammarSymbol) {
	grammar.init_follows_sets(endmarker)
	change := true
	for change {
		change = false
		for _, non_terminal := range grammar.non_terminals {
			change = grammar.make_follow_for(non_terminal)
		}
	}
}

func (grammar *Grammar) derive_in_epsilon(symbol IGrammarSymbol) bool {
	for _, production := range grammar.GetProductions(symbol) {
		if len(production) == 1 && production[0].Epsilon() {
			return true
		}
	}
	return false
}
`
}
