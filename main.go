package main

import (
	"bufio"
	"log"
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
    logger.Println("attached splitter")
    for bufferScanner.Scan() {
        logger.Println("parsing buffer...")
        msg := bufferScanner.Text()
        handleMessage(logger, msg);
    }
    logger.Println("finished scanning")
}

// handle the message received from client
// UNIMPLEMENTED yet
func handleMessage(logger *log.Logger, message any) {
    logger.Println(message)
}

// create a file to write logs to and return it
func getLogger(path string) *log.Logger {
    logfile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

    if (err != nil) {
        panic("Invalid file path provided")
    }
    return log.New(logfile, "[custom-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
