package Website

func (pg PageQuestion) GetProgressState(pos int) string {
	if pg.Progress >= pos {
		return ` progressitemdone`
	} else {
		return ``
	}
}
