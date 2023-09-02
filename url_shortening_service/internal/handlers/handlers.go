package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/handlers/middlewares"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/service"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
	"github.com/rs/cors"
	"net/http"
	"time"
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
	r.Use(middlewares.CommonMiddleware)

	r.Post("/short", h.Create)

	h.http = r
	return h
}

func (h *Handler) HTTP() http.Handler {
	return h.http
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input CreateShortURLInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	input.ExpireAt = input.ExpireAt * time.Hour

	err := h.services.ShortUrl.Create(r.Context(), 1, input.OriginalURL, input.ExpireAt)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.WriteHeader(201)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

}
