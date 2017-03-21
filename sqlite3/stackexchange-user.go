package sqlite3

import "github.com/aframevr/slackoverflow/stackexchange"

// StackExchangeUser table
type StackExchangeUser struct {
	UID          int
	DisplayName  string
	ProfileImage string
	Link         string
	Reputation   int
	AcceptRate   int
	BadgeBronze  int
	BadgeSilver  int
	BadgeGold    int
}

// GetSchema return table schema for StackUser
func (StackExchangeUser) GetSchema() string {
	return `
      CREATE TABLE IF NOT EXISTS "StackExchangeUser" (
        "UID" INTEGER PRIMARY KEY,
        "displayName" TEXT,
        "profileImage" TEXT,
        "link" TEXT,
        "reputation" INTEGER,
        "acceptRate" INTEGER,
        "badgeBronze" INTEGER,
        "badgeSilver" INTEGER,
        "badgeGold" INTEGER
      )`
}

// Find User by ID
func (StackExchangeUser) Find(UID int) StackExchangeUser {
	user := StackExchangeUser{}
	stmt, err := sqlite.DB.Prepare(`SELECT * FROM StackExchangeUser WHERE UID = ?`)
	defer stmt.Close()
	if err != nil {
		return user
	}
	_ = stmt.QueryRow(UID).Scan(
		&user.UID,
		&user.DisplayName,
		&user.ProfileImage,
		&user.Link,
		&user.Reputation,
		&user.AcceptRate,
		&user.BadgeBronze,
		&user.BadgeSilver,
		&user.BadgeGold,
	)
	return user
}

// SyncShallowUser Create or update Stack Exchange User
func (seu StackExchangeUser) SyncShallowUser(user stackexchange.ShallowUserObj) (msg string, err error) {

	existingUser := seu.Find(user.UID)
	// Map the user
	seu.UID = user.UID
	seu.DisplayName = user.DisplayName
	seu.ProfileImage = user.ProfileImage
	seu.Link = user.Link
	seu.Reputation = user.Reputation
	seu.AcceptRate = user.AcceptRate
	seu.BadgeBronze = user.BadgeCounts.Bronze
	seu.BadgeSilver = user.BadgeCounts.Silver
	seu.BadgeGold = user.BadgeCounts.Gold

	// If there is no update needed
	if existingUser == seu {
		return "User: " + seu.DisplayName + " is already up to date", err
	}

	// Create new user or update existing one
	if existingUser.UID > 0 {
		msg, err = seu.Update()
	} else {
		msg, err = seu.Create()
	}
	return msg, err
}

// Create new User
func (seu *StackExchangeUser) Create() (msg string, err error) {
	tx, err := sqlite.DB.Begin()
	if err != nil {
		return msg, err
	}

	stmt, err := sqlite.DB.Prepare(`INSERT INTO StackExchangeUser
      (UID, displayName, profileImage, link, reputation, acceptRate, badgeBronze, badgeSilver, badgeGold)
      VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9);`)
	defer stmt.Close()
	if err != nil {
		return msg, err
	}

	if _, err := stmt.Exec(
		seu.UID,
		seu.DisplayName,
		seu.ProfileImage,
		seu.Link,
		seu.Reputation,
		seu.AcceptRate,
		seu.BadgeBronze,
		seu.BadgeSilver,
		seu.BadgeGold,
	); err != nil {
		tx.Rollback()
		return "Error had to Rollback", err
	}
	return "User " + seu.DisplayName + " created.", nil
}

// Update existsing User
func (seu *StackExchangeUser) Update() (msg string, err error) {
	tx, err := sqlite.DB.Begin()
	if err != nil {
		return msg, err
	}
	stmt, err := sqlite.DB.Prepare(`UPDATE StackExchangeUser SET
      displayName=?, profileImage=?, link=?, reputation=?, acceptRate=?, badgeBronze=?, badgeSilver=?, badgeGold=?
      WHERE UID=?;`)

	if err != nil {
		return msg, err
	}

	if _, err := stmt.Exec(
		seu.DisplayName,
		seu.ProfileImage,
		seu.Link,
		seu.Reputation,
		seu.AcceptRate,
		seu.BadgeBronze,
		seu.BadgeSilver,
		seu.BadgeGold,
		seu.UID,
	); err != nil {
		tx.Rollback()
		return "Error had to Rollback", err
	}
	defer stmt.Close()
	return "User :" + seu.DisplayName + " updated.", nil
}
