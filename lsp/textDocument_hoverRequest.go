package lsp

type HoverRequest struct {
    Request
    HoverParams
}

type HoverParams struct {
    Params TextDocumentPositionParams `json:"params"`
}

type HoverResponse struct {
    Response
    Result HoverResult `json:"result"`
}

type HoverResult struct {
    Contents string `json:"contents"`
}