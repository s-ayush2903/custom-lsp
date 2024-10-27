package lsp

type DefinitionRequest struct {
    Request
    DefinitionParams
}

type DefinitionParams struct {
    Params TextDocumentPositionParams `json:"params"`
}

type DefinitionResponse struct {
    Response
    Result Location `json:"result"`
}
