package service

type CreateSubscriptionRequest struct {
	ServiceName string  `json:"service_name" binding:"required"`
	Price       int     `json:"price" binding:"required,min=0"`
	UserID      string  `json:"user_id" binding:"required,uuid"`
	StartDate   string  `json:"start_date" binding:"required,datetime=2006-01"`
	EndDate     *string `json:"end_date,omitempty" binding:"omitempty,datetime=2006-01"`
}

type UpdateSubscriptionRequest struct {
	ServiceName *string `json:"service_name,omitempty" binding:"omitempty,required_with=Price"`
	Price       *int    `json:"price,omitempty" binding:"omitempty,min=0"`
	EndDate     *string `json:"end_date" binding:"omitempty,datetime=2006-01"`
}
