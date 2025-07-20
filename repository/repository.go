package repository

import (
	"subService/model"
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

type SubscriptionRepository interface {
	Create(sub *model.Subscription) error
	GetByID(id string) (*model.Subscription, error)
	Update(sub *model.Subscription) error
	Delete(id string) error
	GetAll(filter SubscriptionFilter) ([]*model.Subscription, error)
	GetSum(filter SumFilter) (int, error)
}
