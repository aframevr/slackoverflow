package sqlite3

import "time"

// StackExchangeQuestion table
type StackExchangeQuestion struct {
	QID          int
	UID          int
	CommentCount int
	IsAnswered   bool
	ViewCount    int
	AnswerCount  int
	Score        int
	CreationDate time.Time
	Link         string
	Title        string
	Site         string
}

// GetSchema return table schema for StackQuestion
func (StackExchangeQuestion) GetSchema() string {
	return `
      CREATE TABLE IF NOT EXISTS "StackExchangeQuestion" (
        "QID" INTEGER PRIMARY KEY,
        "UID" INTEGER,
        "commentCount" INTEGER,
        "isAnswered" INTEGER,
        "viewCount" INTEGER,
        "answerCount" INTEGER,
        "score" INTEGER,
        "creationDate" TIMESTAMP,
        "link" TEXT,
        "title" TEXT,
        "site" TEXT)`
}

// Latest question from database
func (q *StackExchangeQuestion) Latest() error {
	err := sqlite.DB.QueryRow("SELECT * FROM StackExchangeQuestion ORDER BY ID DESC LIMIT 1").Scan(
		&q.QID,
		&q.UID,
		&q.CommentCount,
		&q.IsAnswered,
		&q.ViewCount,
		&q.AnswerCount,
		&q.Score,
		&q.CreationDate,
		&q.Link,
		&q.Title,
		&q.Site,
	)
	return err
}
