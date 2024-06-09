package parser

import "github.com/scatternoodle/wflang/lang/token"

// precedence levels for expressions parsing. Predence weighting in ascending order
// where 1 is the lowest precedence.
const (
	_ int = iota
	p_LOWEST
)

func (p *Parser) currentPrecedence() int {
	return p.lookupPrecedence(p.current.Type)
}

func (p *Parser) peekPrecedence() int {
	return p.lookupPrecedence(p.next.Type)
}

func (p *Parser) lookupPrecedence(t token.Type) int {
	pr, ok := p.precedenceMap[t]
	if !ok {
		return p_LOWEST
	}
	return pr
}

// DON'T USE THIS, get precedence map from the Parser instead. This function is called
// once during New(), and the map is stored in the parser, to avoid a package-level
// variable and also to avoid this function being called every time we need to look
// up the precedence of a token. This would result in a great many copies made of
// what could end up being a rather large map.
func precedenceMap() map[token.Type]int {
	return map[token.Type]int{}
}
