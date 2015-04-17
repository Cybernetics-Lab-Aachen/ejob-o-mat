package Algorithm

func kindConditionalPresence(answer, productValue string) (diff int) {

	switch ans := answer; ans {
	case `0`:
		if productValue == `NotPresent` {
			diff = 1
		} else if productValue == `PresentAndNotConfigurable` {
			diff = -1
		} else if productValue == `Configurable` {
			diff = 1
		}
	case `1`:
		if productValue == `NotPresent` {
			diff = -1
		} else if productValue == `PresentAndNotConfigurable` {
			diff = 1
		} else if productValue == `Configurable` {
			diff = 1
		}
	case `*`:
		diff = 0
	}

	return
}
