package Algorithm

func kindAppropriateCountStudents(answer, productValue string) (diff int) {

	switch ans := answer; ans {
	case `small`:
		if productValue == `OneStudent` {
			diff = 1
		} else if productValue == `SmallSizeGroups` {
			diff = 1
		} else if productValue == `MidSizeGroups` {
			diff = 1
		} else if productValue == `HugeAmount` {
			diff = -1
		} else if productValue == `Masses` {
			diff = -1
		}
	case `mid`:
		if productValue == `OneStudent` {
			diff = 1
		} else if productValue == `SmallSizeGroups` {
			diff = 1
		} else if productValue == `MidSizeGroups` {
			diff = 1
		} else if productValue == `HugeAmount` {
			diff = 1
		} else if productValue == `Masses` {
			diff = -1
		}
	case `huge`:
		if productValue == `OneStudent` {
			diff = -1
		} else if productValue == `SmallSizeGroups` {
			diff = -1
		} else if productValue == `MidSizeGroups` {
			diff = -1
		} else if productValue == `HugeAmount` {
			diff = 1
		} else if productValue == `Masses` {
			diff = 1
		}
	case `masses`:
		if productValue == `OneStudent` {
			diff = -1
		} else if productValue == `SmallSizeGroups` {
			diff = -1
		} else if productValue == `MidSizeGroups` {
			diff = -1
		} else if productValue == `HugeAmount` {
			diff = -1
		} else if productValue == `Masses` {
			diff = 1
		}
	}

	return
}
