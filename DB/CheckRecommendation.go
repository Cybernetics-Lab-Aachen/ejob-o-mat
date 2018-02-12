package DB

import (
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

// CheckRecommendation returns whether recommendation already exists in database.
func CheckRecommendation(session string) bool {

	dbSession, db := CustomerDB.DB()
	defer dbSession.Close()

	if db == nil { // Database not found
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return false
	}

	selector := bson.D{{"Session", session}}

	// Select number of recommendations
	n, errN := db.C(collRecommendations).Find(selector).Count()
	if errN != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameDATABASE, `Was not able to check the recommendations.`, errN.Error())
		return false
	}
	return n > 0 // whether there is a recommendation
}
