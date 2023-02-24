package infra

import (
	"app/ent"
	"app/interfaces/controller"
	"encoding/json"
	"net/http"
	"time"

	"github.com/chihiros/logger"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(conn *ent.Client) *chi.Mux {
	r := chi.NewRouter()
	r.Use(logger.Logger)
	r.Use(middleware.Recoverer)

	// Access-Control-Allow-Originを許可する
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	controller := controller.NewController(conn)
	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", controller.Get)
			r.Get("/query", controller.GetByID)
			r.Post("/", controller.Post)
			r.Delete("/", controller.Delete)
		})

		r.Route("/now", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				jst, err := time.LoadLocation("Asia/Tokyo")
				if err != nil {
					panic(err)
				}
				now := time.Now().In(jst)
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(now)
			})
		})
	})

	return r
}
