package handler

import "github.com/labstack/echo/v4"

type (
	IMobile interface {
		PutUserCall(c echo.Context) error
		GetUserBilling(c echo.Context) error
	}
)
