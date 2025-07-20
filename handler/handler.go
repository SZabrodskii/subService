package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"subService/model"
	"subService/repository"
	"subService/service"
	"time"
)

func Provide() fx.Option {
	return fx.Provide(NewSubscriptionHandler)
}

type SubscriptionHandler struct {
	service service.SubscriptionServiceInterface
}

func NewSubscriptionHandler(service service.SubscriptionServiceInterface) *SubscriptionHandler {
	return &SubscriptionHandler{service: service}
}

func (h *SubscriptionHandler) Register(r *gin.Engine) {
	grp := r.Group("/api/v1/subscriptions")
	grp.POST("", h.Create)
	grp.GET("/:id", h.GetByID)
	grp.PUT("/:id", h.Update)
	grp.DELETE("/:id", h.Delete)
	grp.GET("", h.GetAll)

	r.GET("/api/v1/summary", h.GetSum)
}

// Create создаёт новую подписку
// @Summary Create subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body  service.CreateSubscriptionRequest true "Subscription to create"
// @Success 201 {object} model.SubscriptionResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/v1/subscriptions [post]
func (h *SubscriptionHandler) Create(c *gin.Context) {
	var req service.CreateSubscriptionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := sub.ToResponse()
	c.JSON(http.StatusCreated, resp)
}

// GetByID возвращает подписку по ID
// @Summary Get subscription by ID
// @Param id path string true "Subscription ID"
// @Success 200 {object} model.SubscriptionResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /subscriptions/{id} [get]
func (h *SubscriptionHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	sub, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if sub == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}
	c.JSON(http.StatusOK, sub.ToResponse())
}

// Update изменяет существующую подписку
// @Tags subscriptions
// @Accept json
// @Produce json
// @Summary Update subscription
// @Param id path string true "Subscription ID"
// @Param subscription body service.UpdateSubscriptionRequest true "Fields to update"
// @Success 200 {object} model.SubscriptionResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/v1/subscriptions/{id} [put]
func (h *SubscriptionHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req service.UpdateSubscriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sub, err := h.service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if sub == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}
	c.JSON(http.StatusOK, sub.ToResponse())
}

// Delete удаляет подписку по ID
// @Summary Delete subscription
// @Param id path string true "Subscription ID"
// @Success 204 {object} nil
// @Failure 500 {object} model.ErrorResponse
// @Router /subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "subscription deleted"})
}

// GetAll возвращает список подписок по фильтру
// @Summary List subscriptions
// @Param user_id query string false "Filter by user ID"
// @Param service_name query string false "Filter by service name"
// @Param from query string false "Start period YYYY-MM"
// @Param to query string false "End period YYYY-MM"
// @Success 200 {array} model.SubscriptionResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /subscriptions [get]
func (h *SubscriptionHandler) GetAll(c *gin.Context) {
	var filter repository.SubscriptionFilter
	filter.UserID = c.Query("user_id")
	filter.ServiceName = c.Query("service_name")
	if from := c.Query("from"); from != "" {
		if t, err := time.Parse("2006-01", from); err == nil {
			filter.From = t
		}
	}
	if to := c.Query("to"); to != "" {
		if t, err := time.Parse("2006-01", to); err == nil {
			filter.To = t
		}
	}
	subs, err := h.service.GetAll(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]model.SubscriptionResponse, len(subs))
	for i, sub := range subs {
		resp[i] = sub.ToResponse()
	}
	c.JSON(http.StatusOK, resp)

}

// GetSum возвращает суммарную стоимость подписок
// @Summary Summary of subscriptions
// @Param user_id query string false "Filter by user ID"
// @Param service_name query string false "Filter by service name"
// @Param from query string false "Start period YYYY-MM"
// @Param to query string false "End period YYYY-MM"
// @Success 200 {object} map[string]int
// @Failure 500 {object} model.ErrorResponse
// @Router /summary [get]
func (h *SubscriptionHandler) GetSum(c *gin.Context) {
	var filter repository.SumFilter
	filter.UserID = c.Query("user_id")
	filter.ServiceName = c.Query("service_name")
	if from := c.Query("from"); from != "" {
		if t, err := time.Parse("2006-01", from); err == nil {
			filter.From = t
		}
	}
	if to := c.Query("to"); to != "" {
		if t, err := time.Parse("2006-01", to); err == nil {
			filter.To = t
		}
	}
	sum, err := h.service.GetSum(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sum": sum})
}
