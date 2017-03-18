package slackoverflow

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aframevr/slackoverflow/std"

	yaml "gopkg.in/yaml.v2"
)

// Yaml mapping to yaml configuration
type yamlContents struct {
	SlackOverflow struct {
		LogLevel string `yaml:"log_level"`
		Watch    int    `yaml:"watch"`
	} `yaml:"slackoverflow"`
	Slack struct {
		Channel string `yaml:"channel"`
		Token   string `yaml:"token"`
		APIHost string `yaml:"api-host"`
	} `yaml:"slack"`
	StackExchange struct {
		Key        string `yaml:"key"`
		APIVersion string `yaml:"api-version"`
		APIHost    string `yaml:"api-host"`
		Parameters struct {
			Site   string `yaml:"site"`
			Tagged string `yaml:"tagged"`
		} `yaml:"parameters"`
	} `yaml:"stackexchange"`
}

// Validate configuration
func (yamlContents *yamlContents) Validate() bool {
	if err := yamlContents.readConfig(slackoverflow.configFile); err != nil {
		Warning("Stackoverflow is not configured")
		return false
	}
	return true
}

// Create new configuration file
func (yamlContents *yamlContents) Reconfigure() {
	Info("Configuration location set to: %s", slackoverflow.projectPath)
	os.MkdirAll(slackoverflow.projectPath, os.ModePerm)

	reader := bufio.NewReader(os.Stdin)

	// Log level
	std.Hr()
	Notice("Set stackoverflow logleve and verbosity. valid options are:")
	std.Hr()
	std.Msg("notice")
	std.Msg("info")
	std.Msg("debug")
	std.Hr()
	Info("Make your choice!")
	logLevel, _ := reader.ReadString('\n')
	yamlContents.SlackOverflow.LogLevel = strings.TrimSpace(logLevel)

	// stackexchange
	yamlContents.StackExchange.APIHost = "https://api.stackexchange.com"
	yamlContents.StackExchange.APIVersion = "2.2"

	// stackexchange site
	std.Hr()
	Notice("Name one Stack Exchange site where you want to track questions from e.g: stackoverflow.")
	std.Msg("For full list of available sites check: http://stackexchange.com/sites")
	std.Hr()
	site, _ := reader.ReadString('\n')
	yamlContents.StackExchange.Parameters.Site = strings.TrimSpace(site)

	// stackexchange tagged
	std.Hr()
	Notice("Set tag which you want to track from selected site")
	std.Msg("You can also set multible tags separated with (;) e.g: aframe;three.js")
	tagged, _ := reader.ReadString('\n')
	yamlContents.StackExchange.Parameters.Tagged = strings.TrimSpace(tagged)

	// stackexchange clientKey
	std.Hr()
	Notice("Without having Stack Exchange API APP key's you can make 300 requests per day.")
	std.Msg("When you register for an Stack Exchange API App Key you can make 10000 requests per day")
	std.Msg("You can register for an APP KEY here: http://stackapps.com/apps/oauth/register")
	registerStackExchange := std.AskForConfirmation("Do you want to set keys now?")
	if registerStackExchange {
		Notice("Enter your Client Key")
		clientKey, _ := reader.ReadString('\n')
		yamlContents.StackExchange.Key = strings.TrimSpace(clientKey)
	}

	// Slack
	yamlContents.Slack.APIHost = "https://slack.com/api"

	// slack token
	Notice("Enter your @stackoverflow Slack BOT API Token.")
	std.Msg("You can create a bot and get token at https://<your-team>.slack.com/apps/manage/custom-integrations")
	slackToken, _ := reader.ReadString('\n')
	yamlContents.Slack.Token = strings.TrimSpace(slackToken)

	slackoverflow.Slack.SetHost(yamlContents.Slack.APIHost)
	slackoverflow.Slack.SetToken(yamlContents.Slack.Token)

	// Print the available channel list
	hasChannels := slackoverflow.Slack.ListChannels()
	if !hasChannels {
		Emergency("Unable to fetch any channels with provided credentials")
	}

	// Slack channel
	Notice("Enter Channel ID which you want to post the questions")
	slackChannel, _ := reader.ReadString('\n')
	yamlContents.Slack.Channel = strings.TrimSpace(slackChannel)
	slackoverflow.Slack.SetChannel(yamlContents.Slack.Channel)

	// Number of questions to watch
	std.Hr()
	Notice("Set the value for how many latest questions you want to track and update.")
	std.Body("Good value is (25) which means that besides checking new qustions in defined stack exchange site")
	std.Body("also last (n) questions will be checked for comment count, view count, answer count, score and is question accepted or not.")
	std.Body("Emoijs of these stats will be removed from older than (n) questions.")
	std.Hr()
	fmt.Scan(&yamlContents.SlackOverflow.Watch)

	// Save configFile
	std.Hr()
	contents, _ := yaml.Marshal(&yamlContents)
	err := ioutil.WriteFile(slackoverflow.configFile, []byte(contents), 0644)
	if err != nil {
		Emergency("Failed to write: %s", slackoverflow.configFile)
	}

	Ok("Configuration saved: %s", slackoverflow.configFile)

}

// readConfig - Read yaml into struct
func (yamlContents *yamlContents) readConfig(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	inputBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(inputBytes, &yamlContents); err != nil {
		return err
	}
	yamlContents.SlackOverflow.LogLevel = strings.ToUpper(yamlContents.SlackOverflow.LogLevel)
	return nil
}
