package sqlite3

// StackExchangeUser table
type StackExchangeUser struct {
	UID          int
	UserType     string
	Reputation   int
	ProfileImage string
	DisplayName  string
	Link         string
}

// GetSchema return table schema for StackUser
func (StackExchangeUser) GetSchema() string {
	return `
      CREATE TABLE IF NOT EXISTS "StackExchangeUser" (
        "UID" INTEGER PRIMARY KEY,
        "userType" TEXT,
        "reputation" INTEGER,
        "profileImage" TEXT,
        "displayName" TEXT,
        "link" TEXT
      )`
}
