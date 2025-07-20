package repository

import (
	"time"
)

type SubscriptionFilter struct {
	UserID      string
	ServiceName string
	From        time.Time
	To          time.Time
}

type SumFilter struct {
	UserID      string
	ServiceName string
	From        time.Time
	To          time.Time
}
