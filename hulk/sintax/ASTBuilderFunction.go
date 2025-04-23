package sintax

import (
	"strconv"

	. "hulk.com/app/ast"
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
	. "hulk.com/app/tokens"
)

var symbolsByToken map[TokenType]string
var binaryOperatorFunction map[string]func(a IAST, b IAST, context IContext, collector IErrorCollector) interface{}

func init() {
	symbolsByToken = make(map[TokenType]string)
	binaryOperatorFunction = make(map[string]func(a IAST, b IAST, context IContext, collector IErrorCollector) interface{})

	symbolsByToken[BooleanToken] = "boolean"
	symbolsByToken[NumberToken] = "number"
	symbolsByToken[StringToken] = "string"

	binaryOperatorFunction["+"] = func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
		a := left.Eval(context, collector)
		b := right.Eval(context, collector)
		af, a_ok := a.(float64)
		bf, b_ok := b.(float64)

		if a_ok && b_ok {
			return af + bf
		}
		collector.AddError(NewError("The operator + only can be applied to numbers", left.Line(), left.Column(), Gramatical))
		return nil
	}
	binaryOperatorFunction["-"] = func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
		a := left.Eval(context, collector)
		b := right.Eval(context, collector)
		af, a_ok := a.(float64)
		bf, b_ok := b.(float64)

		if a_ok && b_ok {
			return af - bf
		}
		collector.AddError(NewError("The operator - only can be applied to numbers", left.Line(), left.Column(), Gramatical))
		return nil
	}
	binaryOperatorFunction["*"] = func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
		a := left.Eval(context, collector)
		b := right.Eval(context, collector)
		af, a_ok := a.(float64)
		bf, b_ok := b.(float64)

		if a_ok && b_ok {
			return af * bf
		}
		collector.AddError(NewError("The operator * only can be applied to numbers", left.Line(), left.Column(), Gramatical))
		return nil
	}
	binaryOperatorFunction["/"] = func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
		a := left.Eval(context, collector)
		b := right.Eval(context, collector)
		af, a_ok := a.(float64)
		bf, b_ok := b.(float64)

		if a_ok && b_ok {
			return af / bf
		}
		collector.AddError(NewError("The operator / only can be applied to numbers", left.Line(), left.Column(), Gramatical))
		return nil
	}
	binaryOperatorFunction["%"] = func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
		a := left.Eval(context, collector)
		b := right.Eval(context, collector)
		af, a_ok := a.(float64)
		bf, b_ok := b.(float64)

		if a_ok && b_ok {

			if float64(int(af)) == af && float64(int(bf)) == bf {
				return int(af) % int(bf)
			}
			collector.AddError(NewError("The operator % only can be applied to integers", left.Line(), left.Column(), Semantic))
			return nil
		}
		collector.AddError(NewError("The operator % only can be applied to numbers", left.Line(), left.Column(), Gramatical))
		return nil
	}
}

func HulkASTBuilder(token IToken, endmarker string) IAST {
	switch token.Type() {
	case BooleanToken:
		if token.Text() == "false" {
			return NewAtomicAST(symbolsByToken[BooleanToken], token.Line(), token.Column(), false)
		}
		return NewAtomicAST(symbolsByToken[token.Type()], token.Line(), token.Column(), true)
	case StringToken:
		return NewAtomicAST(symbolsByToken[token.Type()], token.Line(), token.Column(), token.Text())
	case NumberToken:
		value, _ := strconv.ParseFloat(token.Text(), 64)
		return NewAtomicAST(symbolsByToken[token.Type()], token.Line(), token.Column(), value)
	case OperatorToken:
		if f, ok := binaryOperatorFunction[token.Text()]; ok {
			return NewBinaryAST(token.Text(), token.Line(), token.Column(), f)
		}
		panic("Operator " + token.Text() + " not implemented")
	case SymbolToken:
		if token.Text() == endmarker {
			return NewAtomicAST(endmarker, token.Line(), token.Column(), token.Text())
		}
		return NewAtomicAST(token.Text(), token.Line(), token.Column(), token.Text())
	default:
		return nil
	}
}
