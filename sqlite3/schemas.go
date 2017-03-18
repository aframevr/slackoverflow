package sqlite3

// SlackOverflowSchemas are table structures for Slack Overflow
type SlackOverflowSchemas struct {
	StackQuestion     string
	StackQuestionLink string
	StackUser         string
}

func (schemas *SlackOverflowSchemas) loadSchemas() {

	// Tracked slack questions table
	schemas.StackQuestion = `
      CREATE TABLE IF NOT EXISTS "StackQuestion" (
        "ID" INTEGER PRIMARY KEY AUTOINCREMENT,
        "UID" INTEGER,
        "QID" INTEGER,
        "commentCount" INTEGER,
        "isAnswered" INTEGER,
        "viewCount" INTEGER,
        "answerCount" INTEGER,
        "score" INTEGER,
        "creationDate" TIMESTAMP,
        "link" TEXT,
        "title" TEXT,
        "site" TEXT)`

	// Tracking table for questions
	schemas.StackQuestionLink = `
      CREATE TABLE IF NOT EXISTS "StackQuestionLink" (
        "ID" INTEGER PRIMARY KEY AUTOINCREMENT,
        "LQID" INTEGER,
        "QID" INTEGER,
        "channel" TEXT,
        "ts" TIMESTAMP,
        "isOld" INTEGER
      )`

	// StackExchange user table
	schemas.StackUser = `
      CREATE TABLE IF NOT EXISTS "StackUser" (
        "ID" INTEGER PRIMARY KEY AUTOINCREMENT,
        "UID" INTEGER,
        "userType" TEXT,
        "reputation" INTEGER,
        "profileImage" TEXT,
        "displayName" TEXT,
        "link" TEXT
      )`
}
