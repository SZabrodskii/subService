package model

import "time"

type Subscription struct {
	ID          string     `gorm:"type:uuid;default:gen_random_uuid();primary_key" json:"id"`
	ServiceName string     `dorm:"not null" json:"service_name"`
	Price       int        `gorm:"not null" json:"price"`
	UserID      string     `gorm:"type:uuid;not null" json:"user_id"`
	StartDate   time.Time  `gorm:"type:date;not null" json:"start_date"`
	EndDate     *time.Time `gorm:"type:date" json:"end_date,omitempty"`
	CreatedAt   time.Time  `gorm:"type:autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:autoUpdateTime" json:"updated_at"`
}

type SubscriptionResponse struct {
	ID          string  `json:"id"`
	ServiceName string  `json:"service_name"`
	Price       int     `json:"price"`
	UserID      string  `json:"user_id"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date,omitempty"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func (s *Subscription) ToResponse() SubscriptionResponse {
	start := s.StartDate.Format("2006-01")
	var end *string
	if s.EndDate != nil {
		e := s.EndDate.Format("2006-01")
		end = &e
	}
	return SubscriptionResponse{
		ID:          s.ID,
		ServiceName: s.ServiceName,
		Price:       s.Price,
		UserID:      s.UserID,
		StartDate:   start,
		EndDate:     end,
		CreatedAt:   s.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   s.UpdatedAt.Format(time.RFC3339),
	}
}
