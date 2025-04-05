package automaton

func (a *Automaton[T]) Concat(other IAutomaton[T]) IAutomaton[T] {
	states := a.states
	new_alphabet := a.alphabet
	for _, symbol := range other.Alphabet() {
		new_alphabet = append(new_alphabet, symbol)
	}
	for _, state := range other.States() {
		states = append(states, state)
	}
	for _, state := range a.Finals() {
		state.Epsilon(other.Start())
	}
	return NewAutomaton(a.start, states, new_alphabet).ToDeterministic()
}
