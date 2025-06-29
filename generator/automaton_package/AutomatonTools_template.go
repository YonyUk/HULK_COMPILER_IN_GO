package automatonpackage

func AutomatonToolsCode() string {
	return `package automaton

import (
	. "hulk.com/app/tools"
)

func SetStatesEquals[T comparable](states1 []IState[T], states2 []IState[T]) bool {
	if len(states1) != len(states2) {
		return false
	}
	for _, state1 := range states1 {
		exists := false
		for _, state2 := range states2 {
			if state2.ID() == state1.ID() {
				exists = true
				break
			}
		}
		if !exists {
			return false
		}
	}
	return true
}

func SetStatesIsInSetOfSetStates[T comparable](set_set_states [][]IState[T], set_states []IState[T]) bool {
	for _, set_states_base := range set_set_states {
		if SetStatesEquals(set_states_base, set_states) {
			return true
		}
	}
	return false
}

func CopyAutomaton[T comparable](aut IAutomaton[T]) IAutomaton[T] {
	states := copy_states(aut)
	start_index, _ := IndexOf(states, func(state IState[T]) bool { return state.ID() == aut.Start().ID() })
	return NewAutomaton(states[start_index], states, aut.Alphabet())
}

func copy_states[T comparable](automaton IAutomaton[T]) []IState[T] {
	states := []IState[T]{}
	transitions := make(map[string]map[T]string)
	epsilons := make(map[string][]string)
	for _, state := range automaton.States() {
		states = append(states, NewState[T](state.ID(), state.IsAccepting(), state.IsFault()))
		transitions[state.ID()] = make(map[T]string)
		for _, ep := range state.Epsilons() {
			if _, ok := epsilons[state.ID()]; !ok {
				epsilons[state.ID()] = []string{}
			}
			epsilons[state.ID()] = append(epsilons[state.ID()], ep.ID())
		}
		for _, symbol := range automaton.Alphabet() {
			if !state.HasTransition(symbol) {
				continue
			}
			transitions[state.ID()][symbol] = state.Next(symbol).ID()
		}
	}
	for i := 0; i < len(states); i++ {
		if _, ok := epsilons[states[i].ID()]; ok {
			for _, id := range epsilons[states[i].ID()] {
				ep_index, _ := IndexOf(states, func(state IState[T]) bool { return state.ID() == id })
				states[i].Epsilon(states[ep_index])
			}
		}
		for key, value := range transitions[states[i].ID()] {
			s_index, _ := IndexOf(states, func(state IState[T]) bool { return state.ID() == value })
			states[i].AddTransition(key, states[s_index])
		}
	}
	return states
}
`
}
