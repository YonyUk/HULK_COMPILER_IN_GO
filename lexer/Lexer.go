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
}

func NewLexer() *Lexer {
	return &Lexer{
		tokens_grammars: make(map[TokenType]IGrammar),
		code:            "",
		text_readed:     "",
		text_pointer:    0,
		automatons:      make(map[TokenType]IAutomaton[rune]),
	}
}

func (l *Lexer) AddToken(token_type TokenType, re IGrammar) {
	l.tokens_grammars[token_type] = re
}

func (l *Lexer) Tokenize() (<-chan IToken, error) {
	ch := make(chan IToken)
	err := l.build_automatons()
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		tokens_types_selection := []TokenType{}
		for l.text_pointer < len(l.code) {
			walked := false
			tokens_types := []TokenType{}
			for token_type, aut := range l.automatons {
				aut.Walk(rune(l.code[0]))
				if !aut.CurrentState().IsFault() {
					walked = true
					tokens_types = append(tokens_types, token_type)
				}
			}
			if walked {
				l.text_readed += string(l.code[0])
				l.code = l.code[1:len(l.code)]
				l.text_pointer++
				copy(tokens_types_selection, tokens_types)
			} else {
				token, err := l.token_extractor.GetToken(tokens_types_selection, 0, 0, l.text_readed)
				if err != nil {
					panic(err)
				}
				ch <- token
				l.text_pointer = 0
				for _, aut := range l.automatons {
					aut.Restart()
				}
			}
		}
	}()
	return ch, nil
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
