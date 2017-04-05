package slackoverflow

import (
	"os"
	"os/signal"

	"github.com/robfig/cron"
)

// slackoverflow run
// Execute slackoverflow once
type cmdRun struct {
	KeepAlive bool `long:"keep-alive" description:"Keep on rumning every minute"`
}

// Execute
func (a *cmdRun) Execute(args []string) error {
	if a.KeepAlive {
		slackoverflow.Quiet = true
		srv := cron.New()
		srv.AddFunc("@every 1m", runFull)
		go srv.Start()
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt, os.Kill)
		<-sig
	} else {
		runFull()
	}
	return nil
}

func runFull() {

	var args []string

	// Refresh the session before running this command
	slackoverflow.SessionRefresh(true)
	stackechangeSync := cmdStackExchangeQuestions{}
	stackechangeSync.Sync = true
	stackechangeSync.Execute(args)

	slackSync := cmdSlackQuestions{}
	slackSync.All = true
	slackSync.Execute(args)
}
