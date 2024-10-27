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
