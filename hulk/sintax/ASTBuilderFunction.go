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
	. "hulk.com/app/tools"
)

var symbolsByToken map[TokenType]string
var binaryOperatorFunction map[string]func(a IAST, b IAST, context IContext, collector IErrorCollector) interface{}
var unaryOperatorFunction map[string]func(target IAST, context IContext, collector IErrorCollector) interface{}
var InFrontOperators []string
var InBackOperators []string

func init() {
	symbolsByToken = make(map[TokenType]string)
	binaryOperatorFunction = make(map[string]func(a IAST, b IAST, context IContext, collector IErrorCollector) interface{})
	unaryOperatorFunction = make(map[string]func(target IAST, context IContext, collector IErrorCollector) interface{})

	InFrontOperators = []string{
		"!",
	}

	InBackOperators = []string{}

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
		var error_line int
		var error_column int
		if a_err != nil {
			error_line = left.Line()
			error_column = left.Column()
		} else {
			error_line = right.Line()
			error_column = right.Column()
		}
		collector.AddError(NewError("The operator + only can be applied to arithmetic expressions", error_line, error_column, Gramatical))
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
		var error_line int
		var error_column int
		if a_err != nil {
			error_line = left.Line()
			error_column = left.Column()
		} else {
			error_line = right.Line()
			error_column = right.Column()
		}
		collector.AddError(NewError("The operator - only can be applied to arithmetic expressions", error_line, error_column, Gramatical))
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
		var error_line int
		var error_column int
		if a_err != nil {
			error_line = left.Line()
			error_column = left.Column()
		} else {
			error_line = right.Line()
			error_column = right.Column()
		}
		collector.AddError(NewError("The operator * only can be applied to arithmetic expressions", error_line, error_column, Gramatical))
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
		var error_line int
		var error_column int
		if a_err != nil {
			error_line = left.Line()
			error_column = left.Column()
		} else {
			error_line = right.Line()
			error_column = right.Column()
		}
		collector.AddError(NewError("The operator / only can be applied to arithmetic expressions", error_line, error_column, Gramatical))
		return nil
	}
	binaryOperatorFunction["%"] = func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
		a := left.Eval(context, collector)
		b := right.Eval(context, collector)
		a_value, a_err := isNumber(a)
		b_value, b_err := isNumber(b)
		var error_line int
		var error_column int
		if a_err == nil && b_err == nil {
			if float64(int(a_value)) == a_value && float64(int(b_value)) == b_value {
				return int(a_value) % int(b_value)
			}
			if float64(int(a_value)) != a_value {
				error_line = left.Line()
				error_column = left.Column()
			} else {
				error_line = right.Line()
				error_column = right.Column()
			}
			collector.AddError(NewError("The operator % only can be applied to integers", error_line, error_column, Semantic))
			return nil
		}
		if a_err != nil {
			error_line = left.Line()
			error_column = left.Column()
		} else {
			error_line = right.Line()
			error_column = right.Column()
		}
		collector.AddError(NewError("The operator % only can be applied to arithmetic expressions", error_line, error_column, Gramatical))
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
		var error_line int
		var error_column int
		if a_err != nil {
			error_line = left.Line()
			error_column = left.Column()
		} else {
			error_line = right.Line()
			error_column = right.Column()
		}
		collector.AddError(NewError("The operator ^ only can be applied to arithmetic expressions", error_line, error_column, Gramatical))
		return nil
	}
	binaryOperatorFunction["&"] = func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
		a := left.Eval(context, collector)
		b := right.Eval(context, collector)
		a_bool, a_ok := a.(bool)
		b_bool, b_ok := b.(bool)

		if a_ok && b_ok {
			return a_bool && b_bool
		}
		var error_line int
		var error_column int
		if !a_ok {
			error_line = left.Line()
			error_column = left.Column()
		} else {
			error_line = right.Line()
			error_column = right.Column()
		}
		collector.AddError(NewError("The operator & only can be applied to boolean expressions", error_line, error_column, Gramatical))
		return nil
	}
	binaryOperatorFunction["|"] = func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{} {
		a := left.Eval(context, collector)
		b := right.Eval(context, collector)
		a_bool, a_ok := a.(bool)
		b_bool, b_ok := b.(bool)

		if a_ok && b_ok {
			return a_bool || b_bool
		}
		var error_line int
		var error_column int
		if !a_ok {
			error_line = left.Line()
			error_column = left.Column()
		} else {
			error_line = right.Line()
			error_column = right.Column()
		}
		collector.AddError(NewError("The operator & only can be applied to boolean expressions", error_line, error_column, Gramatical))
		return nil
	}

	unaryOperatorFunction["!"] = func(target IAST, context IContext, collector IErrorCollector) interface{} {
		a := target.Eval(context, collector)
		a_bool, a_ok := a.(bool)
		if a_ok {
			return !a_bool
		}
		collector.AddError(NewError("The operator ! only can be applied to boolean expressions", target.Line(), target.Column(), Gramatical))
		return nil
	}
}

func HulkASTBuilder(token IToken, endmarker string) IAST {
	switch token.Type() {
	case BooleanToken:
		value, _ := strconv.ParseBool(token.Text())
		return NewAtomicAST(symbolsByToken[token.Type()], token.Line(), token.Column(), value)
	case StringToken:
		return NewAtomicAST(symbolsByToken[token.Type()], token.Line(), token.Column(), token.Text())
	case NumberToken:
		value, _ := strconv.ParseFloat(token.Text(), 64)
		return NewAtomicAST(symbolsByToken[token.Type()], token.Line(), token.Column(), value)
	case OperatorToken:
		if f, ok := binaryOperatorFunction[token.Text()]; ok {
			return NewBinaryAST(token.Text(), token.Line(), token.Column(), f)
		}
		if f, ok := unaryOperatorFunction[token.Text()]; ok {
			if _, err := IndexOf(InFrontOperators, func(s string) bool { return s == token.Text() }); err == nil {
				return NewInFrontOperatorAST(token.Text(), token.Line(), token.Column(), f)
			}
			return nil
		}
		panic("Operator " + token.Text() + " not implemented")
	case SymbolToken:
		if token.Text() == endmarker {
			return NewAtomicAST(endmarker, token.Line(), token.Column(), token.Text())
		}
		return NewAtomicAST(token.Text(), token.Line(), token.Column(), token.Text())
	default:
		return NewGarbageAST(token.Text(), token.Line(), token.Column(), token.Text())
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
