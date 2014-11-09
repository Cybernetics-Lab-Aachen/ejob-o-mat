package Website

import (
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
)

func init() {
	betaPassword = ConfigurationDB.Read(`BetaPassword`)
}
