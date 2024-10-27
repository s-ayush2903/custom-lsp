package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentItem
type TextDocumentItem struct {
    /**
	 * The text document's URI.
	 */
    Uri string `json:"uri"`

    /**
	 * The text document's language identifier.
	 */
    LanguageId string `json:"languageId"`

    /**
	 * The version number of this document (it will increase after each
	 * change, including undo/redo).
	 */
    Version int `json:"version"`

    /**
	 * The content of the opened text document.
	 */
    Text string `json:"text"`
}

type TextDocumentIdentifier struct {
    Uri string `json:"uri"`
}

type VersionedTextDocumentIdentifier struct {
    Uri string `json:"uri"`
    Version int `json:"version"`
}
