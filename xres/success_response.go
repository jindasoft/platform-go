package xres

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

const SuccessfulMessage = "Successful."

// 200 OK
type SuccessResponse struct {
	Success bool   `json:"success" example:"true"`
	Type    string `json:"type" example:"success"`
	Message string `json:"message" example:"Successful."`
	Data    any    `json:"data,omitempty"`
}

func Success[T any](c *echo.Context, data T) error {
	r := SuccessResponse{
		Success: true,
		Type:    "success",
		Message: SuccessfulMessage,
		Data:    data,
	}

	return c.JSON(http.StatusOK, r)
}

type PagingMeta struct {
	Offset int64 `json:"offset" example:"0"`
	Limit  int64 `json:"limit" example:"10"`
	Total  int64 `json:"total" example:"100"`
}

type PagingResponse struct {
	Success bool       `json:"success" example:"true"`
	Type    string     `json:"type" example:"success"`
	Message string     `json:"message" example:"Successful."`
	Meta    PagingMeta `json:"meta"`
	Data    any        `json:"data,omitempty"`
}

func Paging[T any](c *echo.Context, data T, meta PagingMeta) error {
	r := PagingResponse{
		Success: true,
		Type:    "success",
		Message: SuccessfulMessage,
		Data:    data,
		Meta:    meta,
	}

	return c.JSON(http.StatusOK, r)
}

type DynamoPagingMeta struct {
	Limit             int64  `json:"limit" example:"10"`
	ExclusiveStartKey string `json:"exclusive_start_key,omitempty" example:"eyJtZW1iZXJfaWQiOiAiMTIzNDU2IiwgInRpbWVzdGFtcCI6IDE2ODk0ODAwMDAwMDAwfQ=="`
	HasMore           bool   `json:"has_more" example:"true"`
}

type DynamoPagingResponse struct {
	Success bool             `json:"success" example:"true"`
	Type    string           `json:"type" example:"success"`
	Message string           `json:"message" example:"Successful."`
	Meta    DynamoPagingMeta `json:"meta"`
	Data    any              `json:"data,omitempty"`
}

func DynamoPaging[T any](c *echo.Context, data T, meta DynamoPagingMeta) error {
	r := DynamoPagingResponse{
		Success: true,
		Type:    "success",
		Message: SuccessfulMessage,
		Data:    data,
		Meta:    meta,
	}

	return c.JSON(http.StatusOK, r)
}

// 201 Created
type CreatedResponse struct {
	Success bool   `json:"success" example:"true"`
	Type    string `json:"type" example:"created"`
	Message string `json:"message" example:"Created."`
	Data    any    `json:"data,omitempty"`
}

func Created[T any](c *echo.Context, data T) error {
	r := CreatedResponse{
		Success: true,
		Type:    "created",
		Message: "Created.",
		Data:    data,
	}

	return c.JSON(http.StatusCreated, r)
}

// 202 Accepted
type AcceptedResponse struct {
	Success bool   `json:"success" example:"true"`
	Type    string `json:"type" example:"accepted"`
	Message string `json:"message" example:"Accepted."`
	Data    any    `json:"data,omitempty"`
}

func Accepted[T any](c *echo.Context, data T) error {
	r := AcceptedResponse{
		Success: true,
		Type:    "accepted",
		Message: "Accepted.",
		Data:    data,
	}

	return c.JSON(http.StatusAccepted, r)
}

// 204 No Content
func NoContent(c *echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
