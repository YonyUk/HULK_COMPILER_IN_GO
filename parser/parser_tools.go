package parser

import (
	"fmt"
	"os"
	"path/filepath"

	. "hulk.com/app/grammar"
	. "hulk.com/app/tools"
)

// Gets the collection of items lr(0) associated to a production
func GetLR0Collection(head IGrammarSymbol, production []IGrammarSymbol) IItemLR0Collection {
	result := []IItemLR0{}
	for i := 0; i < len(production)+1; i++ {
		result = append(result, NewItemLR0(head, production[:i], production[i:]))
	}
	return NewItemLR0Collection(result)
}

// Gets the clousure of a given set of items lr(0)
func ItemLR0Clousure(items IItemLR0Collection, g IGrammar) IItemLR0Collection {
	result := []IItemLR0{}
	for _, item := range items.Items() {
		result = append(result, item)
	}
	change := true
	for change {
		change = false
		for _, item := range result {
			if len(item.RightTail()) > 0 {
				productions := g.GetProductions(item.RightTail()[0])
				for _, production := range productions {
					new_item := NewItemLR0(item.RightTail()[0], []IGrammarSymbol{}, production)
					_, err := IndexOf(result, func(i IItemLR0) bool { return i.ID() == new_item.ID() })
					if err != nil {
						change = true
						result = append(result, new_item)
					}
				}
			}
		}
	}
	return NewItemLR0Collection(result)
}

func GOTO(I IItemLR0Collection, symbol IGrammarSymbol, g IGrammar) IItemLR0Collection {
	items := []IItemLR0{}
	for _, item := range I.Items() {
		if len(item.RightTail()) > 0 && item.RightTail()[0].Symbol() == symbol.Symbol() {
			new_item := NewItemLR0(item.Head(), append(item.LeftTail(), item.RightTail()[0]), item.RightTail()[1:])
			items = append(items, new_item)
		}
	}
	if len(items) > 0 {
		result := ItemLR0Clousure(NewItemLR0Collection(items), g)
		return result
	}
	return nil
}

func GetCanonicalLR0Collection(g IGrammar) []IItemLR0Collection {
	old_start := g.GetProductions(g.StartSymbol())[0][0]
	start_item := NewItemLR0(g.StartSymbol(), []IGrammarSymbol{}, []IGrammarSymbol{old_start})
	sets := []IItemLR0Collection{ItemLR0Clousure(NewItemLR0Collection([]IItemLR0{start_item}), g)}
	symbols := []IGrammarSymbol{}
	for _, nt := range g.NonTerminals() {
		if nt.Symbol() == g.StartSymbol().Symbol() {
			continue
		}
		symbols = append(symbols, nt)
	}
	for _, ter := range g.Terminals() {
		symbols = append(symbols, ter)
	}
	change := true
	for change {
		change = false
		for _, collection := range sets {
			for _, symbol := range symbols {
				goto_ := GOTO(collection, symbol, g)
				if goto_ != nil {
					_, err := IndexOf(sets, func(i IItemLR0Collection) bool { return i.ID() == goto_.ID() })
					if err != nil {
						sets = append(sets, goto_)
						change = true
					}
				}
			}
		}
	}
	return sets
}

func ShowCollection(collection IItemLR0Collection) {
	for _, item := range collection.Items() {
		left := Map(item.LeftTail(), func(s IGrammarSymbol) string { return s.Symbol() })
		right := Map(item.RightTail(), func(s IGrammarSymbol) string { return s.Symbol() })
		fmt.Println(item.Head().Symbol(), "---->", left, right)
	}
	fmt.Println()
	fmt.Println()
}

func DumpParser(parser ParserSLR, path string) {
	if path == "" {
		path, _ = os.Getwd()
	}
	status, err := os.Stat(path)
	if err != nil {
		panic("The given path does not exists")
	}
	if !status.IsDir() {
		panic("The given path is not a directory")
	}
	path_ := filepath.Join(path, "PARSER.json")
	text := "{\"START\":\"" + parser.StartState() + "\","
	text += "\"ENDMARKER\": \"" + parser.EndMarker() + "\","
	text += "\"ACTION\":{"
	for k1, val := range parser.ActionTable() {
		text += "\"" + k1 + "\":{"
		for k2, act := range val {
			switch act.Action {
			case SHIFT:
				text += "\"" + k2 + "\": [\"SHIFT\", \"" + act.NextState + "\"],"
				break
			case REDUCE:
				text += "\"" + k2 + "\": [\"REDUCE\"],"
				break
			case ACCEPT:
				text += "\"" + k2 + "\": [\"ACCEPT\"],"
				break
			default:
				break
			}
		}
		text = text[:len(text)-1]
		text += "},"
	}
	text = text[:len(text)-1]
	text += "}, \"REDUCE\":{"
	for k1, val := range parser.ReduceTable() {
		text += "\"" + k1 + "\":{"
		for k2, red := range val {
			symbols := "["
			for _, s := range red.Symbols {
				symbols += "\"" + s + "\","
			}
			symbols = symbols[:len(symbols)-1] + "]"
			text += "\"" + k2 + "\":{\"head\": \"" + red.NewSymbol + "\",\"symbols\":" + symbols + "},"
		}
		text = text[:len(text)-1]
		text += "},"
	}
	text = text[:len(text)-1]
	text += "}}"
	os.WriteFile(path_, []byte(text), 0644)
}
