package Algorithm

func kindConditionalYesNo(answer, productValue string) (diff int) {

	switch ans := answer; ans {
	case `0`:
		if productValue == `Yes` {
			diff = -1
		} else if productValue == `No` {
			diff = 1
		} else if productValue == `Perhaps` {
			diff = 1
		} else if productValue == `Decision` {
			diff = 1
		}
	case `1`:
		if productValue == `Yes` {
			diff = 1
		} else if productValue == `No` {
			diff = -1
		} else if productValue == `Perhaps` {
			diff = 1
		} else if productValue == `Decision` {
			diff = 1
		}
	case `*`:
		diff = 1
	}

	return
}
