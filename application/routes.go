package application

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"

	"github.com/bagussans/ms-support-golang/handler"
)

func loadRoutes(db *bun.DB) *chi.Mux {
	// Load .env file (ignore if missing)
	if os.Getenv("APP_ENV") != "prod" {
		_ = godotenv.Load()
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// --- CORS setup ---
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"https://tools.bagussan.my.id",
			"http://localhost:3333",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Token", "X-API-KEY"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // 5 minutes
	}))
	// -------------------
	// router.Route("/authext", loadAuthExtRoutes)

	router.Route("/tools", func(r chi.Router) {
		r.Use(handler.AuthCheckApiKey())
		loadToolsRoutes(r)
	})

	return router
}

func loadToolsRoutes(router chi.Router) {
	toolsHandler := &handler.Tools{}

	router.Post("/calc-student-avg-score", toolsHandler.CalcStudentAvgScore)
	router.Post("/image-editor", toolsHandler.ImageEditor)
}
