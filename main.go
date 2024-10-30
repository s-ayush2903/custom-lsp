package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"lsp-go/analysis"
	"lsp-go/lsp"
	"lsp-go/rpc"
	"os"
)

func main() {
    logger := getLogger("/Users/ayush/Desktop/learn/lsp-go/logs.txt")
    logger.Printf("logger initiated bruv..")

    logger.Printf("hiii");

    bufferScanner := bufio.NewScanner(os.Stdin);

    state := analysis.NewState();

    // split using the custom defined splitter function
    bufferScanner.Split(rpc.Split)
    var idx int = 1
    logger.Println("attached splitter")
    for bufferScanner.Scan() {
        logger.Printf("parsing buffer... %d", idx)
        msg := bufferScanner.Bytes()
        method, content, err :=  rpc.DecodeMessage(msg)

        if (err != nil) {
            logger.Printf("An error occurred %s", err)
        }

        // I previously erroneously used stderr as writing stream and hence wasn't able to receive further responses from the client
        writer := os.Stdout 

        handleMessage(logger, writer, state, method, content);
        logger.Printf("[progress] Buffer parsed for [%d]th time", idx)
        idx += 1
    }
    logger.Println("finished scanning")
}

// handle the message received from client
// UNIMPLEMENTED yet
func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, content []byte) {
    logger.Printf("Received message with method %s", method)

    switch method {
    case "initialize":
        var request lsp.InitializeRequest
        if err := json.Unmarshal(content, &request); err != nil {
            logger.Printf("received contents cannot be parsed : %s %s", content,  err)
        }
        logger.Printf("[INITIALIZE] Connected to client %s with version %s", request.InitializeRequestParams.ClientInfo.Name, request.InitializeRequestParams.ClientInfo.Version)

        response := lsp.NewInitializeResponse(request.ID)

        writeResponse(logger, response, writer)

    case "textDocument/didOpen":
        var request lsp.DidOpenTextDocumentNotification
        if err := json.Unmarshal(content, &request) ; err != nil {
            logger.Printf("[textdoc/didOpen]received contents cannot be parsed : %s %s", content,  err)
        }
        logger.Printf("[textdoc/didOpen] OPENED loaded file at: [ %s ]", request.DidOpenTextDocumentParams.TextDocumentItem.Uri)
        state.OpenDocument(request.DidOpenTextDocumentParams.TextDocumentItem.Uri, request.DidOpenTextDocumentParams.TextDocumentItem.Text)

    case "textDocument/didChange":
        var request lsp.DidChangeTextDocumentNotification
        if err := json.Unmarshal(content, &request) ; err != nil {
            logger.Printf("[textdoc/didChange] received contents cannot be parsed : %s %s", content,  err)
        }
        logger.Printf("[textdoc/didChange] UPDATED loaded file at: [ %s ]", request.Params.TextDocument.Uri)

        for _, change := range request.Params.Changes {
            state.UpdateDocument(request.Params.TextDocument.Uri, change.Text)
        }

    case "textDocument/hover":
        var request lsp.HoverRequest
        if err := json.Unmarshal(content, &request) ; err != nil {
            logger.Printf("[textdoc/hover] received contents cannot be parsed : %s %s", content,  err)
        }
        logger.Printf("[textdoc/hover] HOVER loaded file at: [ %s ]", request.Params.TextDocument.Uri)
        // prepare resposne
        response := state.Hover(request.ID, request.Params.TextDocument.Uri)       // write it back to the stream
        writeResponse(logger, response, writer)

    case "textDocument/definition":
        var request lsp.DefinitionRequest
        if err := json.Unmarshal(content, &request) ; err != nil {
            logger.Printf("[textdoc/defnition] received contents cannot be parsed : %s %s", content,  err)
        }
        logger.Printf("[textdoc/definition] DEFINITION loaded file at: [ %s ]", request.Params.TextDocument.Uri)
        // prepare resposne
        response := state.Definition(request.ID, request.Params.TextDocument.Uri, request.Params.Position)       // write it back to the stream
        writeResponse(logger, response, writer)

    case "textDocument/codeAction":
        var request lsp.DefinitionRequest
        if err := json.Unmarshal(content, &request) ; err != nil {
            logger.Printf("[textdoc/codeAction] received contents cannot be parsed : %s %s", content,  err)
        }
        logger.Printf("[textdoc/codeAction] CODE ACTION loaded file at: [ %s ]", request.Params.TextDocument.Uri)

        // prepare resposne
        response := state.CodeAction(request.ID, request.Params.TextDocument.Uri)       // write it back to the stream

        writeResponse(logger, response, writer)

    case "textDocument/completion":
        var request lsp.CompletionRequest
        if err := json.Unmarshal(content, &request) ; err != nil {
            logger.Printf("[textdoc/completion] received contents cannot be parsed : %s %s", content,  err)
        }
        logger.Printf("[textdoc/completion] COMPLETION loaded file at: [ %s ]", request.Params.TextDocument.Uri)

        // prepare resposne
        response := state.Completion(logger, request.ID, request.Params.TextDocument.Uri)       // write it back to the stream

        writeResponse(logger, response, writer)
    }
}

func writeResponse(logger *log.Logger, response any, writer io.Writer) {
        encodedResponse := rpc.EncodeMessage(response);
        writer.Write([]byte(encodedResponse))
        logger.Printf("responded back to client with %s", encodedResponse)
}

// create a file to write logs to and return it
func getLogger(path string) *log.Logger {
    logfile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

    if (err != nil) {
        panic("Invalid file path provided")
    }
    return log.New(logfile, "[custom-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
