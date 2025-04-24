package sintax

import (
	"errors"
	"math"
	"reflect"
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
		a_value, a_err := isNumber(a)
		b_value, b_err := isNumber(b)
		if a_err == nil && b_err == nil {
			return a_value + b_value
		}
		collector.AddError(NewError("The operator + only can be applied to numbers", left.Line(), left.Column(), Gramatical))
		return nil
	}
	binaryOperatorFunction["-"] = func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
		a := left.Eval(context, collector)
		b := right.Eval(context, collector)
		a_value, a_err := isNumber(a)
		b_value, b_err := isNumber(b)
		if a_err == nil && b_err == nil {
			return a_value - b_value
		}
		collector.AddError(NewError("The operator + only can be applied to numbers", left.Line(), left.Column(), Gramatical))
		return nil
	}
	binaryOperatorFunction["*"] = func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
		a := left.Eval(context, collector)
		b := right.Eval(context, collector)
		a_value, a_err := isNumber(a)
		b_value, b_err := isNumber(b)
		if a_err == nil && b_err == nil {
			return a_value * b_value
		}
		collector.AddError(NewError("The operator + only can be applied to numbers", left.Line(), left.Column(), Gramatical))
		return nil
	}
	binaryOperatorFunction["/"] = func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
		a := left.Eval(context, collector)
		b := right.Eval(context, collector)
		a_value, a_err := isNumber(a)
		b_value, b_err := isNumber(b)
		if a_err == nil && b_err == nil {
			return a_value / b_value
		}
		collector.AddError(NewError("The operator + only can be applied to numbers", left.Line(), left.Column(), Gramatical))
		return nil
	}
	binaryOperatorFunction["%"] = func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
		a := left.Eval(context, collector)
		b := right.Eval(context, collector)
		a_value, a_err := isNumber(a)
		b_value, b_err := isNumber(b)
		if a_err == nil && b_err == nil {
			if int(a_value) == int(a_value) && int(b_value) == int(b_value) {
				return int(a_value) % int(b_value)
			}
			collector.AddError(NewError("The operator % only can be applied to integers", left.Line(), left.Column(), Semantic))
		}
		collector.AddError(NewError("The operator % only can be applied to numbers", left.Line(), left.Column(), Gramatical))
		return nil
	}
	binaryOperatorFunction["^"] = func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
		a := left.Eval(context, collector)
		b := right.Eval(context, collector)
		a_value, a_err := isNumber(a)
		b_value, b_err := isNumber(b)
		if a_err == nil && b_err == nil {
			return math.Pow(a_value, b_value)
		}
		collector.AddError(NewError("The operator ^ only can be applied to numbers", left.Line(), left.Column(), Gramatical))
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

func isNumber(value interface{}) (float64, error) {
	type_ := reflect.TypeOf(value).Kind()
	switch type_ {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(value.(int)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(value.(uint)), nil
	case reflect.Float32, reflect.Float64:
		return value.(float64), nil
	default:
		return 0, errors.New("The value is not a number")
	}
}
