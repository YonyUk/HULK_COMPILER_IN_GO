package automaton

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
