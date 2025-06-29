package automatonpackage

func IStateCode() string {
	return `package automaton

// Automaton state interface
type IState[T comparable] interface {
	// Identifier for the state
	ID() string
	// Says when the state is accepting or not
	IsAccepting() bool
	// Says if the state is fault
	IsFault() bool
	// Return true if the state has a transition for the given symbol
	HasTransition(symbol T) bool
	// Return the states that can be reached from this state throug an epsilon transition
	Epsilons() []IState[T]
	// Return the next state for the given symbol
	Next(symbol T) IState[T]
	// Adds a new transition for this state
	AddTransition(symbol T, state IState[T])
	// Adds an epsilon transition for this state
	Epsilon(state IState[T])
	// Gets the epsilon clousure of this state
	Clousure() []IState[T]
}
`
}
