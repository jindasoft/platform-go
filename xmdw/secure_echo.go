package xmdw

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func SecureMiddleware() echo.MiddlewareFunc {
	secureConfig := middleware.SecureConfig{
		ContentTypeNosniff: "nosniff",
		ReferrerPolicy:     "same-origin",
	}

	return middleware.SecureWithConfig(secureConfig)
}
