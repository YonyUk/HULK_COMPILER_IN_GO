package hulk

import (
	. "hulk.com/app/compiler"
	. "hulk.com/app/interpreter"
	. "hulk.com/app/lexer"
	. "hulk.com/app/lexical_analisys"
	. "hulk.com/app/tokens"
)

type HulkInterpreter struct {
	_interpreter IInterpreter
}

func NewHulkInterpreter() *HulkInterpreter {
	lexer := NewLexer()
	lexer.AddTokenExpression(KeywordToken, 0, KeywordsGrammar)
	lexer.AddTokenExpression(BooleanToken, 1, BooleanGrammar)
	lexer.AddTokenExpression(OperatorToken, 2, OperatorGrammar)
	lexer.AddTokenExpression(NumberToken, 3, NumberGrammar)
	lexer.AddTokenExpression(VariableToken, 4, VariableGrammar)
	lexer.AddTokenExpression(SymbolToken, 5, SymbolGrammar)
	lexer.AddTokenExpression(StringToken, 6, StringGrammar)

	var_rule := NewLexicalRule("The var names most starts with a letter", func(t IToken) bool {
		return (rune(48) > rune(t.Text()[0]) || rune(t.Text()[0]) > rune(57)) && string(t.Text()[0]) != "_"
	})
	num_rule := NewLexicalRule("Numbers most have only one zero at the begining", func(t IToken) bool {
		switch len(t.Text()) {
		case 1:
			return true
		default:
			if string(t.Text()[0]) == "+" || string(t.Text()[0]) == "-" {
				if len(t.Text()) == 2 {
					return true
				}
				if string(t.Text()[1]) == "0" && string(t.Text()[2]) != "." {
					return false
				}
				return true
			}
			if string(t.Text()[0]) == "0" && string(t.Text()[1]) != "." {
				return false
			}
			return true
		}
	})

	lexical := NewLexicalAnalizer()
	lexical.AddRule(VariableToken, var_rule)
	lexical.AddRule(NumberToken, num_rule)
	collector := NewErrorCollector()

	return &HulkInterpreter{
		_interpreter: NewInterpreter(lexer, lexical, collector),
	}
}

func (hi *HulkInterpreter) Execute(code string) {
	hi._interpreter.Execute(code)
}
