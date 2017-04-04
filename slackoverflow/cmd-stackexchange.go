package slackoverflow

// slackoverflow stackexchange
type cmdStackExchange struct {
	Questions cmdStackExchangeQuestions `command:"questions" description:"Work with stackexchange questions based on the config"`
	Watch     cmdStackExchangeWatch     `command:"watch" description:"Watch new questions from Stack Exchange site (updated every minute nothing stored to db or posted to slack)"`
}
