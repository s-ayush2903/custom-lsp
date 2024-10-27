package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize
type InitializeRequest struct {
    Request
    InitializeRequestParams InitializeRequestParams `json:"params"`
}

type ClientInfo struct {
    Name string `json:"name"`
    Version string `json:"version,omitempty"`
}

type InitializeRequestParams struct {
    ProcessId int `json:"processId,omitempty"`
    ClientInfo *ClientInfo `json:"clientInfo"`
    Locale string `json:"locale"`
}

type InitializeResponse struct {
    Response
    Result InitializeResult `json:"result"`
}

type InitializeResult struct {
    Capabilities ServerCapabilities `json:"capabilities"`
    ServerInfo ServerInfo `json:"serverInfo"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#serverCapabilities
type ServerCapabilities struct {
    TextDocumentSyncKind int `json:"textDocumentSync"`
}

type ServerInfo struct {
    Name string `json:"name"`
    Version string `json:"version,omitempty"`
}

func NewInitializeResponse(id int) InitializeResponse {
    return InitializeResponse{
        Response: Response{
            RPC: "2.0",
            ID: &id,
        },
        Result: InitializeResult{
            Capabilities: ServerCapabilities{
                TextDocumentSyncKind: 1,
            },
            ServerInfo: ServerInfo{
                Name: "custom-lsp",
                Version: "0.0.0.0.beta-1",
            },
        },
    }
}
