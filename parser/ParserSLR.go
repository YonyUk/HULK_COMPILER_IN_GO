package parser

import (
	"strconv"

	. "hulk.com/app/ast"
	. "hulk.com/app/automaton"
	. "hulk.com/app/grammar"
	. "hulk.com/app/tools"
)

type ParserSLR struct {
	automaton IAutomaton[string]
	action    map[string]map[string]ParserAction
	reduce    map[string]map[string][]string
	stack     []IAST
}

func NewParserSLRFromGrammar(g IGrammar, endmarker IGrammarSymbol) *ParserSLR {
	G := AugmentGrammar(g)
	G.MakeFirstsAndFollows(endmarker)

	action := make(map[string]map[string]ParserAction)
	reduce := make(map[string]map[string][]string)

	states := GetCanonicalLR0Collection(G)

	// for i := 0; i < len(states); i++ {
	// 	fmt.Println("I" + strconv.Itoa(i))
	// 	ShowCollection(states[i])
	// }
	pos := -1
	states_ := Map(states, func(col IItemLR0Collection) IState[string] {
		pos++
		return NewState[string]("I"+strconv.Itoa(pos), true, false)
	})
	start_state_index, _ := IndexOf(states, func(collection IItemLR0Collection) bool {
		for _, item := range collection.Items() {
			if len(item.LeftTail()) > 0 {
				return false
			}
		}
		return true
	})
	for i := 0; i < len(states); i++ {
		for _, item := range states[i].Items() {
			if len(item.RightTail()) > 0 {
				for _, terminal := range G.Terminals() {
					goto_ := GOTO(states[i], terminal, G)
					if goto_ != nil {
						next_index, _ := IndexOf(states, func(coll IItemLR0Collection) bool { return coll.ID() == goto_.ID() })
						states_[i].AddTransition(terminal.Symbol(), states_[next_index])
						if _, ok := action[states_[i].ID()]; !ok {
							action[states_[i].ID()] = make(map[string]ParserAction)
						}
						action[states_[i].ID()][terminal.Symbol()] = SHIFT
					}
				}
			}
			if len(item.RightTail()) == 0 && item.Head().Symbol() != G.StartSymbol().Symbol() {
				for _, terminal := range G.FOLLOW(item.Head()) {
					if _, ok := action[states_[i].ID()]; !ok {
						action[states_[i].ID()] = make(map[string]ParserAction)
					}
					action[states_[i].ID()][terminal.Symbol()] = REDUCE
					if _, ok := reduce[states_[i].ID()]; !ok {
						reduce[states_[i].ID()] = make(map[string][]string)
					}
					reduce[states_[i].ID()][terminal.Symbol()] = Map(append([]IGrammarSymbol{item.Head()}, item.LeftTail()...), func(s IGrammarSymbol) string { return s.Symbol() })
				}
			}
			if len(item.RightTail()) == 0 && item.Head().Symbol() == G.StartSymbol().Symbol() {
				if _, ok := action[states_[i].ID()]; !ok {
					action[states_[i].ID()] = make(map[string]ParserAction)
				}
				action[states_[i].ID()][endmarker.Symbol()] = ACCEPT
			}
		}
	}
	internal_automaton := NewAutomaton(states_[start_state_index], states_, Map(G.Terminals(), func(s IGrammarSymbol) string { return s.Symbol() }))
	// for k1, val := range action {
	// 	for k2, to := range val {
	// 		if to == SHIFT {
	// 			index, _ := strconv.ParseInt(k1[1:], 10, 32)
	// 			next := states_[index].Next(k2)
	// 			fmt.Println("ACTION[", k1, ",", k2, "] = SHIFT", next.ID())
	// 		}
	// 		if to == REDUCE {
	// 			fmt.Println("ACTION[", k1, ",", k2, "] = REDUCE", reduce[k1][k2])
	// 		}
	// 		if to == ACCEPT {
	// 			fmt.Println("ACTION[", k1, ",", k2, "] = ACCEPT")
	// 		}
	// 	}
	// }
	return &ParserSLR{
		automaton: internal_automaton,
		stack:     []IAST{},
		action:    action,
		reduce:    reduce,
	}
}
