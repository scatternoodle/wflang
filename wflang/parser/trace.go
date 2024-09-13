package parser

import "strings"

type trace struct {
	lvl int
	s   *strings.Builder
}

func (t *trace) String() string {
	return t.s.String()
}

func (t *trace) trace(msg string) {
	t.lvl++
	t.writeMsg("BEGIN " + msg)
}

func (t *trace) untrace(msg string) {
	t.writeMsg("END " + msg)
	t.lvl--
}

func (t *trace) writeMsg(msg string) {
	ind := strings.Repeat("\t", t.lvl-1)
	t.s.WriteString(ind + msg + "\n")
}
