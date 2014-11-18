package Algorithm

func kindConditionalPossibility(answer, productValue string) (diff int) {

	switch ans := answer; ans {
	case `0`:
		if productValue == `Possible` {
			diff = -1
		} else if productValue == `NotPossible` {
			diff = 1
		} else if productValue == `Decision` {
			diff = 1
		}
	case `1`:
		if productValue == `Possible` {
			diff = 1
		} else if productValue == `NotPossible` {
			diff = -1
		} else if productValue == `Decision` {
			diff = 1
		}
	case `*`:
		diff = 1
	}

	return
}
