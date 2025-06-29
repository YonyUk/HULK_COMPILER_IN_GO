package automatonpackage

func AutomatonCode() string {
	return `package automaton

import (
	"errors"

	. "hulk.com/app/tools"
)

type Automaton[T comparable] struct {
	start    IState[T]
	states   []IState[T]
	current  IState[T]
	alphabet []T
}

func NewAutomaton[T comparable](start IState[T], states []IState[T], alphabet []T) *Automaton[T] {
	return &Automaton[T]{
		start:    start,
		states:   states,
		current:  start,
		alphabet: alphabet,
	}
}

func (a *Automaton[T]) Start() IState[T] {
	return a.start
}

func (a *Automaton[T]) Finals() []IState[T] {
	return FiltBy(a.states, func(state IState[T]) bool {
		return state.IsAccepting()
	})
}

func (a *Automaton[T]) States() []IState[T] {
	return a.states
}

func (a *Automaton[T]) IsDeterministic() bool {
	counter := Count(a.states, func(state IState[T]) bool {
		return len(state.Epsilons()) > 0
	})
	return counter == 0
}

func (a *Automaton[T]) ToDeterministic() IAutomaton[T] {
	initial_id, states, transitions := a.toDeterministic()
	States := []IState[T]{}
	var InitialState IState[T]
	for _, _states := range states {
		new_state := a.make_state_from_states(_states)
		if new_state.ID() == initial_id {
			InitialState = new_state
		}
		States = append(States, new_state)
	}
	for id, dict := range transitions {
		index, _ := IndexOf(States, func(s IState[T]) bool { return s.ID() == id })
		for symbol, state_id := range dict {
			index2, _ := IndexOf(States, func(s IState[T]) bool { return s.ID() == state_id })
			States[index].AddTransition(symbol, States[index2])
		}
	}
	return NewAutomaton(InitialState, States, a.alphabet)
}

func (a *Automaton[T]) Restart() {
	a.current = a.start
}

func (a *Automaton[T]) toDeterministic() (string, [][]IState[T], map[string]map[T]string) {
	// calculando las clausuras
	initial, clousures := a.get_clousures()
	states := [][]IState[T]{initial}
	transitions := make(map[string]map[T]string)
	state_added := true

	for state_added {
		state_added = false
		for _, state := range states {

			a_state := a.make_state_from_states(state)
			if _, ok := transitions[a_state.ID()]; !ok {
				transitions[a_state.ID()] = make(map[T]string)
			}

			for _, symbol := range a.alphabet {
				move := []IState[T]{}
				for _, st := range state {
					if !st.HasTransition(symbol) {
						continue
					}
					next := st.Next(symbol)
					move = append(move, next)
					for _, e_st := range clousures[next.ID()] {
						if e_st.ID() != next.ID() {
							move = append(move, e_st)
						}
					}
				}
				if len(move) == 0 {
					continue
				}
				if !SetStatesIsInSetOfSetStates(states, move) {
					states = append(states, move)
					state_added = true
					transitions[a_state.ID()][symbol] = a.make_state_from_states(move).ID()
					break
				} else {
					transitions[a_state.ID()][symbol] = a.make_state_from_states(move).ID()
				}
			}
			if state_added {
				break
			}
		}
	}
	return a.make_state_from_states(initial).ID(), states, transitions
}

func (a *Automaton[T]) CurrentState() IState[T] {
	return a.current
}

func (a *Automaton[T]) Walk(symbol T) error {
	if a.current.IsFault() {
		return nil
	}
	if !a.IsDeterministic() {
		return errors.New("Automaton is not deterministic, make it deterministic first")
	}
	a.current = a.current.Next(symbol)
	return nil
}

func (a *Automaton[T]) make_state_from_states(states []IState[T]) IState[T] {
	id := states[0].ID()
	for i := 1; i < len(states); i++ {
		id += "-" + states[i].ID()
	}
	accepting := Count(states, func(state IState[T]) bool { return state.IsAccepting() }) > 0
	fault := Count(states, func(state IState[T]) bool { return state.IsFault() }) == len(states)
	return NewState[T](id, accepting, fault)
}

func (a *Automaton[T]) get_clousures() ([]IState[T], map[string][]IState[T]) {
	result := make(map[string][]IState[T])
	for _, state := range a.states {
		result[state.ID()] = state.Clousure()
	}
	return a.start.Clousure(), result
}

func (a *Automaton[T]) AddState(state IState[T]) {
	a.states = append(a.states, state)
}

func (a *Automaton[T]) Alphabet() []T {
	return a.alphabet
}
`
}
