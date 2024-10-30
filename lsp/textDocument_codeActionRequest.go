package lsp

type CodeActionRequest struct {
    Request
    Params CodeActionParams `json:"params"`
}

type CodeActionContext struct {
}

type CodeActionParams struct {
    TextDocumentIdentifier `json:"textDocument"`
    Range `json:"range"`
    Context CodeActionContext `json:"context"`
}

type CodeActionResponse struct {
    Response
    Result []CodeAction `json:"result"`
}

type CodeAction struct {
    Title string `json:"title"`
    Edit *WorkspaceEdit `json:"edit,omitempty"`
    Command *Command `json:"command,omitempty"`
}

type Kind struct {
    QuickFix string
    Refactor string
}

type Diagnostic struct {
}

type Command struct {
    Title string `json:"title"`
    Command string `json:"command"`
    Arguments []interface{} `json:"arguments,omitempty"`
}
