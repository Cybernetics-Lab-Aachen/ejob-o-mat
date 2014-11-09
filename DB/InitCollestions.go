package DB

import (
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

func initCollections() {

	// Get the database:
	dbSession, db := CustomerDB.DB()
	defer dbSession.Close()

	if db == nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	// Get all collections:
	ocollAnswers := db.C(collAnswers)

	// Ensure the indexes:
	ocollAnswers.EnsureIndexKey(`Session`)
	ocollAnswers.EnsureIndexKey(`TimeUTC`)
}
