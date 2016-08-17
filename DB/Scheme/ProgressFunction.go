package Scheme

func (pg ProductGroup) GetProgressState(influence int) string {
	if influence > 0 {
		return ` progressitemdone`
	} else if influence < 0 {
		return ` progressitemundone`
	} else {
		return ``
	}
}
