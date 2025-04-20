package parser

import (
	. "hulk.com/app/tools"
)

type ItemLR0Collection struct {
	id    string
	items []IItemLR0
}

func NewItemLR0Collection(items []IItemLR0) *ItemLR0Collection {
	items_ := MergeSort(items, func(a IItemLR0, b IItemLR0) int {
		if CompareString(a.ID(), b.ID()) < 0 {
			return -1
		}
		if CompareString(a.ID(), b.ID()) > 0 {
			return 1
		}
		return 0
	})
	id := items_[0].ID()
	for _, item := range items_ {
		id += "-" + item.ID()
	}
	return &ItemLR0Collection{
		id:    id,
		items: items_,
	}
}

func (collection *ItemLR0Collection) ID() string {
	return collection.id
}

func (collection *ItemLR0Collection) Items() []IItemLR0 {
	return collection.items
}
