package models

import (
	"time"
)

const (
	timeLayout = time.RFC822
	timeFormat = "15:04:05.00000"
	dateFormat = "2006/01/02"
)

type Event struct {
	Id            int    `json:"id"`
	Uid           int    `json:"user_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	TimestampFrom int64  `json:"timestamp_from"`
	TimestampTo   int64  `json:"timestamp_to"`
}

type EventResult struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Timezone string `json:"timezone"`
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
	TimeFrom string `json:"timeFrom"`
	TimeTo   string `json:"timeTo"`
}

func (*Event) TimestampToDateTime(t int64, timezone string) (string, string, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return "", "", err
	}

	unixTime := time.Unix(t, 0)
	timeString := unixTime.In(location).Format(timeFormat)
	dateString := unixTime.In(location).Format(dateFormat)
	return dateString, timeString, err
}

func (*Event) TimeToTimestamp(t, timezone string) (int64, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}

	tzName, _ := time.Now().In(location).Zone()
	form := t + " " + tzName

	timeUnix, err := time.Parse(timeLayout, form)
	if err != nil {
		return 0, err
	}

	return timeUnix.Unix(), err
}
