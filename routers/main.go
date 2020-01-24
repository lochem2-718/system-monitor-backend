package routers

import (
	"fmt"
	"net/http"
	"system-monitor-backend/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Run - Starts server's router
func Run(userStore *models.UserStore, port int) {
	corsPolicy := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://monitor.weinberger.systems"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	mainRouter := chi.NewRouter()
	mainRouter.Use(corsPolicy.Handler)
	mainRouter.Mount("/auth", BuildAuthRouter(userStore))

	http.ListenAndServe(fmt.Sprintf(":%d", port), mainRouter)
}
