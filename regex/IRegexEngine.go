package regex

import (
	. "hulk.com/app/automaton"
	. "hulk.com/app/grammar"
)

// Regex engine interface
type IRegexEngine interface {
	// Return the automaton that recognizes the given grammar
	Regex(grammar IGrammar) (IAutomaton[rune], error)
}
