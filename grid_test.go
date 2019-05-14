package glife

import "testing"

func TestThatNewGridHasAllDeadCells(t *testing.T) {
	g := NewGrid(1, 1)

	if g.IsAlive(1, 1) {
		t.Errorf("expected cell at (1,1) to be dead")
	}
}
