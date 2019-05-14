package glife

import "testing"

func expectDead(g *Grid, row, col int, t *testing.T) {
	if g.IsAlive(row, col) {
		t.Errorf("expected cell to be dead")
	}
}

func expectAlive(g *Grid, row, col int, t *testing.T) {
	if !g.IsAlive(row, col) {
		t.Errorf("expected cell to be alive")
	}
}

func TestThatNewGridHasAllDeadCells(t *testing.T) {
	g := NewGrid(1, 1)

	expectDead(g, 0, 0, t)
}

func TestThatACellCanBeMarkedAlive(t *testing.T) {
	g := NewGrid(1, 1)

	g.MarkAlive(0, 0)

	expectAlive(g, 0, 0, t)
}

func TestThatCellsCanHaveDifferentStates(t *testing.T) {
	g := NewGrid(1, 2)

	g.MarkAlive(0, 0)

	expectAlive(g, 0, 0, t)
	expectDead(g, 0, 1, t)
}

func TestGridWithSecondRow(t *testing.T) {
	g := NewGrid(2, 2)

	g.MarkAlive(1, 0)

	expectDead(g, 0, 0, t)
	expectDead(g, 0, 1, t)
	expectAlive(g, 1, 0, t)
	expectDead(g, 1, 1, t)
}
