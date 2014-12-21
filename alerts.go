package main

import (
	"time"
)

type Alert struct {
	Id        int64
	Location  string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
