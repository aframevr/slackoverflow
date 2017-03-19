package slackoverflow

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aframevr/slackoverflow/slack"
	"github.com/aframevr/slackoverflow/std"

	yaml "github.com/go-yaml/yaml"
)

// Yaml mapping to yaml configuration
type yamlContents struct {
	SlackOverflow struct {
		LogLevel string `yaml:"log_level"`
		Watch    int    `yaml:"watch"`
	} `yaml:"slackoverflow"`
	Slack struct {
		Enabled bool   `yaml:"enabled"`
		Channel string `yaml:"channel"`
		Token   string `yaml:"token"`
		APIHost string `yaml:"api-host"`
	} `yaml:"slack"`
	StackExchange struct {
		Enabled        bool              `yaml:"enabled"`
		Key            string            `yaml:"key"`
		APIVersion     string            `yaml:"api-version"`
		APIHost        string            `yaml:"api-host"`
		SearchAdvanced map[string]string `yaml:"search-advanced"`
	} `yaml:"stackexchange"`
}

// readConfig - Read yaml into struct
func (yc *yamlContents) readConfig(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	inputBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(inputBytes, &yc); err != nil {
		return err
	}
	yc.SlackOverflow.LogLevel = strings.ToUpper(yc.SlackOverflow.LogLevel)
	return nil
}

// Validate configuration
func (yc *yamlContents) IsConfigured() bool {
	if err := yc.readConfig(slackoverflow.configFile); err != nil {
		Warning("Stackoverflow is not configured")
		return false
	}
	return yc.SlackOverflow.LogLevel != ""
}

// Save saves configuration contents to ~/.slackoverflow/slackoverflow.yaml
func (yc *yamlContents) Save() {
	// Save configFile
	std.Hr()
	contents, _ := yaml.Marshal(&yc)
	err := ioutil.WriteFile(slackoverflow.configFile, []byte(contents), 0644)
	if err != nil {
		Emergency("Failed to write: %s", slackoverflow.configFile)
	}

	Ok("Configuration saved to: %s", slackoverflow.configFile)
}

// ReconfigureAll Create new configuration file or replace existsing
func (yc *yamlContents) ReconfigureAll() {
	Info("Configuration location set to: %s", slackoverflow.projectPath)
	os.MkdirAll(slackoverflow.projectPath, os.ModePerm)
	std.Hr()

	configureStackExchange := std.AskForConfirmation("Would you like to Configure StackExchange?")
	if configureStackExchange {
		yc.ConfigureStackExchange()
	}

	configureSlack := std.AskForConfirmation("Would you like to Configure Slack?")
	if configureSlack {
		yc.ConfigureSlack()
	}

	configureSlackOverflow := std.AskForConfirmation("Would you like to Configure Slack Overflow?")
	if configureSlackOverflow {
		yc.ConfigureSlackOverflow()
	}
}

// ConfigureSlackOverflow set general configuration options
func (yc *yamlContents) ConfigureSlackOverflow() {
	std.Hr()
	std.Body("Configuring Slack Overflow")

	// Force to reload the configuration file if exists
	yc.IsConfigured()

	reader := bufio.NewReader(os.Stdin)
	// Log level

	std.Hr()
	std.Body("Set stackoverflow logleve and verbosity. valid options are:")
	std.Nl()
	std.Body("notice")
	std.Body("info")
	std.Body("debug")
	std.Hr()
	std.Body("Make your choice!")
	logLevel, _ := reader.ReadString('\n')
	std.Hr()
	yc.SlackOverflow.LogLevel = strings.TrimSpace(logLevel)

	// Number of questions to watch
	std.Hr()
	std.Body("Set the value for how many latest questions you want to track and update.")
	std.Body("Good value is (25) which means that besides checking new qustions in defined stack exchange site")
	std.Body("also last (n) questions will be checked for comment count, view count, answer count, score and is question accepted or not.")
	std.Body("Emoijs of these stats will be removed from older than (n) questions.")
	std.Hr()
	fmt.Scan(&yc.SlackOverflow.Watch)
	yc.Save()
	Ok("Slack Overflow is configured")
	std.Hr()
}

// ConfigureStackExchange se Stack Exchange configuration options
func (yc *yamlContents) ConfigureStackExchange() {
	std.Hr()
	std.Body("Configuring Stack Exchange API Client")

	// Force to reload the configuration file if exists
	yc.IsConfigured()

	reader := bufio.NewReader(os.Stdin)

	if yc.StackExchange.SearchAdvanced == nil {
		yc.StackExchange.SearchAdvanced = make(map[string]string)
	}

	// Set endpoint defaults
	yc.StackExchange.Enabled = true
	yc.StackExchange.APIHost = "https://api.stackexchange.com"
	yc.StackExchange.APIVersion = "2.2"

	// Set Stack Exchange site
	std.Hr()
	std.Body("Name one Stack Exchange site where you want to track questions from e.g: stackoverflow.")
	std.Body("For full list of available sites check: http://stackexchange.com/sites")
	std.Hr()
	site, _ := reader.ReadString('\n')
	yc.StackExchange.SearchAdvanced["site"] = strings.TrimSpace(site)

	// Set tagged parameter value for Stack Exchange search advanced
	std.Hr()
	std.Body("Set tag which you want to track from selected site")
	std.Body("You can also set multible tags separated with (;) e.g: aframe;three.js")
	std.Hr()
	tagged, _ := reader.ReadString('\n')
	yc.StackExchange.SearchAdvanced["tagged"] = strings.TrimSpace(tagged)

	// stackexchange clientKey
	std.Hr()
	std.Body("Without having Stack Exchange API APP key's you can make 300 requests per day.")
	std.Body("When you register for an Stack Exchange API App Key you can make 10000 requests per day")
	std.Body("You can register for an APP KEY here: http://stackapps.com/apps/oauth/register")
	registerStackExchange := std.AskForConfirmation("Do you want to set keys now?")
	std.Hr()
	if registerStackExchange {
		std.Body("Enter your Client Key")
		clientKey, _ := reader.ReadString('\n')
		yc.StackExchange.Key = strings.TrimSpace(clientKey)
	}
	yc.Save()
	Ok("Slack Stack Exchange API Client is configured")
	std.Hr()
}

// Slack Configure slack
func (yc *yamlContents) ConfigureSlack() {
	std.Hr()
	std.Body("Configuring Slack API Client")

	// Force to reload the configuration file if exists
	yc.IsConfigured()

	reader := bufio.NewReader(os.Stdin)

	// Set Slack Defaults
	yc.Slack.APIHost = "https://slack.com/api"
	yc.Slack.Enabled = true

	// slack token
	std.Hr()
	std.Body("Enter your @stackoverflow Slack BOT API Token.")
	std.Body("You can create a bot and get token at https://<your-team>.slack.com/apps/manage/custom-integrations")
	std.Hr()
	slackToken, _ := reader.ReadString('\n')
	yc.Slack.Token = strings.TrimSpace(slackToken)

	// Print the available channel list
	std.Hr()
	Info("Fetching available Slack channels.")
	slackoverflow.Slack = slack.Load()
	slackoverflow.Slack.SetHost(yc.Slack.APIHost)
	slackoverflow.Slack.SetToken(yc.Slack.Token)
	hasChannels := slackoverflow.Slack.ListChannels()
	if !hasChannels {
		Emergency("Unable to fetch any channels with provided credentials")
	}

	// Slack channel
	std.Hr()
	std.Body("Enter Channel ID which you want to post the questions")
	std.Hr()
	slackChannel, _ := reader.ReadString('\n')
	yc.Slack.Channel = strings.TrimSpace(slackChannel)

	yc.Save()
	Ok("Slack Slack API Client is configured")
	std.Hr()
}
