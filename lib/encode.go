package qrc

import (
    qrcode "github.com/skip2/go-qrcode"
)

// EncodeToGrid creates a Grid from the input text using QR error correction level Low.
func EncodeToGrid(text string) (Grid, error) {
    // Create QR code with default size; we'll use the bitmap directly.
    qr, err := qrcode.New(text, qrcode.Low)
    if err != nil {
        return nil, err
    }
    bm := qr.Bitmap()
    h := len(bm)
    w := 0
    if h > 0 {
        w = len(bm[0])
    }
    data := make([][]bool, h)
    for y := 0; y < h; y++ {
        data[y] = make([]bool, w)
        for x := 0; x < w; x++ {
            data[y][x] = bm[y][x]
        }
    }
    return &BoolGrid{W: w, H: h, Data: data}, nil
}
