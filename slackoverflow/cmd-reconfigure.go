package slackoverflow

import "github.com/aframevr/slackoverflow/std"

// slackoverflow reconfigure
// Reconfigure SlackOverflow.
type cmdReconfigure struct{}

func (a *cmdReconfigure) Execute(args []string) error {

	doReconfigure := std.AskForConfirmation("This will overwrite current configuration, Are you sure you want to continue?")
	if doReconfigure {
		stackoverflow.config.Reconfigure()
		Ok("Reconfiguration done")
	} else {
		Warning("Reconfiguration canceled!")
	}

	return nil
}
