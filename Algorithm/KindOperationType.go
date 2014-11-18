package Algorithm

func kindOperationType(answer, productValue string) (diff int) {

	switch ans := answer; ans {
	case `support4lecture`:
		if productValue == `ReplacePresenceLecture` {
			diff = -1
		} else if productValue == `Support4Lecture` {
			diff = 1
		} else if productValue == `Both` {
			diff = 1
		}
	case `replace`:
		if productValue == `ReplacePresenceLecture` {
			diff = 1
		} else if productValue == `Support4Lecture` {
			diff = -1
		} else if productValue == `Both` {
			diff = 1
		}
	}

	return
}
