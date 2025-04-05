package automaton

// Automaton state implementation
type State[T comparable] struct {
	transitions map[T]IState[T]
	epsilons    []IState[T]
	id          string
	accepting   bool
	fault       bool
}

// Return a new automaton state
func NewState[T comparable](id string, accepting bool, fault bool) *State[T] {
	return &State[T]{
		transitions: make(map[T]IState[T]),
		epsilons:    []IState[T]{},
		id:          id,
		accepting:   accepting,
		fault:       fault,
	}
}

func (s *State[T]) ID() string {
	return s.id
}

func (s *State[T]) IsAccepting() bool {
	return s.accepting
}

func (s *State[T]) IsFault() bool {
	return s.fault
}

func (s *State[T]) HasTransition(symbol T) bool {
	if s.fault {
		return false
	}
	_, ok := s.transitions[symbol]
	return ok
}

func (s *State[T]) Epsilons() []IState[T] {
	if s.fault {
		return []IState[T]{}
	}
	return s.epsilons
}

func (s *State[T]) AddTransition(symbol T, state IState[T]) {
	s.transitions[symbol] = state
}

func (s *State[T]) Epsilon(state IState[T]) {
	s.epsilons = append(s.epsilons, state)
}

func (s *State[T]) Next(symbol T) IState[T] {
	if s.IsFault() {
		return s
	}
	if next, ok := s.transitions[symbol]; ok {
		return next
	}
	return NewState[T]("FAULT", false, true)
}

func (s *State[T]) Clousure() []IState[T] {
	clousure := []IState[T]{s}
	for _, item := range s.epsilons {
		clousure = append(clousure, item)
	}
	return clousure
}
