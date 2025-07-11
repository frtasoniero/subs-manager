package handlers

import (
	"net/http"

	"github.com/frtasoniero/subsmanager/internal/usecases"
	"github.com/frtasoniero/subsmanager/pkg/utils"
	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	subscriptionUseCase *usecases.SubscriptionUseCase
}

func NewSubscriptionHandler(subscriptionUseCase *usecases.SubscriptionUseCase) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionUseCase: subscriptionUseCase,
	}
}

func (h *SubscriptionHandler) GetAllSubscriptions(c *gin.Context) {
	subscriptions, err := h.subscriptionUseCase.GetAllSubscriptions(c.Request.Context())
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get subscriptions", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Subscriptions retrieved successfully", subscriptions)
}
