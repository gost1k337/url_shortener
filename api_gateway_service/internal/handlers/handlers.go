package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gost1k337/url_shortener/api_gateway_service/config"
	_ "github.com/gost1k337/url_shortener/api_gateway_service/docs"
	"github.com/gost1k337/url_shortener/api_gateway_service/internal/handlers/middlewares"
	"github.com/gost1k337/url_shortener/api_gateway_service/internal/service"
	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

type Handler struct {
	services *service.Services
	http     *chi.Mux
	logger   logging.Logger
}

func New(services *service.Services, logger logging.Logger, cfg *config.Config) *Handler {
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
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", cfg.App.Host, cfg.App.Port)),
	))
	r.Post("/short", h.CreateUrlShort)
	r.Get("/u/{token}", h.Redirect)

	h.http = r
	return h
}

func (h *Handler) HTTP() http.Handler {
	return h.http
}

// CreateUrlShort godoc
//
// @Summary Create short url
// @Description Create short url from url and return it
// @ID create-short-url
// @Tags ShortUrl
// @Success 201
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /short [post]
func (h *Handler) CreateUrlShort(w http.ResponseWriter, r *http.Request) {
	var input CreateUrlShortInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		h.logger.Error(err.Error())
		return
	}
	input.ExpireAt = input.ExpireAt * time.Hour

	res, err := h.services.UrlShort.Create(r.Context(), input.OriginalUrl, input.ExpireAt)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		h.logger.Error(err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		h.logger.Error(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Redirect godoc
//
// @Summary Redirect to original url
// @Description Redirect user from short url to original url
// @ID redirect-short-url
// @Tags ShortUrl
// @Success 303
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /u/{token} [get]
func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	h.logger.Info(token)
	res, err := h.services.UrlShort.Get(r.Context(), token)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		h.logger.Error(err.Error())
		return
	}
	if res == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		h.logger.Error(err.Error())
		return
	}
	http.Redirect(w, r, res.OriginalUrl, http.StatusSeeOther)
}
