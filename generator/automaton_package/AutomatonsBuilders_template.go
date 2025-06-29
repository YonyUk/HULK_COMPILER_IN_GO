package automatonpackage

func AutomatonsBuildersCode() string {
	return `package automaton

import (
	"errors"
	"fmt"
)

// Returns the automaton that recognizes the given sequence
// id: is used to name the start state of this automaton
func SequenceAutomaton[T comparable](sequence []T, id string) (IAutomaton[T], error) {
	if len(sequence) == 0 {
		return nil, errors.New("The sequence can't be empty")
	}
	start := NewState[T](id, false, false)
	states := []IState[T]{start}
	for i := 0; i < len(sequence); i++ {
		states = append(states, NewState[T](id+"-"+fmt.Sprint(i), i == len(sequence)-1, false))
		states[i].AddTransition(sequence[i], states[i+1])
	}
	return NewAutomaton(states[0], states, sequence), nil
}
`
}
