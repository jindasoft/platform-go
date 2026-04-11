package xmdw

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLogger()
}
