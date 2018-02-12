package XML

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

var (
	senderName LM.Sender = `XML`
	data                 = Data{} // Representation of Data.xml in memory
	// Questions is a name -> QuestionGroup mapping for all questions present in the XML
	Questions map[string]QuestionGroup
)
