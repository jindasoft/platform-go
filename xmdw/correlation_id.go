package xmdw

import (
	"context"

	"github.com/google/uuid"
	"github.com/jindasoft/jinda-platform/xconst"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func CorrelationIDMiddleware() echo.MiddlewareFunc {
	mdw := middleware.RequestIDConfig{
		RequestIDHandler: func(c *echo.Context, id string) {
			correlationID := c.Request().Header.Get(xconst.HeaderXCorrelationID)
			if correlationID == "" {
				correlationID = uuid.New().String()
			}

			// set to context
			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, xconst.ContextCorrelationID, correlationID)
			c.SetRequest(c.Request().WithContext(ctx))

			// set to header
			c.Response().Header().Set(xconst.HeaderXCorrelationID, correlationID)
		},
		TargetHeader: echo.HeaderXCorrelationID,
	}

	return middleware.RequestIDWithConfig(mdw)
}
