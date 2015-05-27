package DB

import (
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
)

func StoreFeedback(result Scheme.Feedback) {

	dbSession, db := CustomerDB.DB()
	defer dbSession.Close()

	if db == nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	ocollFeedback := db.C(collFeedback)
	ocollFeedback.Insert(result)
}
