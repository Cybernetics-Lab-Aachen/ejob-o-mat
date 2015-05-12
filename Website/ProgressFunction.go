package Website

func (pg PageQuestion) GetProgressState(pos int) string {
	scaledPoints := float64(pg.Progress) / 5.555 // 5.555% := 1 full point (at the UI)

	if int(scaledPoints) >= pos {
		return ` progressitemdone`
	} else {
		return ``
	}
}
