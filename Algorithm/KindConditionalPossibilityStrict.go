package Algorithm

func kindConditionalPossibilityStrict(answer, productValue string) (diff int) {

	switch ans := answer; ans {
	case `0`:
		if productValue == `Possible` {
			diff = 1
		} else if productValue == `NotPossible` {
			diff = 1
		}
	case `1`:
		if productValue == `Possible` {
			diff = 1
		} else if productValue == `NotPossible` {
			diff = -2
		}
	case `*`:
		diff = 0
	}

	return
}
