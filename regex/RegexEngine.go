package regex

import (
	"errors"

	. "hulk.com/app/automaton"
	. "hulk.com/app/grammar"
	. "hulk.com/app/tools"
)

type RegexEngine struct {
}

func NewRegexEngine() *RegexEngine {
	return &RegexEngine{}
}

func (re *RegexEngine) Regex(g IGrammar) (IAutomaton[rune], error) {
	states := []IState[rune]{}
	for _, symbol := range g.NonTerminals() {
		states = append(states, NewState[rune](symbol.Symbol(), false, false))
	}
	states = append(states, NewState[rune]("FINAL", true, false))
	for _, symbol := range g.NonTerminals() {
		productions := g.GetProductions(symbol)
		state_index, _ := IndexOf(states, func(s IState[rune]) bool { return s.ID() == symbol.Symbol() })
		for _, production := range productions {
			if len(production) > 2 || len(production) < 1 {
				return nil, errors.New("This grammar is not regular, bad production at symbol: " + symbol.Symbol())
			}
			switch len(production) {
			case 2:
				if production[0].Type() == production[1].Type() {
					return nil, errors.New("This grammar is not regular")
				}
				switch production[0].Type() {

				case Terminal:
					if production[0].Epsilon() {
						return nil, errors.New("This grammar is not regular, bad production at symbol: " + symbol.Symbol())
					}
					trans_state_index, _ := IndexOf(states, func(s IState[rune]) bool { return s.ID() == production[1].Symbol() })
					if states[state_index].HasTransition(rune(production[0].Symbol()[0])) {
						new_state_index, _ := IndexOf(states, func(s IState[rune]) bool {
							return s.ID() == states[state_index].Next(rune(production[0].Symbol()[0])).ID()
						})
						states[new_state_index].Epsilon(states[trans_state_index])
					} else {
						states[state_index].AddTransition(rune(production[0].Symbol()[0]), states[trans_state_index])
					}
				case NonTerminal:
					if production[1].Epsilon() {
						return nil, errors.New("Bad formed grammar")
					}
					trans_state_index, _ := IndexOf(states, func(s IState[rune]) bool { return s.ID() == production[0].Symbol() })
					states[state_index].Epsilon(states[trans_state_index])
					states[trans_state_index].AddTransition(rune(production[1].Symbol()[0]), states[len(states)-1])
				default:
					return nil, errors.New("Symbol type not defined")
				}
			case 1:
				switch production[0].Type() {
				case Terminal:
					switch production[0].Epsilon() {
					case true:
						states[state_index].Epsilon(states[len(states)-1])
					default:
						states[state_index].AddTransition(rune(production[0].Symbol()[0]), states[len(states)-1])
					}
				default:
					return nil, errors.New("This grammar is not regular")
				}
			}
		}
	}
	alphabet := Map(g.Terminals(), func(symbol IGrammarSymbol) rune { return rune(symbol.Symbol()[0]) })
	start_state_index, _ := IndexOf(states, func(s IState[rune]) bool { return s.ID() == g.StartSymbol().Symbol() })
	aut := NewAutomaton(states[start_state_index], states, alphabet)
	return aut.ToDeterministic(), nil
}
