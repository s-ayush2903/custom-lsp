package analysis

import (
	"fmt"
	"log"
	"lsp-go/lsp"
	"strings"
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

func (state *State) CodeAction(id int, filePath string) lsp.CodeActionResponse {
    contents := state.Documents[filePath]
    codeActions := []lsp.CodeAction{}
    for row, line := range strings.Split(contents, "\n") {
        idx := strings.Index(line, "VS Code")
        if idx > -1 {
            replaceChange := map[string][]lsp.TextEdit{}
            replaceChange[filePath] = []lsp.TextEdit{
                {
                    Range: LineRange(row, idx, idx + len("VS Code")),
                    NewText: "nvim",
                },
            }

            codeActions = append(codeActions, lsp.CodeAction{
                Title: "Replace VS C*de with a superior editor ",
                Edit: &lsp.WorkspaceEdit{Changes: replaceChange},
            })

            censorChange := map[string][]lsp.TextEdit{}
            censorChange[filePath] = []lsp.TextEdit{
                {
                    Range: LineRange(row, idx, idx + len("VS Code")),
                    NewText: "VS C*de",
                },
            }

            codeActions = append(codeActions, lsp.CodeAction{
                Title: "Censor to VS C*de",
                Edit: &lsp.WorkspaceEdit{Changes: censorChange},
            })
        }
    }

    codeActionResponse := lsp.CodeActionResponse{
        Response: lsp.Response{
            RPC: "2.0",
            ID: &id,
        },
        Result: codeActions, // curly brackets are for creating a NEW instance
    }

    return codeActionResponse
}

func (state *State) Completion(logger *log.Logger, id int, filePath string) lsp.CompletionResponse {
    completions := []lsp.CompletionItem{
        {
            Label: "sample trigger",
            Detail: "This is the very default granular detail about this completion",
            Documentation: "People rarely read documentation for theses pieces of codes, if you want it more then better search it over the internet",
        },
    }

    completionsResponse := lsp.CompletionResponse{
        Response: lsp.Response{
            RPC: "2.0",
            ID: &id,
        },
        Completions: completions,
    }

    // logger.Printf("COMPLETION RESPONSE: %s", completions[0])

    return completionsResponse
}

func LineRange (line, start, end int) lsp.Range {
    return lsp.Range{
        Start: lsp.Position{Line: line, Character: start},
        End: lsp.Position{Line: line, Character: end},
    }
}
