package xmdw

import (
	"context"

	"github.com/jindasoft/jinda-platform/xconst"
	"github.com/jindasoft/jinda-platform/xenums"
	"github.com/labstack/echo/v5"
)

func RequestLocaleMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			locale := xenums.LocaleDefault

			lang := c.QueryParam("lang")
			if lang != "" {
				locale = ToLocaleCode(lang)
			} else {
				acceptLang := c.Request().Header.Get(xconst.HeaderAcceptLanguage)
				if acceptLang != "" {
					locale = ToLocaleCode(acceptLang)
				}
			}

			// set to context
			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, xconst.ContextLocaleCode, locale)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

func ToLocaleCode(localeCode string) xenums.Locale {
	locale, err := xenums.FromLocaleCode(localeCode)
	if err != nil {
		return xenums.LocaleDefault
	}

	return locale
}
