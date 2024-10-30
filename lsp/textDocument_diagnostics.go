package lsp

type PublishDiagnosticsNotification struct {
    Notification
    Params PublishDiagnosticParams `json:"params"`
}

type PublishDiagnosticParams struct {
    URI string 
    Diagnostics []Diagnostic `json:"diagnostics"`
}

type Diagnostic struct {
    Source string `json:"source"`
    Message string `json:"message"`
    Serverity int `json:"severity"`
    Range Range `json:"range"`
}

