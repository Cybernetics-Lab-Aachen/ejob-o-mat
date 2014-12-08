package DB

import (
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"gopkg.in/mgo.v2/bson"
)

func LoadRecommendation(session string) (result Scheme.Recommendation) {

	dbSession, db := CustomerDB.DB()
	defer dbSession.Close()

	if db == nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	data := Scheme.Recommendation{}
	selector := bson.D{{"Session", session}}
	ocollRecommendations := db.C(collRecommendations)

	if err := ocollRecommendations.Find(selector).One(&data); err != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameDATABASE, `Was not able to load the recommendation.`, err.Error())
	} else {
		result = data
	}

	return
}
