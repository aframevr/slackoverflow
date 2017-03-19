package slackoverflow

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strconv"
	"time"

	"github.com/aframevr/slackoverflow/slack"
	"github.com/aframevr/slackoverflow/sqlite3"
	"github.com/aframevr/slackoverflow/stackexchange"
	"github.com/aframevr/slackoverflow/std"
	flags "github.com/jessevdk/go-flags"
	"github.com/logrusorgru/aurora"
)

// User Session object
var slackoverflow *Application

// Connands mapper
var argv Commands

// CLi parser
var parser = flags.NewParser(&argv, flags.Default)

// Application application
type Application struct {
	cwd           string
	user          *user.User
	group         *user.Group
	startTime     time.Time
	logLevel      int
	projectPath   string
	configFile    string
	databaseFile  string
	hostname      string
	config        yamlContents
	Info          info
	Slack         *slack.Client
	SQLite3       *sqlite3.Client
	StackExchange *stackexchange.Client
}

// Run stackoverflow application
func (so *Application) Run() {
	so.Banner()

	// Handle call
	if _, err := parser.Parse(); err != nil {
		// Failure was fine since -h or --help flag was provided
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			so.Close(0)
		} else {
			// There was some errors
			Warning("Invalid or missing arguments: ")
			std.Body("See slackoverflow --help for more info")
			std.Body("See slackoverflow <command> --help for more info about specific command")
			so.Close(0)
		}
	}
}

// SetStartTime from external source
func (so *Application) SetStartTime(startTime time.Time) {
	so.startTime = startTime
}

// Banner prints stackoverflow application Banner
func (so *Application) Banner() {
	std.Hr()
	std.Body("  %s v: %s date: %s", aurora.Bold("slackoverflow"), so.Info.GetVersion(), so.Info.GetBuildDate())
	std.Nl()
	std.Body("  Copyright © %s A-Frame authors.", so.Info.GetCopyYear())
	std.Body("  Url: https://github.com/aframevr/slackoverflow")
	std.Body("  %s", "The MIT License")
	std.Hr()
}

// Close and prinf final summary
func (so *Application) Close(code int) {
	date := time.Now().Local().Format("15:04:05 - 2 January 2006")

	// Wait database to be closed
	if so.SQLite3 != nil {
		defer so.SQLite3.DB.Close()
	}

	fmt.Println("")
	std.Hr()
	std.Body("Execution elapsed: %s", time.Since(so.startTime))
	std.Body("Date: %s", date)
	std.Hr()
	os.Exit(code)
}

// SessionRefresh refresh session and makes sure that all deps are loaded
func (so *Application) SessionRefresh() {

	// Check configuration
	if !so.config.IsConfigured() {
		Emergency("You must execute 'slackoverflow reconfigure' or correct errors in ~/.slackoverflow/slackoverflow.yaml")
	}

	// Set Log Level from -v or -d flag default to config.Data.SlackOverflow.LogLevel
	UpdateLogLevel()

	// Load Slack Client
	if so.Slack == nil {
		// Configure slack
		so.Slack = slack.Load()
		so.Slack.SetHost(so.config.Slack.APIHost)
		so.Slack.SetToken(so.config.Slack.Token)
		so.Slack.SetChannel(so.config.Slack.Channel)
		Debug("Slack Client is loaded.")
	} else {
		Debug("Slack Client is already loaded.")
	}

	// Load SQLite3 Client
	if so.SQLite3 == nil {
		var err error
		so.SQLite3, err = sqlite3.Load(so.databaseFile)
		// Kill the session if we can not open database file
		if err != nil {
			Emergency(err.Error())
		}
		Ok("SQLite3 Database loaded: %s", so.databaseFile)

		// Check does the StackQuestion table exist
		err = so.SQLite3.VerifyTable("StackExchangeQuestion")
		if err != nil {
			Emergency("Table StackExchangeQuestion: %q", err)
		}
		Ok("Table: StackExchangeQuestion exists.")

		// Check does the SlackQuestion table exist
		err = so.SQLite3.VerifyTable("SlackQuestion")
		if err != nil {
			Emergency("Table SlackQuestion: %q", err)
		}
		Ok("Table: SlackQuestion exists.")

		// Check does the StackExchangeUser table exist
		err = so.SQLite3.VerifyTable("StackExchangeUser")
		if err != nil {
			Emergency("Table StackExchangeUser: %q", err)
		}
		Ok("Table: StackExchangeUser exists.")
	} else {
		Ok("SQLite3 Database is already loaded.")
	}

	// Load Stack Exchange Client
	// Load Slack Client
	if so.StackExchange == nil {
		// Configure slack
		so.StackExchange = stackexchange.Load()
		so.StackExchange.SetHost(so.config.StackExchange.APIHost)
		so.StackExchange.SetAPIVersion(so.config.StackExchange.APIVersion)
		so.StackExchange.SetAPIKey(so.config.StackExchange.Key)

		Debug("Stack Exchange Client is loaded.")
	} else {
		Debug("Stack Exchange Client is already loaded.")
	}
}

// Debugging Either is debugging enabled or not
func (so *Application) Debugging() bool {
	return so.logLevel == 100
}

// Start session
func Start() *Application {
	gid := os.Getegid()
	slackoverflow = &Application{}
	slackoverflow.cwd, _ = os.Getwd()
	slackoverflow.user, _ = user.Current()
	slackoverflow.group, _ = user.LookupGroupId(strconv.Itoa(gid))
	slackoverflow.startTime = time.Now()
	slackoverflow.projectPath = path.Join(slackoverflow.user.HomeDir, ".slackoverflow")
	slackoverflow.configFile = path.Join(slackoverflow.projectPath, "slackoverflow.yaml")
	slackoverflow.databaseFile = path.Join(slackoverflow.projectPath, "slackoverflow.db3")
	slackoverflow.hostname, _ = os.Hostname()

	return slackoverflow
}
