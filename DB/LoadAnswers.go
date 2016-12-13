package DB

import (
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"gopkg.in/mgo.v2/bson"
)

//LoadAnswers returns the answers for this session from db and whether an error occurred.
func LoadAnswers(session string) (Scheme.Answers, bool) {
	answers := Scheme.Answers{}

	// Get the database:
	dbSession, db := CustomerDB.DB()
	defer dbSession.Close()

	if db == nil { // Database not found
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return answers, true
	}

	selector := bson.D{{"Session", session}}

	// Select answers from db
	if err := db.C(collAnswers).Find(selector).One(&answers); err != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameDATABASE, `Was not able to find this session in the database.`, session, err.Error())
		return answers, true
	}

	return answers, false // Everything went ok
}
