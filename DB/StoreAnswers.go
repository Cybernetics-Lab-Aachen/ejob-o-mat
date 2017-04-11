package DB

import (
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"gopkg.in/mgo.v2/bson"
)

// StoreNewAnswers creates a new answers record for this session in the db.
func StoreNewAnswers(ans Scheme.Answers) {

	// Get the database:
	dbSession, db := CustomerDB.DB()
	defer dbSession.Close()

	if db == nil { // Database not found
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	// Insert answers
	db.C(collAnswers).Insert(ans)
}

// UpdateAnswers updates an existing answers record.
func UpdateAnswers(ans Scheme.Answers) {

	// Get the database:
	dbSession, db := CustomerDB.DB()
	defer dbSession.Close()

	if db == nil { // Database not found
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	selector := bson.D{{"Session", ans.Session}}

	// Update answers
	db.C(collAnswers).Update(selector, ans)
}
