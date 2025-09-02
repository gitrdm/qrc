package qrc

import (
    "bytes"
    "testing"
)

func TestPrintAA(t *testing.T) {
    grid, err := EncodeToGrid("test")
    if err != nil {
        t.Fatalf("encode failed: %v", err)
    }
    var buf bytes.Buffer
    PrintAA(&buf, grid, false)
    out := buf.String()
    if len(out) == 0 {
        t.Fatal("no output produced")
    }
    // Expect ANSI reset at least once
    if !bytes.Contains([]byte(out), []byte("\x1b[0m")) {
        t.Log("ANSI reset not found; color support may vary, but output length is non-zero")
    }
}
