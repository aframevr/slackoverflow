package slackoverflow

// Yaml mapping to yaml configuration
type yamlContents struct {
	Slackoverflow struct {
		LogLevel string `yaml:"log_level"`
		Watch    int    `yaml:"watch"`
	} `yaml:"slackoverflow"`
	Slack struct {
		Channel string `yaml:"channel"`
		Token   string `yaml:"token"`
		APIHost string `yaml:"api-host"`
	} `yaml:"slack"`
	StackExchange struct {
		ClientID     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
		Key          string `yaml:"key"`
		APIVersion   string `yaml:"api-version"`
		APIHost      string `yaml:"api-host"`
		Parameters   struct {
			Site   string `yaml:"site"`
			Tagged string `yaml:"tagged"`
		} `yaml:"parameters"`
	} `yaml:"stackexchange"`
}
