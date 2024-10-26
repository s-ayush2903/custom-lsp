package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
    content, err := json.Marshal(msg);
    if (err != nil) {
        panic(err)
    }
    return fmt.Sprintf("Content-Length %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
    Message string `json:"Method"`
}

// currently returns the content length along with error (if applicable)
func DecodeMessage(msg []byte) (string, []byte, error) {
    header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})

    if (found == false) {
        return "", nil, errors.New("Separators not found!");
    }

    _ = content
    contentLengthBytes := header[len("Content-Length "):]
    contentLength, err := strconv.Atoi(string(contentLengthBytes));

    if (err != nil) {
        return "", nil, err;
    }

    var message BaseMessage;
    derr := json.Unmarshal(content[:contentLength], &message);
    if (derr != nil) {
        fmt.Printf("error for %s", content)
        panic(derr)
    }

    return message.Message, content, nil;
}
