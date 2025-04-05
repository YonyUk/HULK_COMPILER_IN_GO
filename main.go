package main

import (
	"fmt"

	. "hulk.com/app/automaton"
)

func main() {
	q0 := NewState[rune]("q0", true, false)
	q1 := NewState[rune]("q1", true, false)

	q0.AddTransition('0', q0)

	q1.AddTransition('1', q1)

	aut0 := NewAutomaton(q0, []IState[rune]{q0}, []rune{'0'})
	aut1 := NewAutomaton(q1, []IState[rune]{q1}, []rune{'1'})

	aut := aut0.Concat(aut1)

	// fmt.Println(aut.CurrentState().HasTransition('1'))

	chain := "00001"

	for _, char := range chain {
		fmt.Println(aut.Walk(char))
	}
	fmt.Println(aut.CurrentState().IsAccepting())

	// for _, state := range aut.States() {
	// 	fmt.Println(state.ID(), state)
	// }
}
