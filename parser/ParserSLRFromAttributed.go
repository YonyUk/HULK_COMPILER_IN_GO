package parser

import (
	"fmt"
	"strconv"

	. "hulk.com/app/ast"
	. "hulk.com/app/automaton"
	. "hulk.com/app/compiler"
	. "hulk.com/app/grammar"
	. "hulk.com/app/tokens"
	. "hulk.com/app/tools"
)

type ParserSLRFromAttributedGrammar struct {
	ParserSLR
	reduction_id_engine func(symbol string, production []string) string
}

func NewParserSLRFromAttributedGrammar(g IAttributedGrammar, endmarker IGrammarSymbol, ast_engine func(token IToken, endmarker string) IAST) *ParserSLRFromAttributedGrammar {
	G := AugmentAttributedGrammar(g)
	G.MakeFirstsAndFollows(endmarker)

	action := make(map[string]map[string]ActionStruct)
	reduce := make(map[string]map[string]ReduceStruct)
	states := GetCanonicalLR0CollectionFromAttributed(G)
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
					goto_ := GOTOFromAttributed(states[i], terminal, G)
					if goto_ != nil {
						next_index, _ := IndexOf(states, func(coll IItemLR0Collection) bool { return coll.ID() == goto_.ID() })
						states_[i].AddTransition(terminal.Symbol(), states_[next_index])
						if _, ok := action[states_[i].ID()]; !ok {
							action[states_[i].ID()] = make(map[string]ActionStruct)
						}
						if action_, ok := action[states_[i].ID()][terminal.Symbol()]; ok {
							if action_.Action != action[states_[i].ID()][terminal.Symbol()].Action {
								panic("The given grammar is not SLR(1)")
							}
						}
						action[states_[i].ID()][terminal.Symbol()] = ActionStruct{Action: SHIFT, NextState: states_[next_index].ID()}
					}
				}
			}
			if len(item.RightTail()) == 0 && item.Head().Symbol() != G.StartSymbol().Symbol() {
				for _, terminal := range G.FOLLOW(item.Head()) {
					if _, ok := action[states_[i].ID()]; !ok {
						action[states_[i].ID()] = make(map[string]ActionStruct)
					}
					if action_, ok := action[states_[i].ID()][terminal.Symbol()]; ok {
						if action_.Action != action[states_[i].ID()][terminal.Symbol()].Action {
							panic("The given grammar is not SLR(1)")
						}
					}
					action[states_[i].ID()][terminal.Symbol()] = ActionStruct{Action: REDUCE, NextState: ""}
					if _, ok := reduce[states_[i].ID()]; !ok {
						reduce[states_[i].ID()] = make(map[string]ReduceStruct)
					}
					reduce[states_[i].ID()][terminal.Symbol()] = ReduceStruct{
						NewSymbol: item.Head().Symbol(),
						Symbols:   Map(item.LeftTail(), func(s IGrammarSymbol) string { return s.Symbol() }),
					}
				}
			}
			if len(item.RightTail()) == 0 && item.Head().Symbol() == G.StartSymbol().Symbol() {
				if _, ok := action[states_[i].ID()]; !ok {
					action[states_[i].ID()] = make(map[string]ActionStruct)
				}
				action[states_[i].ID()][endmarker.Symbol()] = ActionStruct{Action: ACCEPT, NextState: ""}
			}
		}
	}
	for i := 0; i < len(states); i++ {
		for _, nt := range G.NonTerminals() {
			goto_ := GOTOFromAttributed(states[i], nt, G)
			if goto_ != nil {
				next_index, _ := IndexOf(states, func(coll IItemLR0Collection) bool { return coll.ID() == goto_.ID() })
				states_[i].AddTransition(nt.Symbol(), states_[next_index])
				if _, ok := action[states_[i].ID()]; !ok {
					action[states_[i].ID()] = make(map[string]ActionStruct)
				}
				if action_, ok := action[states_[i].ID()][nt.Symbol()]; ok {
					if action_.Action != action[states_[i].ID()][nt.Symbol()].Action {
						panic("The given grammar is not SLR(1)")
					}
				}
				action[states_[i].ID()][nt.Symbol()] = ActionStruct{Action: SHIFT, NextState: states_[next_index].ID()}
			}
		}
	}
	return &ParserSLRFromAttributedGrammar{
		ParserSLR: ParserSLR{
			states:              states_,
			start_state:         states_[start_state_index],
			states_stack:        []IState[string]{states_[start_state_index]},
			stack:               []IAST{},
			action:              action,
			reduce:              reduce,
			ast_engine:          ast_engine,
			reduction_functions: G.Rules(),
			endmarker:           endmarker.Symbol(),
			terminals:           Map(G.Terminals(), func(symbol IGrammarSymbol) string { return symbol.Symbol() }),
		},
		reduction_id_engine: G.GetProductionId,
	}
}

func (parser *ParserSLRFromAttributedGrammar) AST() IAST {
	return parser.stack[0]
}

func (parser *ParserSLRFromAttributedGrammar) ActionTable() map[string]map[string]ActionStruct {
	return parser.action
}
func (parser *ParserSLRFromAttributedGrammar) ReduceTable() map[string]map[string]ReduceStruct {
	return parser.reduce
}

func (parser *ParserSLRFromAttributedGrammar) EndMarker() string {
	return parser.endmarker
}

func (parser *ParserSLRFromAttributedGrammar) SetReduction(production_id string, reductor func([]IAST, string) IAST) {
	panic("Invalid Operation")
}

func (parser *ParserSLRFromAttributedGrammar) Parse(token IToken, collector IErrorCollector) {
	ast := parser.ast_engine(token, parser.endmarker)
	if _, ok := parser.action[parser.states_stack[len(parser.states_stack)-1].ID()][ast.Symbol()]; !ok {
		msg := "expected "
		for k, _ := range parser.action[parser.states_stack[len(parser.states_stack)-1].ID()] {
			if _, err := IndexOf(parser.terminals, func(s string) bool { return s == k }); err == nil {
				msg += "'" + k + "'" + ","
			}
		}
		if ast.Symbol() == parser.endmarker {
			collector.AddError(NewError("Unexpected EOF symbol "+","+msg, ast.Line(), ast.Column(), Gramatical))
		} else {
			collector.AddError(NewError("Unexpected symbol '"+ast.Symbol()+"',"+msg, ast.Line(), ast.Column(), Gramatical))
		}
	} else {
		symbol_accepted := false
		for !symbol_accepted {
			fmt.Println(Map(parser.stack, func(a IAST) string { return a.Symbol() }))
			action := parser.action[parser.states_stack[len(parser.states_stack)-1].ID()][ast.Symbol()]
			if action.Action == SHIFT || action.Action == ACCEPT {
				parser.states_stack = append(parser.states_stack, parser.states_stack[len(parser.states_stack)-1].Next(ast.Symbol()))
				parser.stack = append(parser.stack, ast)
				symbol_accepted = true
			}
			if action.Action == REDUCE {
				reduction := parser.reduce[parser.states_stack[len(parser.states_stack)-1].ID()][ast.Symbol()]
				reduction_id := parser.reduction_id_engine(reduction.NewSymbol, reduction.Symbols)
				symbols_to_reduce := parser.stack[len(parser.stack)-len(reduction.Symbols):]
				parser.stack = parser.stack[:len(parser.stack)-len(reduction.Symbols)]
				if reductor, ok := parser.reduction_functions[reduction_id]; ok {
					new_ast := reductor(symbols_to_reduce, reduction.NewSymbol)
					parser.stack = append(parser.stack, new_ast)
					parser.states_stack = parser.states_stack[:len(parser.states_stack)-len(reduction.Symbols)]
					action = parser.action[parser.states_stack[len(parser.states_stack)-1].ID()][new_ast.Symbol()]
					parser.states_stack = append(parser.states_stack, parser.states_stack[len(parser.states_stack)-1].Next(new_ast.Symbol()))
				} else {
					panic("There is not a reduction defined for " + reduction_id)
				}
			}
		}
		fmt.Println(Map(parser.stack, func(a IAST) string { return a.Symbol() }))
	}
}
