package main

import (
	"strconv"

	. "hulk.com/app/ast"
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
	. "hulk.com/app/grammar"
	. "hulk.com/app/parser"
	. "hulk.com/app/tokens"
)

func main() {
	E := NewGrammarSymbol("E", NonTerminal, false)
	T := NewGrammarSymbol("T", NonTerminal, false)
	F := NewGrammarSymbol("F", NonTerminal, false)
	Plus := NewGrammarSymbol("+", Terminal, false)
	Mul := NewGrammarSymbol("*", Terminal, false)
	LP := NewGrammarSymbol("(", Terminal, false)
	RP := NewGrammarSymbol(")", Terminal, false)
	number := NewGrammarSymbol("number", Terminal, false)

	G := NewAttributedGrammar(E)
	G.AddProduction(E, []IGrammarSymbol{E, Plus, T}, func(asts []IAST, s string) IAST {
		if asts[1].Symbol() == Plus.Symbol() {
			ast_, _ := asts[1].(*BinaryAST)
			ast_.Left = asts[0]
			ast_.Right = asts[2]
			ast_.UpdateSymbol(s)
			return ast_
		}
		msg_error := "Error while parsing, expected '" + Plus.Symbol() + "' and got " + asts[1].Symbol() + " at " + strconv.Itoa(asts[1].Line()) + ", " + strconv.Itoa(asts[1].Column())
		panic(msg_error)
	})
	G.AddProduction(E, []IGrammarSymbol{T}, func(asts []IAST, s string) IAST {
		if asts[0].Symbol() == T.Symbol() {
			asts[0].UpdateSymbol(s)
			return asts[0]
		}
		msg_error := "Error while parsing, expected '" + T.Symbol() + "' and got " + asts[0].Symbol() + " at " + strconv.Itoa(asts[0].Line()) + ", " + strconv.Itoa(asts[0].Column())
		panic(msg_error)
	})
	G.AddProduction(T, []IGrammarSymbol{T, Mul, F}, func(asts []IAST, s string) IAST {
		if asts[1].Symbol() == Mul.Symbol() {
			ast_, _ := asts[1].(*BinaryAST)
			ast_.Left = asts[0]
			ast_.Right = asts[2]
			ast_.UpdateSymbol(s)
			return ast_
		}
		msg_error := "Error while parsing, expected '" + Mul.Symbol() + "' and got " + asts[1].Symbol() + " at " + strconv.Itoa(asts[1].Line()) + ", " + strconv.Itoa(asts[1].Column())
		panic(msg_error)
	})
	G.AddProduction(T, []IGrammarSymbol{F}, func(asts []IAST, s string) IAST {
		if asts[0].Symbol() == F.Symbol() {
			asts[0].UpdateSymbol(s)
			return asts[0]
		}
		msg_error := "Error while parsing, expected '" + F.Symbol() + "' and got " + asts[0].Symbol() + " at " + strconv.Itoa(asts[0].Line()) + ", " + strconv.Itoa(asts[0].Column())
		panic(msg_error)
	})
	G.AddProduction(F, []IGrammarSymbol{LP, E, RP}, func(asts []IAST, s string) IAST {
		if asts[1].Symbol() == E.Symbol() {
			asts[1].UpdateSymbol(s)
			return asts[1]
		}
		msg_error := "Error while parsing, expected '" + E.Symbol() + "' and got " + asts[1].Symbol() + " at " + strconv.Itoa(asts[1].Line()) + ", " + strconv.Itoa(asts[1].Column())
		panic(msg_error)
	})
	G.AddProduction(F, []IGrammarSymbol{number}, func(asts []IAST, s string) IAST {
		if asts[0].Symbol() == number.Symbol() {
			asts[0].UpdateSymbol(s)
			return asts[0]
		}
		msg_error := "Error while parsing, expected '" + number.Symbol() + "' and got " + asts[0].Symbol() + " at " + strconv.Itoa(asts[0].Line()) + ", " + strconv.Itoa(asts[0].Column())
		panic(msg_error)
	})

	parser := NewParserSLRFromAttributedGrammar(G, NewGrammarSymbol("$", Terminal, false), ast_engine)
	tokens_ := []IToken{
		NewToken(0, 0, "0", NumberToken),
		NewToken(0, 0, "+", OperatorToken),
		NewToken(0, 0, "0", NumberToken),
		NewToken(0, 0, "$", SymbolToken),
	}

	for _, token := range tokens_ {
		parser.Parse(token, nil)
	}
}

func ast_engine(token IToken, endmarker string) IAST {
	if token.Text() == "0" {
		return NewAtomicAST("number", token.Line(), token.Column(), 0)
	}
	if token.Text() == "+" {
		return NewBinaryAST("+", token.Line(), token.Column(), func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
			left_, _ := left.Eval(context, collector).(int)
			right_ := right.Eval(context, collector).(int)
			return left_ + right_
		})
	}
	if token.Text() == "*" {
		return NewBinaryAST("+", token.Line(), token.Column(), func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
			left_, _ := left.Eval(context, collector).(int)
			right_ := right.Eval(context, collector).(int)
			return left_ * right_
		})
	}
	if token.Text() == "(" {
		return NewAtomicAST("(", token.Line(), token.Column(), nil)
	}
	if token.Text() == ")" {
		return NewAtomicAST(")", token.Line(), token.Column(), nil)
	}
	return NewAtomicAST("$", token.Line(), token.Column(), nil)
}
