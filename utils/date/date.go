package date

import "time"

const (
	DateFormat = "2006-01-02T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	now := GetNow()
	return now.Format(DateFormat)
}
