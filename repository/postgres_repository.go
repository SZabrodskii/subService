package repository

import (
	"errors"
	"fmt"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"subService/model"
)

func Provide() fx.Option {
	return fx.Provide(NewSubscriptionRepository)

}

type PostgresSubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &PostgresSubscriptionRepository{
		db: db,
	}
}

func (r *PostgresSubscriptionRepository) Create(sub *model.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *PostgresSubscriptionRepository) GetByID(id string) (*model.Subscription, error) {
	var sub model.Subscription
	err := r.db.First(&sub, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription: %w", err)
	}

	return &sub, nil
}

func (r *PostgresSubscriptionRepository) Update(sub *model.Subscription) error {
	return r.db.Save(sub).Error
}

func (r *PostgresSubscriptionRepository) Delete(id string) error {
	return r.db.Delete(&model.Subscription{}, "id = ?", id).Error
}

func (r *PostgresSubscriptionRepository) GetAll(filter SubscriptionFilter) ([]*model.Subscription, error) {
	var subs []*model.Subscription
	query := r.db.Model(&model.Subscription{})

	if filter.UserID != "" {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.ServiceName != "" {
		query = query.Where("service_name = ?", filter.ServiceName)
	}
	if !filter.From.IsZero() {
		query = query.Where("start_date >= ?", filter.From)
	}
	if !filter.To.IsZero() {
		query = query.Where("start_date <= ?", filter.To)
	}
	err := query.Find(&subs).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get subscriptions: %w", err)
	}
	return subs, nil
}

func (r *PostgresSubscriptionRepository) GetSum(filter SumFilter) (int, error) {
	var sum int64
	query := r.db.Model(&model.Subscription{}).Select("COALESCE(SUM(price), 0)")
	if filter.UserID != "" {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.ServiceName != "" {
		query = query.Where("service_name = ?", filter.ServiceName)
	}
	if !filter.From.IsZero() {
		query = query.Where("start_date >= ?", filter.From)
	}
	if !filter.To.IsZero() {
		query = query.Where("start_date <= ?", filter.To)
	}

	err := query.Scan(&sum).Error
	if err != nil {
		return 0, fmt.Errorf("failed to get sum: %w", err)
	}
	return int(sum), nil
}
