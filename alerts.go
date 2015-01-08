package main

import (
	"time"
)

type Alert struct {
	Id       int64  `json:"id"`
	Location string `json:"location"`
	Type     string `json:"type"`
	// where the alert was retrieved; can be either rappler.com,
	// PNP, or user entered.
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Alerts []Alert
