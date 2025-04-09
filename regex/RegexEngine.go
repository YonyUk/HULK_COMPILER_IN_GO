package regex

import (
	. "hulk.com/app/automaton"
	. "hulk.com/app/grammar"
	// . "hulk.com/app/tools"
)

type RegexEngine struct {
}

func NewRegexEngine() *RegexEngine {
	return &RegexEngine{}
}

func (re *RegexEngine) Regex(g IGrammar) (IAutomaton[rune], error) {
	return nil, nil
}
