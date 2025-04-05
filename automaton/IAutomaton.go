package automaton

// Automaton interface
type IAutomaton[T comparable] interface {
	// The initial state for this automaton
	Start() IState[T]
	// The finals state of this automaton
	Finals() []IState[T]
	// The states of this automaton
	States() []IState[T]
	// Return true if the automanton is deterministic
	IsDeterministic() bool
	// Gets the deterministic version of this automaton
	ToDeterministic() IAutomaton[T]
	// Move one step by the given symbol
	Walk(symbol T) error
	// Gets the current state
	CurrentState() IState[T]
	// Adds a new state to the automaton
	AddState(state IState[T])
	// Return the alphabet of this automaton
	Alphabet() []T
}
