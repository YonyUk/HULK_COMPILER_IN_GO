package automaton

import (
	. "hulk.com/app/tools"
)

// Returns the automaton that recognizes the concatenation of both languages
func Concat[T comparable](aut1 IAutomaton[T], aut2 IAutomaton[T]) IAutomaton[T] {
	states := []IState[T]{}
	states1 := copy_states(aut1)
	states2 := copy_states(aut2)
	alphabet := aut1.Alphabet()
	for _, symbol := range aut2.Alphabet() {
		if _, err := IndexOf(alphabet, func(s T) bool { return s == symbol }); err != nil {
			alphabet = append(alphabet, symbol)
		}
	}
	for i := 0; i < len(states1); i++ {
		states = append(states, states1[i])
	}
	for i := 0; i < len(states2); i++ {
		states = append(states, states2[i])
	}
	s_index_to, _ := IndexOf(states, func(s IState[T]) bool { return s.ID() == aut2.Start().ID() })
	for _, state := range aut1.Finals() {
		s_index_from, _ := IndexOf(states, func(s IState[T]) bool { return s.ID() == state.ID() })
		states[s_index_from].Epsilon(states[s_index_to])
	}
	start_index, _ := IndexOf(states, func(s IState[T]) bool { return s.ID() == aut1.Start().ID() })
	return NewAutomaton(states[start_index], states, alphabet).ToDeterministic()
}

func Union[T comparable](aut1 IAutomaton[T], aut2 IAutomaton[T]) IAutomaton[T] {
	start := NewState[T](aut1.Start().ID()+"-"+aut2.Start().ID()+"-"+"union", false, false)
	states := []IState[T]{start}
	states1 := copy_states(aut1)
	states2 := copy_states(aut2)
	alphabet := aut1.Alphabet()
	for _, symbol := range aut2.Alphabet() {
		if _, err := IndexOf(alphabet, func(s T) bool { return s == symbol }); err != nil {
			alphabet = append(alphabet, symbol)
		}
	}
	for i := 0; i < len(states1); i++ {
		states = append(states, states1[i])
	}
	for i := 0; i < len(states2); i++ {
		states = append(states, states2[i])
	}
	s1_index, _ := IndexOf(states, func(s IState[T]) bool { return s.ID() == aut1.Start().ID() })
	s2_index, _ := IndexOf(states, func(s IState[T]) bool { return s.ID() == aut2.Start().ID() })
	start.Epsilon(states[s1_index])
	start.Epsilon(states[s2_index])
	return NewAutomaton(start, states, alphabet).ToDeterministic()
}
