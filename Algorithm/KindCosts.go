package Algorithm

func kindCosts(answer, productValue, productValueFree string) (diff int) {

	switch ans := answer; ans {
	case `none`:
		if productValue == `None` {
			diff = 1
		} else if productValue == `Up to USD 5 (non-recurring)` {
			diff = -1
		} else if productValue == `Up to USD 10 (non-recurring)` {
			diff = -1
		} else if productValue == `Up to USD 25 (non-recurring)` {
			diff = -1
		} else if productValue == `Up to USD 50 (non-recurring)` {
			diff = -1
		} else if productValue == `Up to USD 100 (non-recurring)` {
			diff = -2
		} else if productValue == `Up to USD 300 (non-recurring)` {
			diff = -2
		} else if productValue == `Up to USD 600 (non-recurring)` {
			diff = -2
		} else if productValue == `Up to USD 1,000 (non-recurring)` {
			diff = -4
		} else if productValue == `Up to USD 5,000 (non-recurring)` {
			diff = -5
		} else if productValue == `More than USD 5,000 (non-recurring)` {
			diff = -6
		}
	case `up10`:
		if productValue == `None` {
			diff = 1
		} else if productValue == `Up to USD 5 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 10 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 25 (non-recurring)` {
			diff = 0
		} else if productValue == `Up to USD 50 (non-recurring)` {
			diff = -1
		} else if productValue == `Up to USD 100 (non-recurring)` {
			diff = -1
		} else if productValue == `Up to USD 300 (non-recurring)` {
			diff = -1
		} else if productValue == `Up to USD 600 (non-recurring)` {
			diff = -2
		} else if productValue == `Up to USD 1,000 (non-recurring)` {
			diff = -3
		} else if productValue == `Up to USD 5,000 (non-recurring)` {
			diff = -4
		} else if productValue == `More than USD 5,000 (non-recurring)` {
			diff = -5
		}
	case `up25`:
		if productValue == `None` {
			diff = 1
		} else if productValue == `Up to USD 5 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 10 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 25 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 50 (non-recurring)` {
			diff = -1
		} else if productValue == `Up to USD 100 (non-recurring)` {
			diff = -1
		} else if productValue == `Up to USD 300 (non-recurring)` {
			diff = -2
		} else if productValue == `Up to USD 600 (non-recurring)` {
			diff = -2
		} else if productValue == `Up to USD 1,000 (non-recurring)` {
			diff = -4
		} else if productValue == `Up to USD 5,000 (non-recurring)` {
			diff = -5
		} else if productValue == `More than USD 5,000 (non-recurring)` {
			diff = -6
		}
	case `up100`:
		if productValue == `None` {
			diff = 1
		} else if productValue == `Up to USD 5 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 10 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 25 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 50 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 100 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 300 (non-recurring)` {
			diff = -1
		} else if productValue == `Up to USD 600 (non-recurring)` {
			diff = -2
		} else if productValue == `Up to USD 1,000 (non-recurring)` {
			diff = -3
		} else if productValue == `Up to USD 5,000 (non-recurring)` {
			diff = -4
		} else if productValue == `More than USD 5,000 (non-recurring)` {
			diff = -5
		}
	case `up1000`:
		if productValue == `None` {
			diff = 1
		} else if productValue == `Up to USD 5 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 10 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 25 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 50 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 100 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 300 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 600 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 1,000 (non-recurring)` {
			diff = 1
		} else if productValue == `Up to USD 5,000 (non-recurring)` {
			diff = -3
		} else if productValue == `More than USD 5,000 (non-recurring)` {
			diff = -4
		}
	}

	switch ans := answer; ans {
	case `none`:
		if productValueFree == `Yes` {
			diff = 1
		} else if productValueFree == `No` {
			// No change
		} else if productValueFree == `Perhaps` {
			diff = 1
		} else if productValueFree == `Decision` {
			diff = 1
		}
	}

	return
}
