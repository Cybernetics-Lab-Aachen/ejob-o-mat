package Scheme

import (
	"time"
)

var (
	// CurrentVersion increases every time the database layout changes
	CurrentVersion byte = 3
)

// ProductGroupsByPoints implements sort.Interface for []ProductGroup based on
// the Points field.
type ProductGroupsByPoints []ProductGroup

func (groups ProductGroupsByPoints) Len() int {
	return len(groups)
}

func (groups ProductGroupsByPoints) Less(i, j int) bool {
	return groups[i].Points > groups[j].Points
}

func (groups ProductGroupsByPoints) Swap(i, j int) {
	groups[i], groups[j] = groups[j], groups[i]
}

//ProductGroup Stores the display data (points, recommendation percentage, etc.) for a product group.
type ProductGroup struct {
	Name             string          `bson:"Name"`
	Points           int             `bson:"Points"`
	Percent          int             `bson:"Percent"`
	XMLIndex         int             `bson:"XMLIndex"`
	AnswerInfluences map[string]int8 `bson:"AnswerInfluences"`
}

//Recommendation stores the product group ranking calculated by the algorithm for the given answers.
type Recommendation struct {
	SchemeVersion byte           `bson:"Version"`
	CreateTimeUTC time.Time      `bson:"CreateTimeUTC"`
	Session       string         `bson:"Session"`
	ProductGroups []ProductGroup `bson:"ProductGroups"`
}

//Survey stores the answer, time and weight for each question.
type Survey struct {
	SchemeVersion byte              `bson:"Version"`
	CreateTimeUTC time.Time         `bson:"CreateTimeUTC"`
	Session       string            `bson:"Session"`
	StartTimeQ1   time.Time         `bson:"StartTimeQ1"` // Take this time on deliver Q1
	Answers       map[string]Answer `bson:"Answers"`
	Questions     []string          `bson:"Questions"` // Permutation of the questions
}

//GetByInternalName returns the given answer to a question by its internal question name.
func (survey Survey) GetByInternalName(name string) string {
	return survey.Answers[name].Data
}

//Answer contains the answer, time and weight for a question.
type Answer struct {
	TimeUTC time.Time `bson:"TimeUTC"`
	Data    string    `bson:"Data"`
	Weight  byte      `bson:"Weight"`
}

// Feedback contains the data a feedback.
type Feedback struct {
	Text           string `bson:"Text"`
	Rating         byte   `bson:"Rating"`
	Session        string `bson:"Session"`
	SourceLocation string `bson:"SourceLocation"`
}
