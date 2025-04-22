package sintax

import (
	. "hulk.com/app/grammar"
)

var HulkGrammar IGrammar

func init() {
	HulkGrammar = ArithMeticGrammar
}
