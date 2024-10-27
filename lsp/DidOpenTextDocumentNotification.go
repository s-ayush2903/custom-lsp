package lsp

// fields pertaining to a document
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#didOpenTextDocumentParams
type DidOpenTextDocumentParams struct {
    TextDocumentItem TextDocumentItem `json:"textDocument"`
}

// comprises the base Notification along with document textDocument infn
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didOpen
type DidOpenTextDocumentNotification struct {
    Notification
    DidOpenTextDocumentParams DidOpenTextDocumentParams `json:"params"`
}

