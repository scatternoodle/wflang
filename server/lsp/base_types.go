package lsp

import "fmt"

type (
	Int  int32
	Uint uint32
)

// Decimals in LSP are expressed in Range Notation
type Decimal struct {
	Start float64
	End   float64
}

// DecimalFromFloat gives you a Decimal from a single float. Just puts f in both
// Start and End. Use this to represent a single floating point number in LSP.
func DecimalFromFloat(f float64) Decimal {
	return Decimal{f, f}
}

// DecimalFromRange returns a Decimal with the start and end floats specified.
// Use this if you truly need to express a range in LSP.
func DecimalFromRange(start, end float64) Decimal {
	return Decimal{start, end}
}

// Returns the Decimal represented in LSP-compliant range notation. Use this to
// insert Decimals into LSP messages.
func (d Decimal) String() string {
	return fmt.Sprintf("[%v, %v]", d.Start, d.End)
}
