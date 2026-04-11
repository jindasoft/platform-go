package xmdw

import (
	"context"

	"github.com/jindasoft/jinda-platforms/xconst"
	"github.com/jindasoft/jinda-platforms/xenums"
	"github.com/labstack/echo/v5"
)

func RequestCultureMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			culture := xenums.CultureDefault

			lang := c.QueryParam("lang")
			if lang != "" {
				culture = ToCultureCode(lang)
			} else {
				acceptLang := c.Request().Header.Get(xconst.HeaderAcceptLanguage)
				if acceptLang != "" {
					culture = ToCultureCode(acceptLang)
				}
			}

			// set to context
			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, xconst.ContextCultureCode, culture)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

func ToCultureCode(cultureCode string) xenums.Culture {
	culture, err := xenums.PairCulture(cultureCode)
	if err != nil {
		return xenums.CultureDefault
	}

	return culture
}
