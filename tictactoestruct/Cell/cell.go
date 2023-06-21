package cell

type Cell struct {
	mark string
}

func NewCell() *Cell {
	return &Cell{
		mark: "Empty",
	}
}

func (c *Cell) IsCellMarked() bool {
	return c.mark != "Empty"
}

func (c *Cell) MarkCell(mark string) {
	c.mark = mark
}

func (c *Cell) MarkedCell() string {
	return c.mark
}
