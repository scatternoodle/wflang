package docstring

// FunctionDoc holds various markdown string segments that together comprise the
// documentation for a function.
type FunctionDoc struct {
	Name      string
	Signature string
	Returns   string
	Desc      string
	Params    []*ParamDoc
}

// String returns the full markdown string for the function, incorporating all
// segments in a standard doc comment format.
func (f *FunctionDoc) String() string {
	s := codeBlockStart +
		f.Signature + "\n" +
		codeBlockReturns + f.Returns + "\n" +
		codeBlockEnd +
		"### " + f.Name + "\n\n" +
		f.Desc + "\n\n"
	for _, param := range f.Params {
		s += param.String() + "\n\n"
	}
	return s
}

// ParamDoc holds the string segments that comprise a parameter docstring as part
// of a wider FunctionDoc. Also contains the label encoding for LSP signature help.
type ParamDoc struct {
	Name string
	// for lsp.SignatureHelp - substring position of the FunctionDoc signature [start,end]
	Label [2]int
	Type  string
	Desc  string
}

func (p *ParamDoc) String() string {
	return "@param " + p.Name + ": " + p.Type + " - " + p.Desc
}
