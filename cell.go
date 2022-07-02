package cell

type Cell struct {
	Mark string
}

func (c *Cell) Createcell(sign string) *Cell {
	var nec Cell
	nec.Mark = sign
	return &nec
}
