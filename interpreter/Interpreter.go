package interpreter

import (
	"fmt"

	. "hulk.com/app/compiler"
	. "hulk.com/app/lexer"
	. "hulk.com/app/lexical_analisys"
)

type Interpreter struct {
	lexer            ILexer
	lexical_analizer ILexicalAnalizer
	error_collector  IErrorCollector
	error_type_texts map[ErrorType]string
}

func NewInterpreter(lexer ILexer, lexical_anlizer ILexicalAnalizer, error_collector IErrorCollector) *Interpreter {
	e_t := make(map[ErrorType]string)
	e_t[Lexical] = "Lexical"
	return &Interpreter{
		lexer:            lexer,
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
	}
	for _, e := range interpreter.error_collector.Errors() {
		fmt.Println(interpreter.error_type_texts[e.Type()], "Error at line", e.Line(), "column", e.Column(), ": ", e.Message())
	}
}
