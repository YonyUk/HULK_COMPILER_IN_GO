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

	// TOKEN ----> TOKEN_DECLARATION = STRING_LIST
	TokenGrammar.AddProduction(TOKEN, []IGrammarSymbol{TOKEN_DECLARATION, Equal, STRING_LIST}, func(i []IAST, s string) IAST {
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
		decl, _ := i[0].(*TOKEN_DECLARATION_AST)
		string_sequence, _ := i[2].(*STRING_SEQUENCE_AST)
		tokens := string_sequence.Items
		decl.TokenGrammar = GetWordsGrammar(tokens)
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
}
