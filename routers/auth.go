package routers

import (
	"system-monitor-backend/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// BuildAuthRouter -- builds and returns an auth router
func BuildAuthRouter(userStore models.UserStore) *chi.Mux {
	apiRouter := chi.NewRouter()
	apiRouter.Use(middleware.SetHeader("Content-Type", "application/json"))

	return apiRouter
}
