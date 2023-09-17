package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gost1k337/url_shortener/api_gateway_service/config"
	_ "github.com/gost1k337/url_shortener/api_gateway_service/docs"
	"github.com/gost1k337/url_shortener/api_gateway_service/internal/handlers/middlewares"
	"github.com/gost1k337/url_shortener/api_gateway_service/internal/service"
	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
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
		httpSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", net.JoinHostPort(cfg.App.Host, cfg.App.Port))),
	))
	r.Post("/short", h.CreateURLShort)
	r.Get("/u/{token}", h.Redirect)
	r.Post("/users", h.CreateUser)
	r.Get("/users/{id}", h.GetUser)
	r.Delete("/users/{id}", h.DeleteUser)

	h.http = r

	return h
}

func (h *Handler) HTTP() http.Handler {
	return h.http
}

// CreateURLShort godoc
//
// @Summary Create short url
// @Description Create short url from url and return it
// @ID create-short-url
// @Tags ShortURL
// @Success 201
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /short [post].
func (h *Handler) CreateURLShort(w http.ResponseWriter, r *http.Request) {
	var input CreateURLShortInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		h.logger.Error(err.Error())

		return
	}

	input.ExpireAt *= time.Hour

	res, err := h.services.URLShort.Create(r.Context(), input.OriginalURL, input.ExpireAt)
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
// @Tags ShortURL
// @Success 303
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /u/{token} [get].
func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	h.logger.Info(token)

	res, err := h.services.URLShort.Get(r.Context(), token)
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

	http.Redirect(w, r, res.OriginalURL, http.StatusSeeOther)
}

// CreateUser godoc
//
// @Summary Create user
// @Description Create user
// @ID create-user
// @Tags User
// @Success 201
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /users [post].
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input CreateUserInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		h.logger.Error(err.Error())

		return
	}

	res, err := h.services.User.Create(r.Context(), input.Username, input.Email, input.Password)
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

// GetUser godoc
//
// @Summary Get user
// @Description Get user
// @ID get-user
// @Tags User
// @Success 200
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /users/{id} [post].
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	p := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		h.logger.Error(err.Error())

		return
	}

	res, err := h.services.User.Get(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			h.logger.Error(err.Error())

			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		h.logger.Error(err.Error())

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		h.logger.Error(err.Error())

		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteUser godoc
//
// @Summary Delete user
// @Description Delete user
// @ID delete-user
// @Tags User
// @Success 204
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /users/{id} [delete].
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	p := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		h.logger.Error(err.Error())

		return
	}

	res, err := h.services.User.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			h.logger.Error(err.Error())

			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		h.logger.Error(err.Error())

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		h.logger.Error(err.Error())

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
