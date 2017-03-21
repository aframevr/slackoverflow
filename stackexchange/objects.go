package stackexchange

// QuestionsWrapperObj is a API response type
type QuestionsWrapperObj struct {
	Backoff        int           `json:"backoff"`
	ErrorID        int           `json:"error_id"`
	ErrorName      string        `json:"error_name"`
	ErrorMessage   string        `json:"error_message"`
	HasMore        bool          `json:"has_more"`
	Page           int           `json:"page"`
	QuotaMax       int           `json:"quota_max"`
	QuotaRemaining int           `json:"quota_remaining"`
	Items          []QuestionObj `json:"items"`
}

// QuestionObj is question item returned by StackExchange API
type QuestionObj struct {
	QID              int            `json:"question_id"`
	Title            string         `json:"title"`
	CreationDate     int64          `json:"creation_date"`
	LastActivityDate int64          `json:"last_activity_date"`
	Owner            ShallowUserObj `json:"owner"`
	IsAnswered       bool           `json:"is_answered"`
	ShareLink        string         `json:"share_link"`
	ClosedReason     string         `json:"closed_reason"`
	Tags             []string       `json:"tags"`
	Score            int            `json:"score"`
	ViewCount        int            `json:"view_count"`
	AnswerCount      int            `json:"answer_count"`
	CommentCount     int            `json:"comment_count"`
	UpVoteCount      int            `json:"up_vote_count"`
	DownVoteCount    int            `json:"down_vote_count"`
	DeleteVoteCount  int            `json:"delete_vote_count"`
	FavoriteCount    int            `json:"favorite_count"`
	ReOpenVoteCount  int            `json:"reopen_vote_count"`
}

// BadgeCountsObj of Stack Exchange User
type BadgeCountsObj struct {
	Bronze int `json:"bronze"`
	Silver int `json:"silver"`
	Gold   int `json:"gold"`
}

// ShallowUserObj returned by Stack Exchange API
type ShallowUserObj struct {
	UID          int            `json:"user_id"`
	Reputation   int            `json:"reputation"`
	ProfileImage string         `json:"profile_image"`
	DisplayName  string         `json:"display_name"`
	Link         string         `json:"link"`
	AcceptRate   int            `json:"accept_rate"`
	BadgeCounts  BadgeCountsObj `json:"badge_counts"`
}
