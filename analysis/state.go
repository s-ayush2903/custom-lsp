package analysis

import (
	"fmt"
	"lsp-go/lsp"
)

// mapping of all the documents
type State struct {
    // map of document uris to their content
    Documents map[string]string
}

// returns an empty map on new state request
func NewState() State {
    return State {
        Documents: map[string]string{},
    }
}

func (state *State) OpenDocument(path string, contents string) {
    state.Documents[path] = contents
}

func (state *State) UpdateDocument(path string, contents string) {
    state.Documents[path] = contents
}

func (state *State) Hover(id int, path string) lsp.HoverResponse {
    fileContent := state.Documents[path]
    return lsp.HoverResponse{
        Response: lsp.Response{
            RPC: "2.0",
            ID: &id,
        },
        Result: lsp.HoverResult{
            Contents: fmt.Sprintf("File at %s with %d characters", path, len(fileContent)),
        },
    }
}

func (state *State) Definition(id int, path string, position lsp.Position) lsp.DefinitionResponse {
    _ = lsp.Position{
        Line: position.Line - 1,
        Character: 0,
    }
    return lsp.DefinitionResponse{
        Response: lsp.Response{
            RPC: "2.0",
            ID: &id,
        },
        Result: lsp.Location{
            Uri: path,
            Range: lsp.Range{Start: lsp.Position{Line: position.Line - 1, Character: 0}, End: lsp.Position{Line: position.Line - 1, Character: 0}},
            // Range: lsp.Range{Start: x, End: x},
        },
    }
}
