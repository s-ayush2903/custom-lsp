package main

import (
	"bufio"
	"encoding/json"
	"log"
	"lsp-go/lsp"
	"lsp-go/rpc"
	"os"
)

func main() {
    logger := getLogger("/Users/ayush/Desktop/learn/lsp-go/logs.txt")
    logger.Printf("logger initiated bruv..")

    logger.Printf("hiii");

    bufferScanner := bufio.NewScanner(os.Stdin);

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

        handleMessage(logger, method, content);
        logger.Printf("[progress] Buffer parsed for [%d]th time", idx)
        idx += 1
    }
    logger.Println("finished scanning")
}

// handle the message received from client
// UNIMPLEMENTED yet
func handleMessage(logger *log.Logger, method string, content []byte) {
    logger.Printf("Received message with method %s", method)

    switch method {
    case "initialize":
        var request lsp.InitializeRequest
        if err := json.Unmarshal(content, &request); err != nil {
            logger.Printf("received contents cannot be parsed : %s %s", content,  err)
        }
        logger.Printf("Connected to client %s with version %s", request.InitializeRequestParams.ClientInfo.Name, request.InitializeRequestParams.ClientInfo.Version)

        response := lsp.NewInitializeResponse(request.ID)
        encodedResponse := rpc.EncodeMessage(response);

        // I previously erroneously used stderr as writing stream and hence wasn't able to receive further responses from the client
        writer := os.Stdout 
        writer.Write([]byte(encodedResponse))
        logger.Printf("responded back to client with %s", encodedResponse)
    case "textDocument/didOpen":
        var request lsp.DidOpenTextDocumentNotification
        if err := json.Unmarshal(content, &request) ; err != nil {
            logger.Printf("received contents cannot be parsed : %s %s", content,  err)
        }
        logger.Printf("client loaded file at: [ %s ] with contents: { %s }", request.DidOpenTextDocumentParams.TextDocumentItem.Uri, request.DidOpenTextDocumentParams.TextDocumentItem.Text)
    }
}

// create a file to write logs to and return it
func getLogger(path string) *log.Logger {
    logfile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

    if (err != nil) {
        panic("Invalid file path provided")
    }
    return log.New(logfile, "[custom-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
