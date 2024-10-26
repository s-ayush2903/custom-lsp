package rpc_test

import (
	"lsp-go/rpc"
	"testing"
)

type ArbitStruct struct {
    ArbitVariable bool
    AnotherVariable string
    ThirdVariable int
}

func TestEncoderTest(t *testing.T) {
    expected := "Content-Length 67\r\n\r\n{\"ArbitVariable\":true,\"AnotherVariable\":\"hello\",\"ThirdVariable\":12}";
    observed := rpc.EncodeMessage(ArbitStruct{true, "hello", 12});

    if (expected != observed) {
        t.Fatalf("OBSERVED %s while EXPECTED %s", observed, expected)
    }
}

func TestDecode(t *testing.T) {
    msg := "Content-Length 15\r\n\r\n{\"Method\":\"hi\"}"
    message, content, err := rpc.DecodeMessage([]byte(msg));
    observed := len(content)
    expected := 15

    if (err != nil) {
        t.Fatal(err)
    }

    if (observed != expected) {
        t.Fatalf("OBSERVED %d while EXPECTED %d %s", err, expected, message)
    }
}
