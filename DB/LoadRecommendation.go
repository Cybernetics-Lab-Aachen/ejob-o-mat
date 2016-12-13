package DB

import (
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"gopkg.in/mgo.v2/bson"
)

//LoadRecommendation returns the recommendation for this session from db.
func LoadRecommendation(session string) (result Scheme.Recommendation) {

	dbSession, db := CustomerDB.DB()
	defer dbSession.Close()

	if db == nil { // Database not found
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	selector := bson.D{{"Session", session}}

	// Select result from db
	if err := db.C(collRecommendations).Find(selector).One(&result); err != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameDATABASE, `Was not able to load the recommendation.`, err.Error())
		return
	}

	return //Everything went ok
}
