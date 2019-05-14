package glife

type Grid struct {
}

func NewGrid(rows, cols int) *Grid {
	return &Grid{}
}

func (g *Grid) IsAlive(row, col int) bool {
	return false
}
