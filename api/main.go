package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
)

type HTTPServer interface {
	Start() error
}

type HttpServerImpl struct {
	httpServerAddress *url.URL
}

func (s *HttpServerImpl) Start() error {
	e := echo.New()
	e.Use(CorsMiddleware())
	e.Validator = &Validator{validator: validator.New()}

	publicApiGroup := e.Group("/api")
	publicApiGroup.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Test Connection")
	})

	return e.Start(s.httpServerAddress.Host)
}

func NewHttpServer(httpServerAddress *url.URL) *HttpServerImpl {
	return &HttpServerImpl{
		httpServerAddress: httpServerAddress,
	}
}
