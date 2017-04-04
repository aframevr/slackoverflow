package slackoverflow

// slackoverflow stackexchange
type cmdStackExchange struct {
	Questions cmdStackExchangeQuestions `command:"questions" description:"Work with stackexchange questions based on the config"`
}
