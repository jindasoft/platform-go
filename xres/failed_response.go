package xres

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

// 401 Unauthorized
type UnauthorizedResponse struct {
	Success bool   `json:"success" example:"false"`
	Type    string `json:"type" example:"unauthorized"`
	Message string `json:"message" example:"Unauthorized."`
}

func Unauthorized(c *echo.Context) error {
	r := UnauthorizedResponse{
		Success: false,
		Type:    "unauthorized",
		Message: "Unauthorized.",
	}

	return c.JSON(http.StatusUnauthorized, r)
}

// 403 Forbidden
type ForbiddenResponse struct {
	Success bool   `json:"success" example:"false"`
	Type    string `json:"type" example:"forbidden"`
	Message string `json:"message" example:"Forbidden."`
}

func Forbidden(c *echo.Context) error {
	r := ForbiddenResponse{
		Success: false,
		Type:    "forbidden",
		Message: "Forbidden.",
	}

	return c.JSON(http.StatusForbidden, r)
}

// 404 Not Found
type ResourceNotFoundResponse struct {
	Success bool   `json:"success" example:"false"`
	Type    string `json:"type" example:"resource_not_found"`
	Message string `json:"message" example:"Resource not found."`
}

func ResourceNotFound(c *echo.Context, msg string) error {
	r := ResourceNotFoundResponse{
		Success: false,
		Type:    "resource_not_found",
		Message: msg,
	}

	return c.JSON(http.StatusNotFound, r)
}

// 400 Bad Request
type BadRequestResponse struct {
	Success bool   `json:"success" example:"false"`
	Type    string `json:"type" example:"bad_request"`
	Message string `json:"message" example:"Bad request."`
}

func badRequest(typeStr string, msg string) (res BadRequestResponse) {
	return BadRequestResponse{
		Success: false,
		Type:    typeStr,
		Message: msg,
	}
}

func BadRequestBindData(c *echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, badRequest("bind_data_error", msg))
}

func BadRequestValidation(c *echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, badRequest("validation_error", msg))
}

func BadRequestSpecific(c *echo.Context, typeStr string, msg string) error {
	return c.JSON(http.StatusBadRequest, badRequest(typeStr, msg))
}

// 409 Conflict
type ConflictResponse struct {
	Success bool   `json:"success" example:"false"`
	Type    string `json:"type" example:"conflict"`
	Message string `json:"message" example:"Conflict."`
}

func Conflict(c *echo.Context, msg string) error {
	r := ConflictResponse{
		Success: false,
		Type:    "conflict",
		Message: msg,
	}

	return c.JSON(http.StatusConflict, r)
}

// 422 Unprocessable Entity
type UnprocessableEntityResponse struct {
	Success bool   `json:"success" example:"false"`
	Type    string `json:"type" example:"unprocessable_entity"`
	Message string `json:"message" example:"Unprocessable entity."`
}

func unprocessable(typeStr string, msg string) (res UnprocessableEntityResponse) {
	return UnprocessableEntityResponse{
		Success: false,
		Type:    typeStr,
		Message: msg,
	}
}

func UnprocessableEntity(c *echo.Context, msg string) error {
	return c.JSON(http.StatusUnprocessableEntity, unprocessable("unprocessable_entity", msg))
}

func UnprocessableSpecific(c *echo.Context, typeStr string, msg string) error {
	return c.JSON(http.StatusUnprocessableEntity, unprocessable(typeStr, msg))
}

// 429 Many Requests
type TooManyRequestsResponse struct {
	Success bool   `json:"success" example:"false"`
	Type    string `json:"type" example:"too_many_requests"`
	Message string `json:"message" example:"Too many requests."`
}

func TooManyRequests(c *echo.Context, msg string) error {
	r := TooManyRequestsResponse{
		Success: false,
		Type:    "too_many_requests",
		Message: msg,
	}

	return c.JSON(http.StatusTooManyRequests, r)
}
