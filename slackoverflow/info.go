package slackoverflow

import (
	"fmt"
	"time"
)

type info struct {
	version        string
	buildDate      time.Time
	firstBuildYear int
}

// SetBuildDate update application build date
func (info *info) SetBuildDate(buildDate string) {
	rfc3339 := "2006-01-02T15:04:05Z07:00"
	info.firstBuildYear = 2016
	if len(buildDate) == 0 {
		buildDate = time.Now().Local().Format(rfc3339)
	}
	info.buildDate, _ = time.Parse(rfc3339, buildDate)
}

// GetBuildDate - Get current version build date
func (info *info) GetBuildDate() string {
	return info.buildDate.Local().Format("2 Jan 2006 15:04:05")
}

// GetVersion - Get current version
func (info *info) GetVersion() string {
	return info.buildDate.Format("06.1.2")
}

// GetCopyYear - Get copy year(s)
func (info *info) GetCopyYear() string {
	buildYear := info.buildDate.Local().Year()
	copyYear := info.buildDate.Local().Format("2006")
	if buildYear > info.firstBuildYear {
		copyYear = fmt.Sprintf("%d - %d", info.firstBuildYear, buildYear)
	}
	return copyYear
}
