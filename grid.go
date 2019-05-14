package glife

type Grid struct {
	rows, cols int
	cells      []bool
}

func NewGrid(rows, cols int) *Grid {
	return &Grid{
		rows:  rows,
		cols:  cols,
		cells: make([]bool, rows*cols),
	}
}

func (g *Grid) IsAlive(row, col int) bool {
	return g.cells[row*g.cols+col]
}

func (g *Grid) MarkAlive(row, col int) {
	g.cells[row*g.cols+col] = true
}
