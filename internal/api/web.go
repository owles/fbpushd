package api

import (
	"fbpushd/internal/push"
	"firebase.google.com/go"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
	"net/http"
)

type PushRequest struct {
	Token string `json:"token" binding:"required"`
	Push  struct {
		Data  map[string]string `json:"data"`
		Title string            `json:"title"`
		Body  string            `json:"body"`
	} `json:"push" binding:"required"`
}

type PushResponse struct {
	Token   string `json:"token"`
	Success bool   `json:"success"`
}

type ErrorResponse struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

type API struct {
	App *firebase.App
}

func NewAPI(app *firebase.App) *API {
	return &API{app}
}

// RegisterRoutes sets up the API routes
// @Summary Send push notifications
// @Description Send push notifications to specific tokens using Firebase Cloud Messaging.
// @Tags push
// @Accept  json
// @Produce  json
// @Param  requests body []PushRequest true "Push Request"
// @Success 200 {array} PushResponse
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Router /push [post]
func (a *API) controllerPush(c *gin.Context) {
	var requests []PushRequest
	if err := c.ShouldBindJSON(&requests); err != nil {
		slog.Error("Invalid request", "message", err)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Result:  "error",
			Message: "Invalid request",
		})
		return
	}

	var responses []PushResponse
	for _, req := range requests {
		success := push.SendNotification(req.Token, req.Push.Title, req.Push.Body, req.Push.Data, a.App)
		responses = append(responses, PushResponse{
			Token:   req.Token,
			Success: success,
		})
	}

	c.JSON(http.StatusOK, responses)
}

func (a *API) RegisterRoutes(r *gin.Engine, app *firebase.App) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/push", a.controllerPush)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"result":  "error",
			"message": "invalid API method",
		})
	})
}
