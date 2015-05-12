package Scheme

func (pg ProductGroup) GetProgressState(pos int) string {
	scaledPoints := float64(pg.Percent) / 5.0 // 5% := 1 full point (at the UI)

	if int(scaledPoints) >= pos {
		return ` progressitemdone`
	} else {
		return ``
	}
}
