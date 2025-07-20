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

func (h *SubscriptionHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "subscription deleted"})
}

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
