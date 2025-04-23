package sintax

import (
	. "hulk.com/app/ast"
)

func BinaryOperatorReductor(asts []IAST, new_symbol string) IAST {
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
