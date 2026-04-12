package xmdw

import (
	"context"

	"github.com/jindasoft/jinda-platform/xconst"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func TraceIDMiddleware() echo.MiddlewareFunc {
	mdw := middleware.RequestIDConfig{
		RequestIDHandler: func(c *echo.Context, requestID string) {
			traceID := c.Request().Header.Get(xconst.HeaderXTraceID)

			// set to context
			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, xconst.ContextTraceID, traceID)
			c.SetRequest(c.Request().WithContext(ctx))

			// set to header
			c.Response().Header().Set(xconst.HeaderXTraceID, traceID)
		},
		TargetHeader: xconst.HeaderXTraceID,
	}

	return middleware.RequestIDWithConfig(mdw)
}
