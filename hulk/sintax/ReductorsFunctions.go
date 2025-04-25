package sintax

import (
	. "hulk.com/app/ast"
)

func BinaryOperatorReductor(asts []IAST, new_symbol string) IAST {
	if _, garbage := asts[1].(*GarbageAST); garbage {
		return asts[1]
	}
	operator, _ := asts[1].(*BinaryAST)
	operator.Left = asts[0]
	operator.Right = asts[2]
	operator.UpdateSymbol(new_symbol)
	return operator
}

func AtomicReductor(asts []IAST, new_symbol string) IAST {
	asts[0].UpdateSymbol(new_symbol)
	return asts[0]
}

func InBettwenExtractorReductor(asts []IAST, new_symbol string) IAST {
	asts[1].UpdateSymbol(new_symbol)
	return asts[1]
}

func InFrontOperatorReductor(asts []IAST, new_symbol string) IAST {
	if _, garbage := asts[0].(*GarbageAST); garbage {
		return asts[0]
	}
	operator, _ := asts[0].(*InFrontOperatorAST)
	operator.Target = asts[1]
	operator.UpdateSymbol(new_symbol)
	return operator
}
