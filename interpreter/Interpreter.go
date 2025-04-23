package interpreter

import (
	"fmt"

	. "hulk.com/app/compiler"
	. "hulk.com/app/lexer"
	. "hulk.com/app/lexical_analisys"
	. "hulk.com/app/parser"
	. "hulk.com/app/tokens"
)

type Interpreter struct {
	lexer            ILexer
	parser           IParser
	lexical_analizer ILexicalAnalizer
	error_collector  IErrorCollector
	error_type_texts map[ErrorType]string
}

func NewInterpreter(lexer ILexer, lexical_anlizer ILexicalAnalizer, parser IParser, error_collector IErrorCollector) *Interpreter {
	e_t := make(map[ErrorType]string)
	e_t[Lexical] = "Lexical"
	e_t[Gramatical] = "Gramatical"
	e_t[Semantic] = "Semantic"
	return &Interpreter{
		lexer:            lexer,
		parser:           parser,
		lexical_analizer: lexical_anlizer,
		error_collector:  error_collector,
		error_type_texts: e_t,
	}
}

func (interpreter *Interpreter) Execute(code string) {
	interpreter.lexer.LoadCode(code)
	for interpreter.lexer.Next() {
		err := interpreter.lexical_analizer.Check(interpreter.lexer.Current())
		if err != nil {
			interpreter.error_collector.AddError(err)
		}
		interpreter.parser.Parse(interpreter.lexer.Current(), interpreter.error_collector)
	}
	EOF := NewToken(interpreter.lexer.Current().Line(), interpreter.lexer.Current().Column()+len(interpreter.lexer.Current().Text()), interpreter.parser.EndMarker(), SymbolToken)
	interpreter.parser.Parse(EOF, interpreter.error_collector)
	if len(interpreter.error_collector.Errors()) == 0 {
		code_result := interpreter.parser.AST().Eval(nil, interpreter.error_collector)
		if len(interpreter.error_collector.Errors()) == 0 {
			fmt.Println(code_result)
		} else {
			for _, e := range interpreter.error_collector.Errors() {
				fmt.Println(interpreter.error_type_texts[e.Type()], "Error at line", e.Line(), "column", e.Column(), ": ", e.Message())
			}
		}
	}
}
