package model

import "time"

type Subscription struct {
	ID          string     `gorm:"type:uuid;default:gen_random_uuid();primary_key" json:"id"`
	ServiceName string     `dorm:"not null" json:"service_name"`
	Price       int        `dorm:"not null" json:"price"`
	UserID      string     `dorm:"type:uuid;not null" json:"user_id"`
	StartDate   time.Time  `dorm:"type:date;not null" json:"start_date"`
	EndDate     *time.Time `dorm:"type:date" json:"end_date,omitempty"`
	CreatedAt   time.Time  `dorm:"type:autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `dorm:"type:autoUpdateTime" json:"updated_at"`
}
