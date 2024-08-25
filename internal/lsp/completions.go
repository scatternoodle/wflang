package lsp

import "github.com/scatternoodle/wflang/internal/jrpc2"

// Text Document Completion
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_completion

// ComletionOptions
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionOptions
type CompletionOptions struct {
	TriggerCharacters   []string               `json:"triggerCharacters,omitempty"`
	AllCommitCharacters []string               `json:"allCommitCharacters,omitempty"`
	ResolveProvider     bool                   `json:"resolveProvider,omitempty"`
	CompletionItem      *CompletionItemOptions `json:"completionItem,omitempty"`
}

type CompletionItemOptions struct {
	LabelDetailsSupport bool `json:"labelDetailsSupport,omitempty"`
}

// CompletionRequest
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_completion
type CompletionRequest struct {
	jrpc2.Request
	Params CompletionParams `json:"params"`
}

// CompletionParams
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionParams
type CompletionParams struct {
	TextDocumentPositionParams
	Context *CompletionContext `json:"context,omitempty"`
}

// CompletionContext
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionContext
type CompletionContext struct {
	TriggerKind CompletionTriggerKind `json:"triggerKind"`
	TriggerChar string                `json:"triggerCharacter,omitempty"`
}

// CompletionTriggerKind
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionTriggerKind
type CompletionTriggerKind int

const (
	_ CompletionTriggerKind = iota
	CompTriggerInvoked
	CompTriggerTriggerCharacter
	CompTriggerIncomplete
)

type CompletionResponse struct {
	jrpc2.Response
	Result []CompletionItem `json:"result"`
}

// CompletionItem
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItem
type CompletionItem struct {
	// The first and most prominent text to display for the completion item.
	Label string `json:"label"`

	// Additional details to display after the label
	LabelDetails *CompletionItemLabelDetails `json:"labelDetails,omitempty"`
	Kind         CompletionItemKind          `json:"kind,omitempty"`
	Tags         []CompletionItemTag         `json:"tags,omitempty"`

	// Additional info, e.g. type. For beefier stuff might want to use Documentation
	Detail string `json:"detail,omitempty"`

	// Doc comment string
	Documentation MarkupContent `json:"documentation,omitempty"`

	// Deprecated (is) Deprecated: use Tags instead
	Deprecated bool `json:"deprecated,omitempty"`

	// Indicates to client that they should select this item when presenting to the user. Client can only select
	// one, and will make the ultimate decision.
	Preselect bool `json:"preselect,omitempty"`

	// If omitted, Label is used instead.
	SortText string `json:"sortText,omitempty"`

	// If omitted, Label is used instead.
	FilterText string `json:"filterText,omitempty"`

	// Simply gives the client the text to insert without additional info - subject to client-side interpretation.
	// For a more deterministic completion, use TextEdit instead.
	InsertText string `json:"insertText,omitempty"`

	// Applies to both InsertText and TextEdit.NewText. If omitted, InsFormatPlainText is used.
	InsertFormat InsertTextFormat `json:"insertTextFormat,omitempty"`

	// Affects how whitespaces are applied. Default is as-is.
	InsertTextMode InsertTextMode `json:"insertTextMode,omitempty"`
	TextEdit       *TextEdit      `json:"textEdit,omitempty"`

	// You can put additional ancillary edits here, e.g. adding an import statement to the top of a file.
	AddtlEdits  []TextEdit `json:"additionalTextEdits,omitempty"`
	CommitChars []string   `json:"commitCharacters,omitempty"`

	// Optional editor command to be executed AFTER the completion is inserted. For additional insertions use
	// AddtlEdits instead.
	Command *Command `json:"command,omitempty"`
	Data    any      `json:"data,omitempty"`
}

// CompletionItemLabelDetails
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItemLabelDetails
type CompletionItemLabelDetails struct {
	// Displays directly after the label, without space. Typical use case is a function signature.
	Detail string `json:"detail,omitempty"`
	// Displays after the detail, even less prominent.
	Description string `json:"description,omitempty"`
}

// CompletionItemKind
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItemKind
type CompletionItemKind int

const (
	_ CompletionItemKind = iota
	CompItemText
	CompItemMethod
	CompItemFunc
	CompItemConstructor
	CompItemField
	CompItemvar
	CompItemClass
	CompItemInterface
	CompItemModule
	CompItemProperty
	CompItemUnit
	CompItemValue
	CompItemEnum
	CompItemKeyword
	CompItemSnippet
	CompItemColor
	CompItemFile
	CompItemReference
	CompItemFolder
	CompItemEnumMember
	CompItemConstant
	CompItemStruct
	CompItemEvent
	CompItemOperator
	CompItemTypeParam
)

// CompletionItemTag
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItemTag
type CompletionItemTag int

const (
	_ CompletionItemTag = iota
	CompTagDeprecated
)

// InsertTextFormat
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#insertTextFormat
type InsertTextFormat int

const (
	_ InsertTextFormat = iota
	InsFormatPlainText
	InsFormatSnippet
)

// InsertTextMode
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#insertTextMode
type InsertTextMode int

const (
	_ InsertTextMode = iota
	InsModeAsIs
	InsModeAdjustIndent
)
