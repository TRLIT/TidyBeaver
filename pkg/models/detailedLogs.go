package models

import (
	"time"
)

type AdaptedLogs struct {
	AdaptedLog []AdaptedLog
}

type AdaptedLog struct {
	Level      string    `json:"Level"`
	Service    string    `json:"Service"`
	Message    string    `json:"Message"`
	Time       time.Time `json:"Time"`
	Index      int       `json:"Index"`
	EntryType  string    `json:"EntryType"`
	Source     string    `json:"Source"`
	InstanceID int       `json:"InstanceID"`
}
