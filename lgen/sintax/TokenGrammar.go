package sintax

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/grammar"
	. "hulk.com/app/lgen/grammar_symbols"
	. "hulk.com/app/lgen/lgen_ast"
)

var TokenGrammar IAttributedGrammar

func init() {
	TokenGrammar = NewAttributedGrammar(TOKEN)

	// TOKEN ----> TOKEN_DEFINED_BY_LIST
	TokenGrammar.AddProduction(TOKEN, []IGrammarSymbol{TOKEN_DEFINED_BY_LIST}, func(i []IAST, s string) IAST {
		if TOKEN_DEFINED_BY_LIST.Symbol() != i[0].Symbol() {
			msg := "Unexpected symbol: '" + i[0].Symbol() + "', expected " + TOKEN_DEFINED_BY_LIST.Symbol()
			panic(msg)
		}
		i[0].UpdateSymbol(s)
		return i[0]
	})

	// TOKEN ----> TOKEN_DEFINED_BY_GRAMMAR
	TokenGrammar.AddProduction(TOKEN, []IGrammarSymbol{TOKEN_DEFINED_BY_GRAMMAR}, func(i []IAST, s string) IAST {
		if TOKEN_DEFINED_BY_GRAMMAR.Symbol() != i[0].Symbol() {
			msg := "Unexpected symbol: '" + i[0].Symbol() + "', expected " + TOKEN_DEFINED_BY_GRAMMAR.Symbol()
			panic(msg)
		}
		i[0].UpdateSymbol(s)
		return i[0]
	})

	// TOKEN_DEFINED_BY_GRAMMAR ----> TOKEN_DECLARATION = RIGHT_REGULAR_GRAMMAR
	TokenGrammar.AddProduction(TOKEN_DEFINED_BY_GRAMMAR, []IGrammarSymbol{TOKEN_DECLARATION, Equal, RIGHT_REGULAR_GRAMMAR}, func(i []IAST, s string) IAST {
		if TOKEN_DECLARATION.Symbol() != i[0].Symbol() {
			msg := "Unexpected symbol: '" + i[0].Symbol() + "', expected " + TOKEN_DECLARATION.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != Equal.Symbol() {
			msg := "Unexpected symbol: '" + i[1].Symbol() + "', expected " + Equal.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != RIGHT_REGULAR_GRAMMAR.Symbol() {
			msg := "Unexpected symbol: '" + i[2].Symbol() + "', expected " + RIGHT_REGULAR_GRAMMAR.Symbol()
			panic(msg)
		}
		decl, _ := i[0].(*TOKEN_DECLARATION_DEFINED_BY_GRAMMAR_AST)
		token_grammar_ast, _ := i[2].(*RIGHT_REGULAR_GRAMMAR_AST)
		token_grammar := token_grammar_ast.RegularGrammar
		decl.TokenGrammar = token_grammar
		decl.UpdateSymbol(s)
		return decl
	})

	// TOKEN_DEFINED_BY_LIST ----> TOKEN_DECLARATION = STRING_LIST
	TokenGrammar.AddProduction(TOKEN_DEFINED_BY_LIST, []IGrammarSymbol{TOKEN_DECLARATION, Equal, STRING_LIST}, func(i []IAST, s string) IAST {
		if TOKEN_DECLARATION.Symbol() != i[0].Symbol() {
			msg := "Unexpected symbol: '" + i[0].Symbol() + "', expected " + TOKEN_DECLARATION.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != Equal.Symbol() {
			msg := "Unexpected symbol: '" + i[1].Symbol() + "', expected " + Equal.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != STRING_LIST.Symbol() {
			msg := "Unexpected symbol: '" + i[2].Symbol() + "', expected " + STRING_LIST.Symbol()
			panic(msg)
		}
		decl, _ := i[0].(*TOKEN_DECLARATION_DEFINED_BY_LIST_AST)
		string_sequence, _ := i[2].(*STRING_SEQUENCE_AST)
		tokens := string_sequence.Items
		decl.Tokens = tokens
		decl.UpdateSymbol(s)
		return decl
	})

	// TOKEN_DECLARATION ----> Token Variable
	TokenGrammar.AddProduction(TOKEN_DECLARATION, []IGrammarSymbol{TokenKeyword, Variable}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != TokenKeyword.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + TokenKeyword.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != Variable.Symbol() {
			msg := "Unexpected symbol '" + i[1].Symbol() + "', expected " + Variable.Symbol()
			panic(msg)
		}
		decl := NewTokenDeclarationAST(TOKEN_DECLARATION.Symbol(), i[0].Line(), i[0].Column())
		return decl
	})

	// STRING_LIST ----> [ STRING_SEQUENCE ]
	TokenGrammar.AddProduction(STRING_LIST, []IGrammarSymbol{OpenCorchet, STRING_SEQUENCE, ClosedCorchet}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != OpenCorchet.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + OpenCorchet.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != STRING_SEQUENCE.Symbol() {
			msg := "Unexpected symbol '" + i[1].Symbol() + "', expected " + STRING_SEQUENCE.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != ClosedCorchet.Symbol() {
			msg := "Unexpected symbol '" + i[2].Symbol() + "', expected " + ClosedCorchet.Symbol()
			panic(msg)
		}
		i[1].UpdateSymbol(s)
		return i[1]
	})

	// STRING_SEQUENCE ----> Text
	TokenGrammar.AddProduction(STRING_SEQUENCE, []IGrammarSymbol{Text}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != Text.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + Text.Symbol()
			panic(msg)
		}
		text, _ := i[0].Eval(nil, nil).(string)
		seq := NewStringSequenceAST(STRING_SEQUENCE.Symbol(), i[0].Line(), i[0].Column())
		seq.Items = []string{text}
		return seq
	})

	// STRING_SEQUENCE ----> STRING_SEQUENCE , Text
	TokenGrammar.AddProduction(STRING_SEQUENCE, []IGrammarSymbol{STRING_SEQUENCE, Comma, Text}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != STRING_SEQUENCE.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + STRING_SEQUENCE.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != Comma.Symbol() {
			msg := "Unexpected symbol '" + i[1].Symbol() + "', expected " + Comma.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != Text.Symbol() {
			msg := "Unexpected symbol '" + i[2].Symbol() + "', expected " + Text.Symbol()
			panic(msg)
		}
		seq, _ := i[0].(*STRING_SEQUENCE_AST)
		word, _ := i[2].(*AtomicAST).Eval(nil, nil).(string)
		seq.Items = append(seq.Items, word)
		return seq
	})

	// RIGHT_REGULAR_GRAMMAR ----> GRAMMAR_DECLARATION = RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE
	TokenGrammar.AddProduction(RIGHT_REGULAR_GRAMMAR, []IGrammarSymbol{GRAMMAR_DECLARATION, Equal, RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != GRAMMAR_DECLARATION.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + GRAMMAR_DECLARATION.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != Equal.Symbol() {
			msg := "Unexpected symbol '" + i[1].Symbol() + "', expected " + Equal.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE.Symbol() {
			msg := "Unexpected symbol '" + i[2].Symbol() + "', expected " + RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE.Symbol()
			panic(msg)
		}
		decl, _ := i[0].(*RIGHT_REGULAR_GRAMMAR_AST)
		productions_sequence, _ := i[2].(*RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE_AST)
		productions := productions_sequence.Productions
		for _, production := range productions {
			decl.RegularGrammar.AddProduction(production.Head, production.Production)
		}
		decl.UpdateSymbol(s)
		return decl
	})

	// GRAMMAR_DECLARATION ----> grammar variable ( Text )
	TokenGrammar.AddProduction(GRAMMAR_DECLARATION, []IGrammarSymbol{GrammarKeyword, Variable, OpenParent, Variable, ClosedParent}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != GrammarKeyword.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + GrammarKeyword.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != Variable.Symbol() {
			msg := "Unexpected symbol '" + i[1].Symbol() + "', expected " + Variable.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != OpenParent.Symbol() {
			msg := "Unexpected symbol '" + i[2].Symbol() + "', expected " + Variable.Symbol()
			panic(msg)
		}
		if i[3].Symbol() != Variable.Symbol() {
			msg := "Unexpected symbol '" + i[3].Symbol() + "', expected " + Variable.Symbol()
			panic(msg)
		}
		if i[4].Symbol() != ClosedParent.Symbol() {
			msg := "Unexpected symbol '" + i[4].Symbol() + "', expected " + ClosedParent.Symbol()
			panic(msg)
		}
		decl, _ := i[0].(*GRAMMAR_DECLARATION_AST)
		grammar_name_variable_ast, _ := i[1].(*VARIABLE_AST)
		start_symbol_name_variable_ast, _ := i[3].(*VARIABLE_AST)
		start_symbol := NewGrammarSymbol(start_symbol_name_variable_ast.VariableName, NonTerminal, false)
		decl.GrammarName = grammar_name_variable_ast.VariableName
		decl.GrammarValue = NewGrammar(start_symbol)
		decl.UpdateSymbol(s)
		return decl
	})

	// RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE ----> RIGHT_REGULAR_GRAMMAR_PRODUCTION
	TokenGrammar.AddProduction(RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE, []IGrammarSymbol{RIGHT_REGULAR_GRAMMAR_PRODUCTION}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != RIGHT_REGULAR_GRAMMAR_PRODUCTION.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + RIGHT_REGULAR_GRAMMAR_PRODUCTION.Symbol()
			panic(msg)
		}
		i[0].UpdateSymbol(s)
		return i[0]
	})

	// RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE ----> RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE , RIGHT_REGULAR_GRAMMAR_PRODUCTION
	TokenGrammar.AddProduction(RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE, []IGrammarSymbol{RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE, Comma, RIGHT_REGULAR_GRAMMAR_PRODUCTION}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != Comma.Symbol() {
			msg := "Unexpected symbol '" + i[1].Symbol() + "', expected " + Comma.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != RIGHT_REGULAR_GRAMMAR_PRODUCTION.Symbol() {
			msg := "Unexpected symbol '" + i[2].Symbol() + "', expected " + RIGHT_REGULAR_GRAMMAR_PRODUCTION.Symbol()
			panic(msg)
		}
		left_sequence, _ := i[0].(*RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE_AST)
		right_sequence, _ := i[2].(*RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE_AST)
		left_sequence.Productions = append(left_sequence.Productions, right_sequence.Productions...)
		return left_sequence
	})

	// RIGHT_REGULAR_GRAMMAR_PRODUCTION ----> < variable > ---> RIGHT_REGULAR_DERIVATION_SEQUENCE
	TokenGrammar.AddProduction(RIGHT_REGULAR_GRAMMAR_PRODUCTION, []IGrammarSymbol{LessThan, Variable, GreaterThan, GrammarProductionArrow, RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != LessThan.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + LessThan.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != Variable.Symbol() {
			msg := "Unexpected symbol '" + i[1].Symbol() + "', expected " + Variable.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != GreaterThan.Symbol() {
			msg := "Unexpected symbol '" + i[2].Symbol() + "', expected " + GreaterThan.Symbol()
			panic(msg)
		}
		if i[3].Symbol() != GrammarProductionArrow.Symbol() {
			msg := "Unexpected symbol '" + i[3].Symbol() + "', expected " + GrammarProductionArrow.Symbol()
			panic(msg)
		}
		if i[4].Symbol() != RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE.Symbol() {
			msg := "Unexpected symbol '" + i[4].Symbol() + "', expected " + RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE.Symbol()
			panic(msg)
		}
		production := NewRightRegularGrammarProductionSequenceAST(s, i[0].Line(), i[0].Column())
		derivations_ast, _ := i[4].(*RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE_AST)
		var_ast, _ := i[1].(*VARIABLE_AST)
		head_symbol := NewGrammarSymbol(var_ast.VariableName, NonTerminal, false)
		for _, derivation := range derivations_ast.Derivations {
			production.Productions = append(production.Productions, GrammarProduction{
				Head:       head_symbol,
				Production: derivation,
			})
		}
		return production
	})

	// RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE ----> Text < variable >
	TokenGrammar.AddProduction(RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE, []IGrammarSymbol{Text, LessThan, Variable, GreaterThan}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != Text.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + Text.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != LessThan.Symbol() {
			msg := "Unexpected symbol '" + i[1].Symbol() + "', expected " + LessThan.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != Variable.Symbol() {
			msg := "Unexpected symbol '" + i[2].Symbol() + "', expected " + Variable.Symbol()
			panic(msg)
		}
		if i[3].Symbol() != GreaterThan.Symbol() {
			msg := "Unexpected symbol '" + i[3].Symbol() + "', expected " + GreaterThan.Symbol()
			panic(msg)
		}
		derivation_ast := NewRightRegularGrammarDerivationSequenceAST(s, i[0].Line(), i[0].Column())
		text_value_ast, _ := i[0].(*AtomicAST)
		text := text_value_ast.Eval(nil, nil).(string)
		variable_ast, _ := i[2].(*VARIABLE_AST)
		terminal := NewGrammarSymbol(text, Terminal, false)
		non_terminal := NewGrammarSymbol(variable_ast.VariableName, NonTerminal, false)
		derivation := []IGrammarSymbol{terminal, non_terminal}
		derivation_ast.Derivations = append(derivation_ast.Derivations, derivation)
		return derivation_ast
	})

	// RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE ----> Text
	TokenGrammar.AddProduction(RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE, []IGrammarSymbol{Text}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != Text.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + Text.Symbol()
			panic(msg)
		}
		derivation_ast := NewRightRegularGrammarDerivationSequenceAST(s, i[0].Line(), i[0].Column())
		text_value_ast, _ := i[0].(*AtomicAST)
		text := text_value_ast.Eval(nil, nil).(string)
		terminal := NewGrammarSymbol(text, Terminal, false)
		derivation := []IGrammarSymbol{terminal}
		derivation_ast.Derivations = append(derivation_ast.Derivations, derivation)
		return derivation_ast
	})

	// RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE ----> < variable >
	TokenGrammar.AddProduction(RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE, []IGrammarSymbol{LessThan, Variable, GreaterThan}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != LessThan.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + LessThan.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != Variable.Symbol() {
			msg := "Unexpected symbol '" + i[1].Symbol() + "', expected " + Variable.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != GreaterThan.Symbol() {
			msg := "Unexpected symbol '" + i[2].Symbol() + "', expected " + GreaterThan.Symbol()
			panic(msg)
		}
		derivation_ast := NewRightRegularGrammarDerivationSequenceAST(s, i[0].Line(), i[0].Column())
		variable_ast, _ := i[1].(*VARIABLE_AST)
		non_terminal := NewGrammarSymbol(variable_ast.VariableName, NonTerminal, false)
		derivation := []IGrammarSymbol{non_terminal}
		derivation_ast.Derivations = append(derivation_ast.Derivations, derivation)
		return derivation_ast
	})

	// RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE ----> epsilon
	TokenGrammar.AddProduction(RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE, []IGrammarSymbol{Epsilon}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != Epsilon.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + Epsilon.Symbol()
			panic(msg)
		}
		derivation_ast := NewRightRegularGrammarDerivationSequenceAST(s, i[0].Line(), i[0].Column())
		terminal := NewGrammarSymbol("epsilon", Terminal, true)
		derivation := []IGrammarSymbol{terminal}
		derivation_ast.Derivations = append(derivation_ast.Derivations, derivation)
		return derivation_ast
	})

	// RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE ----> RIGHT_REGULAR_DERIVATION_SEQUENCE | Text < variable >
	TokenGrammar.AddProduction(RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE, []IGrammarSymbol{RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE, Or, Text, LessThan, Variable, GreaterThan}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != Or.Symbol() {
			msg := "Unexpected symbol '" + i[1].Symbol() + "', expected " + Or.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != Text.Symbol() {
			msg := "Unexpected symbol '" + i[2].Symbol() + "', expected " + Text.Symbol()
			panic(msg)
		}
		if i[3].Symbol() != LessThan.Symbol() {
			msg := "Unexpected symbol '" + i[3].Symbol() + "', expected " + LessThan.Symbol()
			panic(msg)
		}
		if i[4].Symbol() != Variable.Symbol() {
			msg := "Unexpected symbol '" + i[4].Symbol() + "', expected " + Variable.Symbol()
			panic(msg)
		}
		if i[5].Symbol() != GreaterThan.Symbol() {
			msg := "Unexpected symbol '" + i[5].Symbol() + "', expected " + GreaterThan.Symbol()
			panic(msg)
		}
		derivation_ast, _ := i[0].(*RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE_AST)
		text_value_ast, _ := i[2].(*AtomicAST)
		text := text_value_ast.Eval(nil, nil).(string)
		variable_ast, _ := i[4].(*VARIABLE_AST)
		terminal := NewGrammarSymbol(text, Terminal, false)
		non_terminal := NewGrammarSymbol(variable_ast.VariableName, NonTerminal, false)
		derivation := []IGrammarSymbol{terminal, non_terminal}
		derivation_ast.Derivations = append(derivation_ast.Derivations, derivation)
		return derivation_ast
	})

	// RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE ----> RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE | Text
	TokenGrammar.AddProduction(RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE, []IGrammarSymbol{RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE, Or, Text}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != Or.Symbol() {
			msg := "Unexpected symbol '" + i[1].Symbol() + "', expected " + Or.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != Text.Symbol() {
			msg := "Unexpected symbol '" + i[2].Symbol() + "', expected " + Text.Symbol()
			panic(msg)
		}
		derivation_ast, _ := i[0].(*RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE_AST)
		text_value_ast, _ := i[2].(*AtomicAST)
		text := text_value_ast.Eval(nil, nil).(string)
		terminal := NewGrammarSymbol(text, Terminal, false)
		derivation := []IGrammarSymbol{terminal}
		derivation_ast.Derivations = append(derivation_ast.Derivations, derivation)
		return derivation_ast
	})

	// RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE ----> RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE | < variable >
	TokenGrammar.AddProduction(RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE, []IGrammarSymbol{RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE, Or, LessThan, Variable, GreaterThan}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != Or.Symbol() {
			msg := "Unexpected symbol '" + i[1].Symbol() + "', expected " + Or.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != LessThan.Symbol() {
			msg := "Unexpected symbol '" + i[2].Symbol() + "', expected " + LessThan.Symbol()
			panic(msg)
		}
		if i[3].Symbol() != Variable.Symbol() {
			msg := "Unexpected symbol '" + i[3].Symbol() + "', expected " + Variable.Symbol()
			panic(msg)
		}
		if i[4].Symbol() != GreaterThan.Symbol() {
			msg := "Unexpected symbol '" + i[4].Symbol() + "', expected " + GreaterThan.Symbol()
			panic(msg)
		}
		derivation_ast, _ := i[0].(*RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE_AST)
		variable_ast, _ := i[4].(*VARIABLE_AST)
		non_terminal := NewGrammarSymbol(variable_ast.VariableName, NonTerminal, false)
		derivation := []IGrammarSymbol{non_terminal}
		derivation_ast.Derivations = append(derivation_ast.Derivations, derivation)
		return derivation_ast
	})

	// RIGHT_REGULAR_GRAMMAR_DERIVATION_SQUENCE ----> RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE | epsilon
	TokenGrammar.AddProduction(RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE, []IGrammarSymbol{RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE, Or, Epsilon}, func(i []IAST, s string) IAST {
		if i[0].Symbol() != RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE.Symbol() {
			msg := "Unexpected symbol '" + i[0].Symbol() + "', expected " + RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE.Symbol()
			panic(msg)
		}
		if i[1].Symbol() != Or.Symbol() {
			msg := "Unexpected symbol '" + i[1].Symbol() + "', expected " + Or.Symbol()
			panic(msg)
		}
		if i[2].Symbol() != Epsilon.Symbol() {
			msg := "Unexpected symbol '" + i[2].Symbol() + "', expected " + Epsilon.Symbol()
			panic(msg)
		}
		derivation_ast, _ := i[0].(*RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE_AST)
		terminal := NewGrammarSymbol("epsilon", Terminal, true)
		derivation := []IGrammarSymbol{terminal}
		derivation_ast.Derivations = append(derivation_ast.Derivations, derivation)
		return derivation_ast
	})

}
