package lsp

type CompletionRequest struct {
    Request
    CompletionParams
}

type CompletionParams struct {
    Params TextDocumentPositionParams `json:"params"`
}

type CompletionResponse struct {
    Response
    Completions []CompletionItem `json:"result"`
}

type CompletionItem struct {
    Label string `json:"label"`
    Detail string `json:"detail"`
    Documentation string `json:"documentation"`
}
