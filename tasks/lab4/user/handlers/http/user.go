package httpdelivery

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (u *UserHandler) Create(c echo.Context) error {
	return nil
}

func (u *UserHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	mes := fmt.Sprintf("User - %v is here\n", id)
	return c.String(http.StatusFound, mes)
}

func (u *UserHandler) GetAll(c echo.Context) error {

	return nil
}

func (u *UserHandler) Update(c echo.Context) error {
	id := c.Param("id")
	mes := fmt.Sprintf("User - %v has been updated\n", id)
	return c.String(http.StatusFound, mes)
}

func (u *UserHandler) Delete(c echo.Context) error {
	return nil
}
