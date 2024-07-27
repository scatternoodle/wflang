package object

import "github.com/scatternoodle/wflang/lang/types"

// Object is the base interface for any formula object in WFLang, and is the ultimate
// expression of a formula. The main requirement of an object is to have a type.
type Object interface {
	Type() types.Type
	Methods() []Function
}

type Number struct {
	Val float64
}

func (n Number) Type() types.Type { return types.T_NUMBER }

func (n Number) Methods() []Function { return nil }
