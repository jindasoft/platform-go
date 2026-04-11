package xrequester

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jindasoft/jinda-platforms/xconst"
	"github.com/jindasoft/jinda-platforms/xutils"
)

type Service struct {
	client *resty.Client
}

func New(ctx context.Context, debug *bool) *Service {
	client := resty.New()
	client.SetDebug(*debug)
	client.SetTimeout(30 * time.Second)

	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		// Set default headers
		req.Header.Set(xconst.HeaderXCorrelationID, xutils.GetCorrelationIDString(req.Context()))
		req.Header.Set(xconst.HeaderXTraceID, xutils.GetTraceIDString(req.Context()))
		req.Header.Set(xconst.HeaderAcceptLanguage, xutils.GetCultureString(req.Context()))

		return nil
	})

	return &Service{
		client: client,
	}
}

// setup base url
func (s *Service) SetBaseURL(baseURL string) *resty.Client {
	s.client.SetBaseURL(baseURL)

	return s.client
}
