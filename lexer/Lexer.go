package lexer

import (
	. "hulk.com/app/automaton"
	. "hulk.com/app/grammar"
	. "hulk.com/app/regex"
	. "hulk.com/app/tokens"
)

type Lexer struct {
	tokens_grammars map[TokenType]IGrammar
	code            string
	text_readed     string
	text_pointer    int
	automatons      map[TokenType]IAutomaton[rune]
	token_extractor ITokenExtractor
	priorities      map[int]TokenType
	current_token   IToken
	last_token      IToken
	line            int
	column          int
}

func NewLexer() *Lexer {
	return &Lexer{
		tokens_grammars: make(map[TokenType]IGrammar),
		code:            "",
		text_readed:     "",
		text_pointer:    0,
		automatons:      make(map[TokenType]IAutomaton[rune]),
		token_extractor: nil,
		priorities:      make(map[int]TokenType),
		line:            1,
		column:          0,
	}
}

func (l *Lexer) AddTokenExpression(token_type TokenType, priority int, re IGrammar) {
	l.tokens_grammars[token_type] = re
	l.priorities[priority] = token_type
}

func (l *Lexer) LoadCode(code string) {
	l.text_pointer = 0
	l.text_readed = ""
	l.code = code
	l.line = 1
	l.column = 1
}

func (l *Lexer) Next() bool {
	err := l.build_automatons()
	if err != nil {
		panic(err)
	}
	next := len(l.code) > 0
	last_types := []TokenType{}
	walked := next
	if l.token_extractor == nil {
		l.token_extractor = NewTokenExtractor(l.priorities)
	}
	for walked && len(l.code) > 0 {
		current_types := []TokenType{}
		walked = false
		for token_type, automaton := range l.automatons {
			automaton.Walk(rune(l.code[l.text_pointer]))
			if !automaton.CurrentState().IsFault() {
				walked = true
				if automaton.CurrentState().IsAccepting() {
					current_types = append(current_types, token_type)
				}
			}
		}
		if walked {
			if l.code[l.text_pointer] == '\n' {
				l.column = 1
				l.line++
			}
			last_types = current_types
			l.text_readed += string(l.code[l.text_pointer])
			l.text_pointer++
			if l.text_pointer == len(l.code) {
				walked = false
			}
		}
		if !walked {

			if l.text_pointer == 0 {
				if l.code[l.text_pointer] != ' ' && l.code[l.text_pointer] != '\n' {
					l.current_token = l.token_extractor.GetToken(last_types, l.line, l.column, string(l.code[0]))
				}
				if l.code[l.text_pointer] == '\n' {
					l.column = 1
					l.line++
				} else {
					l.column++
				}
				l.code = l.code[1:]
				walked = true
			} else {
				l.current_token = l.token_extractor.GetToken(last_types, l.line, l.column, l.text_readed)
				l.code = l.code[l.text_pointer:]
				l.column += l.text_pointer
				l.text_pointer = 0
				l.text_readed = ""
			}
			for _, automaton := range l.automatons {
				automaton.Restart()
			}
		}
	}
	if l.last_token != nil && l.last_token.Line() == l.current_token.Line() && l.last_token.Column() == l.current_token.Column() {
		return false
	}
	l.last_token = l.current_token
	return next
}

func (l *Lexer) Current() IToken {
	return l.current_token
}

func (l *Lexer) build_automatons() error {
	regex_engine := NewRegexEngine()
	for k, g := range l.tokens_grammars {
		aut, err := regex_engine.Regex(g)
		if err != nil {
			return err
		}
		l.automatons[k] = aut
	}
	return nil
}
