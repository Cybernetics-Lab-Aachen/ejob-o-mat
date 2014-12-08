package DB

import (
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

func CheckRecommendation(session string) (existing bool) {

	dbSession, db := CustomerDB.DB()
	defer dbSession.Close()

	if db == nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	selector := bson.D{{"Session", session}}
	ocollRecommendations := db.C(collRecommendations)

	if n, errN := ocollRecommendations.Find(selector).Count(); errN != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameDATABASE, `Was not able to check the recommendations.`, errN.Error())
		existing = false
	} else {
		if n > 0 {
			existing = true
		} else {
			existing = false
		}
	}

	return
}
