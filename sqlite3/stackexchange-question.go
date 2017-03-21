package sqlite3

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aframevr/slackoverflow/stackexchange"
)

// StackExchangeQuestion table
type StackExchangeQuestion struct {
	QID              int
	UID              int
	Title            string
	CreationDate     time.Time
	LastActivityDate time.Time
	ShareLink        string
	ClosedReason     string
	Tags             string
	Site             string
	IsAnswered       bool
	Score            int
	ViewCount        int
	AnswerCount      int
	CommentCount     int
	UpVoteCount      int
	DownVoteCount    int
	DeleteVoteCount  int
	FavoriteCount    int
	ReOpenVoteCount  int
}

// GetSchema return table schema for StackQuestion
func (StackExchangeQuestion) GetSchema() string {
	return `
      CREATE TABLE IF NOT EXISTS "StackExchangeQuestion" (
        "QID" INTEGER PRIMARY KEY,
        "UID" INTEGER,
        "title" TEXT,
        "creationDate" TIMESTAMP,
        "lastActivityDate" TIMESTAMP,
        "shareLink" TEXT,
        "closedReason" TEXT,
        "tags" TEXT,
        "site" TEXT,
        "isAnswered" INTEGER,
        "score" INTEGER,
        "viewCount" INTEGER,
        "answerCount" INTEGER,
        "commentCount" INTEGER,
        "upVoteCount" INTEGER,
        "downVoteCount" INTEGER,
        "deleteVoteCount" INTEGER,
        "favoriteCount" INTEGER,
        "reOpenVoteCount" INTEGER)`
}

// Latest question from database
func (StackExchangeQuestion) Latest() StackExchangeQuestion {
	q := StackExchangeQuestion{}
	sqlite.DB.QueryRow("SELECT * FROM StackExchangeQuestion ORDER BY QID DESC LIMIT 1").Scan(
		&q.QID,
		&q.UID,
		&q.Title,
		&q.CreationDate,
		&q.LastActivityDate,
		&q.ShareLink,
		&q.ClosedReason,
		&q.Tags,
		&q.Site,
		&q.IsAnswered,
		&q.Score,
		&q.ViewCount,
		&q.AnswerCount,
		&q.CommentCount,
		&q.UpVoteCount,
		&q.DownVoteCount,
		&q.DeleteVoteCount,
		&q.FavoriteCount,
		&q.ReOpenVoteCount,
	)
	return q
}

// Find Question by ID
func (StackExchangeQuestion) Find(QID int) StackExchangeQuestion {
	q := StackExchangeQuestion{}
	stmt, err := sqlite.DB.Prepare(`SELECT * FROM StackExchangeQuestion WHERE QID = ?`)
	defer stmt.Close()
	if err != nil {
		return q
	}
	_ = stmt.QueryRow(QID).Scan(
		&q.QID,
		&q.UID,
		&q.Title,
		&q.CreationDate,
		&q.LastActivityDate,
		&q.ShareLink,
		&q.ClosedReason,
		&q.Tags,
		&q.Site,
		&q.IsAnswered,
		&q.Score,
		&q.ViewCount,
		&q.AnswerCount,
		&q.CommentCount,
		&q.UpVoteCount,
		&q.DownVoteCount,
		&q.DeleteVoteCount,
		&q.FavoriteCount,
		&q.ReOpenVoteCount,
	)
	return q
}

// SyncFromSite create or update question recieved from defined site
func (seq StackExchangeQuestion) SyncFromSite(question stackexchange.QuestionObj, site string) (msg string, err error) {
	existingQuestion := seq.Find(question.QID)

	// Map the Question
	seq.QID = question.QID
	seq.UID = question.Owner.UID
	seq.Title = question.Title
	seq.CreationDate = time.Unix(question.CreationDate, 0).UTC()
	seq.LastActivityDate = time.Unix(question.LastActivityDate, 0).UTC()
	seq.ShareLink = question.ShareLink
	seq.ClosedReason = question.ClosedReason
	seq.Tags = strings.Join(question.Tags, ";")
	seq.Site = site
	seq.IsAnswered = question.IsAnswered
	seq.Score = question.Score
	seq.ViewCount = question.ViewCount
	seq.AnswerCount = question.AnswerCount
	seq.CommentCount = question.CommentCount
	seq.UpVoteCount = question.UpVoteCount
	seq.DownVoteCount = question.DownVoteCount
	seq.DeleteVoteCount = question.DeleteVoteCount
	seq.FavoriteCount = question.FavoriteCount
	seq.ReOpenVoteCount = question.ReOpenVoteCount

	// If there is no update needed
	if existingQuestion == seq {
		return fmt.Sprintf("Question: %d is already up to date.", seq.QID), err
	}

	// Create new user or update existing one
	if existingQuestion.QID > 0 {
		msg, err = seq.Update()
	} else {
		msg, err = seq.Create()
	}
	return msg, err
}

// Create new Question
func (seq *StackExchangeQuestion) Create() (msg string, err error) {
	tx, err := sqlite.DB.Begin()
	if err != nil {
		return msg, err
	}

	stmt, err := sqlite.DB.Prepare(`INSERT INTO StackExchangeQuestion
      (QID, UID, title, creationDate, lastActivityDate, shareLink, closedReason,
        tags, site, isAnswered, score, viewCount, answerCount, commentCount,
        upVoteCount, downVoteCount, deleteVoteCount, favoriteCount, reOpenVoteCount)
      VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19);`)

	if err != nil {
		return msg, err
	}

	if _, err := stmt.Exec(
		seq.QID,
		seq.UID,
		seq.Title,
		seq.CreationDate,
		seq.LastActivityDate,
		seq.ShareLink,
		seq.ClosedReason,
		seq.Tags,
		seq.Site,
		seq.IsAnswered,
		seq.Score,
		seq.ViewCount,
		seq.AnswerCount,
		seq.CommentCount,
		seq.UpVoteCount,
		seq.DownVoteCount,
		seq.DeleteVoteCount,
		seq.FavoriteCount,
		seq.ReOpenVoteCount,
	); err != nil {
		tx.Rollback()
		return "Error had to Rollback question table", err
	}
	defer stmt.Close()
	return fmt.Sprintf("Question: %d created.", seq.QID), nil

}

// Update existsing question
func (seq *StackExchangeQuestion) Update() (msg string, err error) {
	tx, err := sqlite.DB.Begin()
	if err != nil {
		return msg, err
	}
	stmt, err := sqlite.DB.Prepare(`UPDATE StackExchangeQuestion SET
    UID=?, title=?, creationDate=?, lastActivityDate=?, shareLink=?, closedReason=?,
      tags=?, site=?, isAnswered=?, score=?, viewCount=?, answerCount=?, commentCount=?,
      upVoteCount=?, downVoteCount=?, deleteVoteCount=?, favoriteCount=?, reOpenVoteCount=?
      WHERE QID=?;`)

	if err != nil {
		return msg, err
	}

	if _, err := stmt.Exec(
		seq.UID,
		seq.Title,
		seq.CreationDate,
		seq.LastActivityDate,
		seq.ShareLink,
		seq.ClosedReason,
		seq.Tags,
		seq.Site,
		seq.IsAnswered,
		seq.Score,
		seq.ViewCount,
		seq.AnswerCount,
		seq.CommentCount,
		seq.UpVoteCount,
		seq.DownVoteCount,
		seq.DeleteVoteCount,
		seq.FavoriteCount,
		seq.ReOpenVoteCount,
		seq.QID,
	); err != nil {
		tx.Rollback()
		return "Error had to Rollback question update", err
	}
	defer stmt.Close()
	return fmt.Sprintf("Question: %d updated.", seq.QID), nil
}

// TrackedIds return tracked question ids
func (seq StackExchangeQuestion) TrackedIds(qToWatch int) (ids string, count int) {

	count = 0
	ids = ""

	stmt, err := sqlite.DB.Prepare(`SELECT QID FROM StackExchangeQuestion ORDER BY creationDate DESC LIMIT ?`)
	if err != nil {
		return ids, count
	}
	rows, err := stmt.Query(qToWatch)
	if err != nil {
		return ids, count
	}
	defer stmt.Close()

	var idsMap []string
	for rows.Next() {
		var QID string
		err = rows.Scan(&QID)
		if err != nil {
			log.Fatal(err)
		}
		// in case of only one question asign that to idsMap
		ids = QID
		idsMap = append(idsMap, QID)
		count++
	}
	if count > 1 {
		ids = strings.Join(idsMap, ";")
	}

	return ids, count
}
