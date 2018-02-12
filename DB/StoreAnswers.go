package DB

import (
	"math/rand"
	"time"

	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/ejob-o-mat/DB/Scheme"
	"github.com/SommerEngineering/ejob-o-mat/XML"
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

	// Create slice of all question
	questions := make([]string, 0, len(XML.Questions))
	for questionName := range XML.Questions {
		questions = append(questions, questionName)
	}
	for i := range questions { // Shuffle questions (Fisher–Yates shuffle)
		j := rand.Intn(i + 1)
		questions[i], questions[j] = questions[j], questions[i]
	}

	// Insert answers
	db.C(collAnswers).Insert(Scheme.Survey{
		SchemeVersion: Scheme.CurrentVersion,
		Session:       session,
		CreateTimeUTC: time.Now().UTC(),
		Questions:     questions,
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
