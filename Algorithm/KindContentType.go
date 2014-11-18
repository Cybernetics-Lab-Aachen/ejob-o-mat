package Algorithm

func kindContentType(answer, productValue string) (diff int) {

	switch ans := answer; ans {
	case `static`:
		if productValue == `Static Content` {
			diff = 1
		} else if productValue == `Changeable Content` {
			diff = -1
		} else if productValue == `Interactive Content` {
			diff = -2
		} else if productValue == `Changeable and/or Interactive Content` {
			diff = -2
		}
	case `change`:
		if productValue == `Static Content` {
			diff = -2
		} else if productValue == `Changeable Content` {
			diff = 1
		} else if productValue == `Interactive Content` {
			diff = 0
		} else if productValue == `Changeable and/or Interactive Content` {
			diff = 1
		}
	case `interact`:
		if productValue == `Static Content` {
			diff = -2
		} else if productValue == `Changeable Content` {
			diff = -1
		} else if productValue == `Interactive Content` {
			diff = 1
		} else if productValue == `Changeable and/or Interactive Content` {
			diff = 1
		}
	case `*`:
		diff = 1
	}

	return
}
