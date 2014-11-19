package DB

import (
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"gopkg.in/mgo.v2/bson"
)

func LoadAnswers(session string) (result Scheme.Answers) {

	// Get the database:
	dbSession, db := CustomerDB.DB()
	defer dbSession.Close()

	if db == nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	answers := Scheme.Answers{}
	selector := bson.D{{"Session", session}}
	ocollAnswers := db.C(collAnswers)
	if err := ocollAnswers.Find(selector).One(&answers); err != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameDATABASE, `Was not able to find this session in the database.`, session, err.Error())
		return
	}

	result = answers
	return
}
