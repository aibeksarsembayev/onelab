package httpdelivery

import (
	"context"
	"net/http"
	"strconv"

	"github.com/aibeksarsembayev/onelab/tasks/lab4/domain"
	"github.com/labstack/echo/v4"
)

// ResponceError
type ResponceError struct {
	Message string `json:"message"`
}

// Create user ...
func (uh *UserHandler) Create(c echo.Context) error {
	user := new(domain.UserRequestDTO)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, user)
}

// Get user by ID ...
func (uh *UserHandler) GetByID(c echo.Context) error {
	// get user id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// get user from db
	user, err := uh.UserUsecase.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusFound, user)
}

// Get all users ...
func (uh *UserHandler) GetAll(c echo.Context) error {
	// getall users with status filter
	users, err := uh.UserUsecase.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

// Update user by id ...
func (uh *UserHandler) Update(c echo.Context) error {
	// get id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// user info bind from form
	user := new(domain.UserRequestDTO)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// update user in db
	us := domain.User{
		ID:      id,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
	}
	err = uh.UserUsecase.Update(context.Background(), &us)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

// Delete user by id ...
func (uh *UserHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = uh.UserUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}
