package xmdw

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jindasoft/template-platform-go/xconst"
	"github.com/labstack/echo/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func JwtExtractMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			if c == nil {
				return next(c)
			}

			auth, err := extractToken(c)
			if err != nil || auth == "" {
				return next(c)
			}

			claims := jwt.MapClaims{}
			_, _, err = new(jwt.Parser).ParseUnverified(auth, claims)
			if err != nil {
				return next(c)
			}

			// Set the claims in the context for use in handlers
			sub, ok := claims["sub"].(string)
			if !ok {
				return next(c)
			}

			accountOID, err := primitive.ObjectIDFromHex(sub)
			if err != nil {
				return next(c)
			}
			accountUUID, err := objectIDToUUID(accountOID)
			if err != nil {
				return next(c)
			}

			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, xconst.ContextAccountOID, accountOID)
			ctx = context.WithValue(ctx, xconst.ContextAccountUUID, accountUUID)

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

func extractToken(c *echo.Context) (string, error) {
	header := c.Request().Header.Get(echo.HeaderAuthorization)

	header = strings.TrimSpace(header)
	if header == "" {
		return "", fmt.Errorf("extractToken: %s", "empty authorization header")
	}

	header = strings.TrimPrefix(header, "Bearer ")
	if header == "" {
		return "", fmt.Errorf("extractToken: %s", "empty token")
	}

	return header, nil
}

func objectIDToUUID(oid primitive.ObjectID) (uuid.UUID, error) {
	objIDBytes := oid[:]

	padded := make([]byte, 16)
	copy(padded[:12], objIDBytes)

	uid, err := uuid.FromBytes(padded)
	if err != nil {
		return uuid.Nil, err
	}

	return uid, nil

}
