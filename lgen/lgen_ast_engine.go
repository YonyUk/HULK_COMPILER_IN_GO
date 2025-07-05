package lgen

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/lgen/lgen_ast"
	. "hulk.com/app/tokens"
)

func AstEngine(token IToken, endmarker string) IAST {
	switch token.Type() {
	case VariableToken:
		return NewVariableAST("variable", token.Line(), token.Column(), token.Text())

	case KeywordToken:
		return keyword_ast_engine(token)

	case SymbolToken:
		return symbol_ast_engine(token)

	case OperatorToken:
		return operator_aste_engine(token)

	case LiteralStringToken:
		return NewAtomicAST("Text", token.Line(), token.Column(), token.Text())

	case EndToken:
		return NewAtomicAST(token.Text(), token.Line(), token.Column(), token.Text())

	default:
		panic("Unknown token: " + token.Text())
	}
}

func operator_aste_engine(token IToken) IAST {
	switch token.Text() {
	case "=":
		return NewAssigmentAst(token.Text(), token.Line(), token.Column())

	case "<":
		return NewAtomicAST("<", token.Line(), token.Column(), "<")

	case ">":
		return NewAtomicAST(">", token.Line(), token.Column(), ">")

	default:
		panic("Unkown symbol " + token.Text())

	}
}

func keyword_ast_engine(token IToken) IAST {
	switch token.Text() {
	case "token":
		return NewTokenDeclarationAST(token.Text(), token.Line(), token.Column())

	case "grammar":
		return NewGrammarDeclarationAST(token.Text(), token.Line(), token.Column())

	case "epsilon":
		return NewAtomicAST("epsilon", token.Line(), token.Column(), "epsilon")

	default:
		panic("Unknown keyword: " + token.Text())
	}
}

func symbol_ast_engine(token IToken) IAST {
	switch token.Text() {
	case "[":
		return NewAtomicAST("[", token.Line(), token.Column(), "[")

	case "]":
		return NewAtomicAST("]", token.Line(), token.Column(), "]")

	case ";":
		return NewAtomicAST(";", token.Line(), token.Column(), ";")

	case ",":
		return NewAtomicAST(",", token.Line(), token.Column(), ",")

	case "--->":
		return NewAtomicAST("--->", token.Line(), token.Column(), "--->")

	case "(":
		return NewAtomicAST("(", token.Line(), token.Column(), "(")

	case ")":
		return NewAtomicAST(")", token.Line(), token.Column(), ")")

	default:
		panic("Unknown symbol: " + token.Text())
	}
}
