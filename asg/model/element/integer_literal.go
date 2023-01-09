package element

import "math/big"

type IntegerLiteral interface {
	CobolDivisionElement

	GetValue() *big.Int
}
