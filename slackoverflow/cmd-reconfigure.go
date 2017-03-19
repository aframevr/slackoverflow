package slackoverflow

import "github.com/aframevr/slackoverflow/std"

// slackoverflow reconfigure
// Reconfigure SlackOverflow.
type cmdReconfigure struct{}

func (a *cmdReconfigure) Execute(args []string) error {

	doReconfigure := std.AskForConfirmation("Start interactive confguration?")
	if doReconfigure {
		slackoverflow.config.ReconfigureAll()
		Ok("Reconfiguration done")
	} else {
		Warning("Reconfiguration canceled!")
	}

	return nil
}
