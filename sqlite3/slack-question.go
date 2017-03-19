package sqlite3

import "time"

// SlackQuestion table
// Records in this table keep track of qustions between Stack Exchange and Slack
type SlackQuestion struct {
	// Id of StackExchangeQuestion question
	QID     int
	Channel string
	TS      time.Time
}

// GetSchema return table schema for StackQuestionLink
func (SlackQuestion) GetSchema() string {
	return `
      CREATE TABLE IF NOT EXISTS "SlackQuestion" (
        "QID" INTEGER,
        "channel" TEXT,
        "ts" TIMESTAMP
      )`
}
