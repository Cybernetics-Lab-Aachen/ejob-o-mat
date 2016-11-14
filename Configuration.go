package main

import "github.com/SommerEngineering/Ocean/ConfigurationDB"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func registerAllAppConfigurationParameters() {

	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all app configuration parameters.`)
	defer Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all app configuration parameters done.`)

	ConfigurationDB.CheckSingleConfigurationPresentsAndAddIfMissing(`BetaPassword`, `ELLI`)
	ConfigurationDB.CheckSingleConfigurationPresentsAndAddIfMissing(`SiteVerificationToken`, `ELLI`)
}
