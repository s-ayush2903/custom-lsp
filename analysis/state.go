package analysis

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
