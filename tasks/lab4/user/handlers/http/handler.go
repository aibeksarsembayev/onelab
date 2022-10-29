package httpdelivery

import (
	"github.com/aibeksarsembayev/onelab/tasks/lab4/domain"
	"github.com/labstack/echo/v4"
)

// UseHandler represent the httphandler for user
type UserHandler struct {
	UserUsecase domain.UserUsecase
}

// NewUserHandler ...
func NewUserHandler(e *echo.Echo, us domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}
	user := e.Group("/user")
	user.POST("/", handler.Create)
	user.GET("/", handler.GetByID)
	user.GET("/all", handler.GetAll)
	user.PUT("/:id", handler.Update)
	user.DELETE("/:id", handler.Delete)
}
