package qrc

// Grid represents a read-only 2D boolean bitmap.
type Grid interface {
    Width() int
    Height() int
    Get(x, y int) bool
}

// BoolGrid is a simple in-memory implementation of Grid.
type BoolGrid struct {
    W, H int
    Data [][]bool
}

func (g *BoolGrid) Width() int  { return g.W }
func (g *BoolGrid) Height() int { return g.H }
func (g *BoolGrid) Get(x, y int) bool { return g.Data[y][x] }
