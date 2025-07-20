package service

import (
	"fmt"
	"go.uber.org/fx"
	"subService/model"
	"subService/repository"
	"time"
)

func Provide() fx.Option {
	return fx.Provide(NewSubscriptionService)
}

type subscriptionService struct {
	repo repository.SubscriptionRepository
}

type SubscriptionServiceInterface interface {
	Create(req CreateSubscriptionRequest) (*model.Subscription, error)
	GetByID(id string) (*model.Subscription, error)
	Update(id string, req UpdateSubscriptionRequest) (*model.Subscription, error)
	Delete(id string) error
	GetAll(filter repository.SubscriptionFilter) ([]*model.Subscription, error)
	GetSum(filter repository.SumFilter) (int, error)
}

func NewSubscriptionService(repo repository.SubscriptionRepository) SubscriptionServiceInterface {
	return &subscriptionService{repo: repo}
}

func (s *subscriptionService) Create(req CreateSubscriptionRequest) (*model.Subscription, error) {
	startDate, err := time.Parse("2006-01", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %w", err)
	}
	var endDate *time.Time
	if req.EndDate != nil && *req.EndDate != "" {
		ed, err := time.Parse("2006-01", *req.EndDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse end date: %w", err)
		}
		endDate = &ed
	}

	sub := &model.Subscription{
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      req.UserID,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	if err := s.repo.Create(sub); err != nil {
		return nil, fmt.Errorf("failed to create subscription: %w", err)
	}
	return sub, nil

}

func (s *subscriptionService) GetByID(id string) (*model.Subscription, error) {
	return s.repo.GetByID(id)
}

func (s *subscriptionService) Update(id string, req UpdateSubscriptionRequest) (*model.Subscription, error) {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription: %w", err)
	}

	if existing == nil {
		return nil, fmt.Errorf("subscription not found")
	}

	if req.ServiceName != nil {
		existing.ServiceName = *req.ServiceName
	}
	if req.Price != nil {
		existing.Price = *req.Price
	}
	if req.EndDate != nil {
		if *req.EndDate == "" {
			existing.EndDate = nil
		} else {
			ed, err := time.Parse("2006-01", *req.EndDate)
			if err != nil {
				return nil, fmt.Errorf("failed to parse end date: %w", err)
			}
			existing.EndDate = &ed
		}
	}
	if err := s.repo.Update(existing); err != nil {
		return nil, fmt.Errorf("failed to update subscription: %w", err)
	}

	return existing, nil
}

func (s *subscriptionService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *subscriptionService) GetAll(filter repository.SubscriptionFilter) ([]*model.Subscription, error) {
	return s.repo.GetAll(filter)
}

func (s *subscriptionService) GetSum(filter repository.SumFilter) (int, error) {
	return s.repo.GetSum(filter)
}
