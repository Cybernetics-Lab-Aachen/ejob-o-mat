package Algorithm

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"strconv"
)

// kindAmountAccesses calculates the influence of an answer given to a amount of accesses question.
func kindAmountAccesses(answer, minString, maxString string) (diff int8) {

	min := 1
	max := 4
	ans := 0

	//Parse strings

	if value, errValue := strconv.Atoi(minString); errValue != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityLow, LM.ImpactLow, LM.MessageNamePARSE, `Was not able to parse the min value for this product.`, fmt.Sprintf("Wrong value='%s'", minString))
		diff = 0
		return
	} else {
		min = value
	}

	if value, errValue := strconv.Atoi(maxString); errValue != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityLow, LM.ImpactLow, LM.MessageNamePARSE, `Was not able to parse the max value for this product.`, fmt.Sprintf("Wrong value='%s'", maxString))
		diff = 0
		return
	} else {
		max = value
	}

	if value, errValue := strconv.Atoi(answer[12:]); errValue != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityLow, LM.ImpactLow, LM.MessageNamePARSE, `Was not able to parse the answer value for this product.`, fmt.Sprintf("Wrong value='%s'", answer))
		diff = 0
		return
	} else {
		ans = value
	}

	// Check if answer lies inside supported interval
	if ans >= min && ans <= max {
		diff = 1
	} else {
		diff = -1
	}

	return
}
