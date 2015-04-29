package Algorithm

func kindCommon(answer, productValue string) (diff int) {

	switch ans := answer; ans {
	case `0`:
		if productValue == `Possible` {
			diff = 1
		} else if productValue == `Impossible` {
			diff = 1
		} else if productValue == `Mandatory` {
			diff = -1
		}
	case `1`:
		if productValue == `Possible` {
			diff = 1
		} else if productValue == `Impossible` {
			diff = -1
		} else if productValue == `Mandatory` {
			diff = 1
		}
	case `*`:
		diff = 0
	}

	return
}
