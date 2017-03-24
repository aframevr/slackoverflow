package sqlite3

import (
	"fmt"
	"log"
)

// SlackQuestion table
// Records in this table keep track of qustions between Stack Exchange and Slack
type SlackQuestion struct {
	// Id of StackExchangeQuestion question
	QID     int
	Channel string
	TS      string
}

// GetSchema return table schema for StackQuestionLink
func (SlackQuestion) GetSchema() string {
	return `
      CREATE TABLE IF NOT EXISTS "SlackQuestion" (
        "QID" INTEGER,
        "channel" TEXT,
        "ts" TEXT
      )`
}

// Find Question link
func (SlackQuestion) Find(QID int) SlackQuestion {
	q := SlackQuestion{}
	stmt, err := sqlite.DB.Prepare(`SELECT * FROM SlackQuestion WHERE QID = ?`)
	defer stmt.Close()
	if err != nil {
		return q
	}
	_ = stmt.QueryRow(QID).Scan(
		&q.QID,
		&q.Channel,
		&q.TS,
	)
	return q
}

// Create new Question
func (slq *SlackQuestion) Create() (msg string, err error) {
	tx, err := sqlite.DB.Begin()
	if err != nil {
		return msg, err
	}

	stmt, err := sqlite.DB.Prepare(`INSERT INTO SlackQuestion
      (QID, Channel, TS)
      VALUES($1,$2,$3);`)

	if err != nil {
		return msg, err
	}

	if _, err := stmt.Exec(
		slq.QID,
		slq.Channel,
		slq.TS,
	); err != nil {
		tx.Rollback()
		return "Error had to Rollback question table", err
	}
	defer stmt.Close()
	return fmt.Sprintf("Question link: %d created.", slq.QID), nil

}

// GetAll linked questions
func (SlackQuestion) GetAll() (links []SlackQuestion, count int) {

	count = 0

	stmt, err := sqlite.DB.Prepare(`SELECT * FROM SlackQuestion ORDER BY ts DESC`)
	if err != nil {
		return links, count
	}
	rows, err := stmt.Query()
	if err != nil {
		return links, count
	}
	defer stmt.Close()

	for rows.Next() {
		ql := SlackQuestion{}
		err = rows.Scan(
			&ql.QID,
			&ql.Channel,
			&ql.TS)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, ql)
		count++
	}

	return links, count
}

// Delete Question by ID
func (slq *SlackQuestion) Delete() error {
	stmt, err := sqlite.DB.Prepare(`DELETE FROM SlackQuestion WHERE QID = ?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(slq.QID)
	defer stmt.Close()
	return err
}
