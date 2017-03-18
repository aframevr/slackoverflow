package slackoverflow

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strconv"
	"time"

	"github.com/aframevr/slackoverflow/slack"
	"github.com/aframevr/slackoverflow/std"
	flags "github.com/jessevdk/go-flags"
	"github.com/logrusorgru/aurora"
)

// User Session object
var stackoverflow *Application

// Connands mapper
var argv Commands

// CLi parser
var parser = flags.NewParser(&argv, flags.Default)

// Application application
type Application struct {
	cwd          string
	user         *user.User
	group        *user.Group
	startTime    time.Time
	logLevel     int
	projectPath  string
	configFile   string
	databaseFile string
	hostname     string
	config       yamlContents
	Info         info
	Slack        *slack.Slack
}

// Run stackoverflow application
func (s *Application) Run() {
	s.Banner()

	// Check configuration
	if !s.config.Validate() {
		s.config.Reconfigure()
	}

	// Handle call
	if _, err := parser.Parse(); err != nil {
		// Failure was fine since -h or --help flag was provided
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			s.Close(0)
		} else {
			// There was some errors
			Warning("Invalid or missing arguments: ")
			std.Body("See slackoverflow --help for more info")
			std.Body("See slackoverflow <command> --help for more info about specific command")
			s.Close(0)
		}
	}
}

// SetStartTime from external source
func (s *Application) SetStartTime(startTime time.Time) {
	stackoverflow.startTime = startTime
}

// Banner prints stackoverflow application Banner
func (s *Application) Banner() {
	std.Hr()
	std.Body("  %s v: %s date: %s", aurora.Bold("slackoverflow"), s.Info.GetVersion(), s.Info.GetBuildDate())
	std.Nl()
	std.Body("  Copyright © %s A-Frame authors.", s.Info.GetCopyYear())
	std.Body("  Url: https://github.com/aframevr/slackoverflow")
	std.Body("  %s", "The MIT License")
	std.Hr()
}

// Close and prinf final summary
func (s *Application) Close(code int) {
	date := time.Now().Local().Format("15:04:05 - 2 January 2006")

	fmt.Println("")
	std.Hr()
	std.Body("Execution elapsed: %s", time.Since(s.startTime))
	std.Body("Date: %s", date)
	std.Hr()
	os.Exit(code)
}

func (s *Application) SessionRefresh() {

	// Set Log Level from -v or -d flag default to config.Data.Slackoverflow.LogLevel
	UpdateLogLevel()

	// Load Slack Client
	if stackoverflow.Slack != nil {
		// Configure slack
		stackoverflow.Slack = slack.Load()
		stackoverflow.Slack.SetHost(stackoverflow.config.Slack.APIHost)
		stackoverflow.Slack.SetToken(stackoverflow.config.Slack.Token)
		stackoverflow.Slack.SetChannel(stackoverflow.config.Slack.Channel)
	} else {
		Debug("Slack Client is already loaded.")
	}

}

// Start session
func Start() *Application {
	gid := os.Getegid()
	stackoverflow = &Application{}
	stackoverflow.cwd, _ = os.Getwd()
	stackoverflow.user, _ = user.Current()
	stackoverflow.group, _ = user.LookupGroupId(strconv.Itoa(gid))
	stackoverflow.startTime = time.Now()
	stackoverflow.projectPath = path.Join(stackoverflow.user.HomeDir, ".slackoverflow")
	stackoverflow.configFile = path.Join(stackoverflow.projectPath, "slackoverflow.yaml")
	stackoverflow.databaseFile = path.Join(stackoverflow.projectPath, "slackoverflow.db3")
	stackoverflow.hostname, _ = os.Hostname()

	return stackoverflow
}
