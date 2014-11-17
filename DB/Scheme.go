package DB

import (
	"time"
)

type Answers struct {
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
	A18TimeUTC  time.Time `bson:"A18TimeUTC"`
	A19TimeUTC  time.Time `bson:"A19TimeUTC"`
	A20TimeUTC  time.Time `bson:"A20TimeUTC"`
	A21TimeUTC  time.Time `bson:"A21TimeUTC"`
	A22TimeUTC  time.Time `bson:"A22TimeUTC"`
	A23TimeUTC  time.Time `bson:"A23TimeUTC"`
	A24TimeUTC  time.Time `bson:"A24TimeUTC"`
	A25TimeUTC  time.Time `bson:"A25TimeUTC"`
	A26TimeUTC  time.Time `bson:"A26TimeUTC"` // Take this time on store A26

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
	A19Data string `bson:"A19Data"`
	A20Data string `bson:"A20Data"`
	A21Data string `bson:"A21Data"`
	A22Data string `bson:"A22Data"`
	A23Data string `bson:"A23Data"`
	A24Data string `bson:"A24Data"`
	A25Data string `bson:"A25Data"`
	A26Data string `bson:"A26Data"`
}
