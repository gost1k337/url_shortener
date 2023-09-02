package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gost1k337/url_shortener/api_gateway_service/internal/service"
	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
	"github.com/rs/cors"
	"net/http"
)

type Handler struct {
	services *service.Services
	http     *chi.Mux
	logger   logging.Logger
}

func New(services *service.Services, logger logging.Logger) *Handler {
	h := &Handler{
		services: services,
		logger:   logger,
	}

	corsCfg := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodOptions,
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders: []string{"Accept", "Content-Type", "Accept-Encoding"},
	})

	r := chi.NewRouter()
	r.Use(corsCfg.Handler)
	r.Use(middleware.DefaultLogger)

	h.http = r
	return h
}

func (h *Handler) HTTP() http.Handler {
	return h.http
}
