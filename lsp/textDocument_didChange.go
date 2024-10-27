package lsp

/**
 * An event describing a change to a text document. If only a text is provided
 * it is considered to be the full content of the document.
 * https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentContentChangeEvent
 */
type TextDocumentContentChangeEvent struct {
	/**
	 * The new text of the whole document.
	 */
     Text string `json:"text"`
};

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#didChangeTextDocumentParams
type DidChangeTextDocumentParams struct {
    TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`
    Changes []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type DidChangeTextDocumentNotification struct {
    Notification
    Params DidChangeTextDocumentParams `json:"params"`
}
