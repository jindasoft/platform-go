package xmdw

import (
	"context"

	"github.com/jindasoft/jinda-platforms/xconst"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func SpanIDMiddleware() echo.MiddlewareFunc {
	mdw := middleware.RequestIDConfig{
		RequestIDHandler: func(c *echo.Context, requestID string) {
			spanID := c.Request().Header.Get(xconst.HeaderXSpanID)

			// set to context
			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, xconst.ContextSpanID, spanID)
			c.SetRequest(c.Request().WithContext(ctx))

			// set to header
			c.Response().Header().Set(xconst.HeaderXSpanID, spanID)
		},
		TargetHeader: xconst.HeaderXSpanID,
	}

	return middleware.RequestIDWithConfig(mdw)
}
