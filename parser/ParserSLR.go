package parser

import (
	"strconv"

	. "hulk.com/app/ast"
	. "hulk.com/app/automaton"
	. "hulk.com/app/compiler"
	. "hulk.com/app/grammar"
	. "hulk.com/app/tokens"
	. "hulk.com/app/tools"
)

type ParserSLR struct {
	states              []IState[string]
	start_state         IState[string]
	states_stack        []IState[string]
	action              map[string]map[string]ActionStruct
	reduce              map[string]map[string]ReduceStruct
	stack               []IAST
	ast_engine          func(token IToken, endmarker string) IAST
	reduction_functions map[string]func(asts []IAST) IAST
	endmarker           string
	terminals           []string
}

func NewParserSLRFromGrammar(g IGrammar, endmarker IGrammarSymbol, ast_engine func(token IToken, endmarker string) IAST) *ParserSLR {
	G := AugmentGrammar(g)
	G.MakeFirstsAndFollows(endmarker)

	action := make(map[string]map[string]ActionStruct)
	reduce := make(map[string]map[string]ReduceStruct)
	states := GetCanonicalLR0Collection(G)
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
			goto_ := GOTO(states[i], nt, G)
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
	return &ParserSLR{
		states:              states_,
		start_state:         states_[start_state_index],
		states_stack:        []IState[string]{states_[start_state_index]},
		stack:               []IAST{},
		action:              action,
		reduce:              reduce,
		ast_engine:          ast_engine,
		reduction_functions: make(map[string]func(asts []IAST) IAST),
		endmarker:           endmarker.Symbol(),
		terminals:           Map(G.Terminals(), func(symbol IGrammarSymbol) string { return symbol.Symbol() }),
	}
}

func (parser *ParserSLR) EndMarker() string {
	return parser.endmarker
}

func (parser *ParserSLR) Parse(token IToken, collector IErrorCollector) {
	ast := parser.ast_engine(token, parser.endmarker)
	if _, ok := parser.action[parser.states_stack[len(parser.states_stack)-1].ID()][ast.Symbol()]; !ok {
		msg := "expected "
		for k, _ := range parser.action[parser.states_stack[len(parser.states_stack)-1].ID()] {
			if _, err := IndexOf(parser.terminals, func(s string) bool { return s == k }); err == nil {
				msg += k + ","
			}
		}
		if ast.Symbol() == parser.endmarker {
			collector.AddError(NewError("Unexpected EOF symbol "+","+msg, ast.Line(), ast.Column(), Gramatical))
		} else {
			collector.AddError(NewError("Unexpected symbol "+ast.Symbol()+","+msg, ast.Line(), ast.Column(), Gramatical))
		}
	} else {
		symbol_accepted := false
		for !symbol_accepted {
			action := parser.action[parser.states_stack[len(parser.states_stack)-1].ID()][ast.Symbol()]
			if action.Action == SHIFT || action.Action == ACCEPT {
				parser.states_stack = append(parser.states_stack, parser.states_stack[len(parser.states_stack)-1].Next(ast.Symbol()))
				parser.stack = append(parser.stack, ast)
				symbol_accepted = true
			}
			if action.Action == REDUCE {
				reduction := parser.reduce[parser.states_stack[len(parser.states_stack)-1].ID()][ast.Symbol()]
				reduction_id := reduction.NewSymbol + "->"
				for _, symbol := range reduction.Symbols {
					reduction_id += symbol
				}
				symbols_to_reduce := parser.stack[len(parser.stack)-len(reduction.Symbols):]
				parser.stack = parser.stack[:len(parser.stack)-len(reduction.Symbols)]
				reductor := parser.reduction_functions[reduction_id]
				new_ast := reductor(symbols_to_reduce)
				parser.stack = append(parser.stack, new_ast)
				parser.states_stack = parser.states_stack[:len(parser.states_stack)-len(reduction.Symbols)]
				action = parser.action[parser.states_stack[len(parser.states_stack)-1].ID()][new_ast.Symbol()]
				parser.states_stack = append(parser.states_stack, parser.states_stack[len(parser.states_stack)-1].Next(new_ast.Symbol()))
			}
		}
	}
}

func (parser *ParserSLR) ActionTable() map[string]map[string]ActionStruct {
	return parser.action
}

func (parser *ParserSLR) ReduceTable() map[string]map[string]ReduceStruct {
	return parser.reduce
}

func (parser *ParserSLR) StartState() string {
	return parser.start_state.ID()
}

func (parser *ParserSLR) SetReduction(reduction_id string, reductor func(asts []IAST) IAST) {
	parser.reduction_functions[reduction_id] = reductor
}

func (parser *ParserSLR) AST() IAST {
	return parser.stack[0]
}
