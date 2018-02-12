package main

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/System"
)

func main() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `ejob-o-mat is starting.`)
	System.InitHandlers()
	registerHandlers()
	registerAllAppConfigurationParameters()
	System.StartAndBlockForever()
}
