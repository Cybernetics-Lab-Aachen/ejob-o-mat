package DB

import (
	"time"

	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/ejob-o-mat/DB/Scheme"
	"gopkg.in/mgo.v2/bson"
)

// StoreNewAnswers creates a new answers record for this session in the db.
func StoreNewAnswers(session string) {

	// Get the database:
	dbSession, db := CustomerDB.DB()
	defer dbSession.Close()

	if db == nil { // Database not found
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	// Insert answers
	db.C(collAnswers).Insert(Scheme.Survey{
		SchemeVersion: Scheme.CurrentVersion,
		Session:       session,
		CreateTimeUTC: time.Now().UTC(),
	})
}

// UpdateAnswers updates an existing answers record.
func UpdateAnswers(survey Scheme.Survey) {

	// Get the database:
	dbSession, db := CustomerDB.DB()
	defer dbSession.Close()

	if db == nil { // Database not found
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	selector := bson.D{{Name: "Session", Value: survey.Session}}

	// Update answers
	db.C(collAnswers).Update(selector, survey)
}
