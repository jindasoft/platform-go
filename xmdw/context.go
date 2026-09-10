package xmdw

import (
	"context"

	"github.com/jindasoft/template-platform-go/xconst"
	"github.com/labstack/echo/v5"
)

func ContextMiddleware(service, env string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, xconst.ContextServiceName, service)
			ctx = context.WithValue(ctx, xconst.ContextEnvironment, env)

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
