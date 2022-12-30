package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrBadRequest = echo.NewHTTPError(http.StatusBadRequest, "bad request body")
)
