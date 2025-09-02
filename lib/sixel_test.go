package qrc

import (
    "bytes"
    "strings"
    "testing"
)

func TestPrintSixel(t *testing.T) {
    grid, err := EncodeToGrid("test")
    if err != nil {
        t.Fatalf("encode failed: %v", err)
    }
    var buf bytes.Buffer
    PrintSixel(&buf, grid, false)
    out := buf.String()
    if !strings.HasPrefix(out, "\x1BPq") {
        t.Fatalf("sixel output should start with DCS introducer, got: %q", out[:min(10, len(out))])
    }
    if !strings.HasSuffix(out, "\x1B\\") { // ST
        t.Fatalf("sixel output should end with ST, got tail: %q", out[max(0, len(out)-10):])
    }
}

func min(a, b int) int { if a < b { return a }; return b }
func max(a, b int) int { if a > b { return a }; return b }
