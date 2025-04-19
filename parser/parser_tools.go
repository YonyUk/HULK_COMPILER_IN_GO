package parser

import (
	. "hulk.com/app/grammar"
)

func GetLR0Collection(head IGrammarSymbol, production []IGrammarSymbol) []IItemLR0 {
	result := []IItemLR0{}
	for i := 0; i < len(production)+1; i++ {
		result = append(result, NewItemLR0(head, production[:i], production[i:]))
	}
	return result
}

func ItemLR0Clousure(items []IItemLR0, g IGrammar) []IItemLR0 {
	result := []IItemLR0{}
	for _, item := range items {
		result = append(result, item)
	}
	change := true
	for change {
		change = false
		for _, item := range result {
			if len(item.RightTail()) > 0 {
				productions := g.GetProductions(item.RightTail()[0])
				for _, production := range productions {
					result = append(result, NewItemLR0(item.RightTail()[0], []IGrammarSymbol{}, production))
					change = true
				}
			}
		}
	}
	return result
}

func CompareItemLR0(a IItemLR0, b IItemLR0) bool {
	if a.Head().Symbol() != b.Head().Symbol() {
		return false
	}
	if len(a.LeftTail()) != len(b.LeftTail()) {
		return false
	}
	if len(a.RightTail()) != len(b.RightTail()) {
		return false
	}
	return true
}
