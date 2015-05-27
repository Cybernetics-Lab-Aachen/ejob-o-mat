package Scheme

import (
	"time"
)

var (
	CURRENT_VERSION byte = 1
)

func (groups ProductGroups) Len() int {
	return len(groups)
}

func (groups ProductGroups) Less(i, j int) bool {
	return groups[i].Points > groups[j].Points
}

func (groups ProductGroups) Swap(i, j int) {
	groups[i], groups[j] = groups[j], groups[i]
}

type ProductGroups []ProductGroup

type ProductGroup struct {
	Name     string `bson:"Name"`
	Points   int    `bson:"Points"`
	Percent  int    `bson:"Percent"`
	XMLIndex int    `bson:"XMLIndex"`
}

type Recommendation struct {
	SchemeVersion byte           `bson:"Version"`
	CreateTimeUTC time.Time      `bson:"CreateTimeUTC"`
	Session       string         `bson:"Session"`
	ProductGroups []ProductGroup `bson:"ProductGroups"`
}

type Answers struct {
	SchemeVersion byte      `bson:"Version"`
	CreateTimeUTC time.Time `bson:"CreateTimeUTC"`
	Session       string    `bson:"Session"`

	StartTimeQ1 time.Time `bson:"StartTimeQ1"` // Take this time on deliver Q1
	A1TimeUTC   time.Time `bson:"A1TimeUTC"`   // Take this time on store A1
	A2TimeUTC   time.Time `bson:"A2TimeUTC"`
	A3TimeUTC   time.Time `bson:"A3TimeUTC"`
	A4TimeUTC   time.Time `bson:"A4TimeUTC"`
	A5TimeUTC   time.Time `bson:"A5TimeUTC"`
	A6TimeUTC   time.Time `bson:"A6TimeUTC"`
	A7TimeUTC   time.Time `bson:"A7TimeUTC"`
	A8TimeUTC   time.Time `bson:"A8TimeUTC"`
	A9TimeUTC   time.Time `bson:"A9TimeUTC"`
	A10TimeUTC  time.Time `bson:"A10TimeUTC"`
	A11TimeUTC  time.Time `bson:"A11TimeUTC"`
	A12TimeUTC  time.Time `bson:"A12TimeUTC"`
	A13TimeUTC  time.Time `bson:"A13TimeUTC"`
	A14TimeUTC  time.Time `bson:"A14TimeUTC"`
	A15TimeUTC  time.Time `bson:"A15TimeUTC"`
	A16TimeUTC  time.Time `bson:"A16TimeUTC"`
	A17TimeUTC  time.Time `bson:"A17TimeUTC"`
	A18TimeUTC  time.Time `bson:"A18TimeUTC"` // Take this time on store A25

	A1Data  string `bson:"A1Data"`
	A2Data  string `bson:"A2Data"`
	A3Data  string `bson:"A3Data"`
	A4Data  string `bson:"A4Data"`
	A5Data  string `bson:"A5Data"`
	A6Data  string `bson:"A6Data"`
	A7Data  string `bson:"A7Data"`
	A8Data  string `bson:"A8Data"`
	A9Data  string `bson:"A9Data"`
	A10Data string `bson:"A10Data"`
	A11Data string `bson:"A11Data"`
	A12Data string `bson:"A12Data"`
	A13Data string `bson:"A13Data"`
	A14Data string `bson:"A14Data"`
	A15Data string `bson:"A15Data"`
	A16Data string `bson:"A16Data"`
	A17Data string `bson:"A17Data"`
	A18Data string `bson:"A18Data"`

	A1Weight  byte `bson:"A1Weight"`
	A2Weight  byte `bson:"A2Weight"`
	A3Weight  byte `bson:"A3Weight"`
	A4Weight  byte `bson:"A4Weight"`
	A5Weight  byte `bson:"A5Weight"`
	A6Weight  byte `bson:"A6Weight"`
	A7Weight  byte `bson:"A7Weight"`
	A8Weight  byte `bson:"A8Weight"`
	A9Weight  byte `bson:"A9Weight"`
	A10Weight byte `bson:"A10Weight"`
	A11Weight byte `bson:"A11Weight"`
	A12Weight byte `bson:"A12Weight"`
	A13Weight byte `bson:"A13Weight"`
	A14Weight byte `bson:"A14Weight"`
	A15Weight byte `bson:"A15Weight"`
	A16Weight byte `bson:"A16Weight"`
	A17Weight byte `bson:"A17Weight"`
	A18Weight byte `bson:"A18Weight"`
}

type Feedback struct {
	Text           string `bson:"Text"`
	Rating         byte   `bson:"Rating"`
	Session        string `bson:"Session"`
	SourceLocation string `bson:"SourceLocation"`
}
