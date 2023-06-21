package infra

import (
	contact_controller "app/elements/contact/interfaces/controller"
	profile_controller "app/elements/profiles/interfaces/controller"
	rec_controller "app/elements/recreations/interfaces/controller"
	"app/middle/applog"
	"app/middle/authrization"
	"encoding/json"
	"net/http"
	"time"

	"github.com/chihiros/logger"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(logger.Logger)
	r.Use(middleware.Recoverer)

	// Access-Control-Allow-Originを許可する
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://*", "http://*"},
		// AllowedOrigins: []string{"http://localhost:3000", "https://stg.topicpost.net"}, // 特定のオリジンのみ許可
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	ccon := contact_controller.NewContactController()
	// Postgresへのコネクションを取得する
	conn1, err := NewPostgresConnection()
	if err != nil {
		applog.Panic(err)
	}
	rcon := rec_controller.NewRecreationController(conn1)

	conn2, err := NewPostgresConnection()
	if err != nil {
		applog.Panic(err)
	}
	pcon := profile_controller.NewProfileController(conn2)
	r.Route("/v1", func(r chi.Router) {
		// プロフィール用のAPI
		r.Route("/profile", func(r chi.Router) {
			r.Use(authrization.AuthMiddleware) // Dockerで開発するときはコメントアウトする

			r.Get("/", pcon.GetProfiles)
			r.Post("/", pcon.PostProfiles)
			r.Put("/", pcon.PutProfiles)
			r.Delete("/", pcon.DeleteProfiles)
		})

		// レクリエーション用のAPI
		r.Route("/recreation", func(r chi.Router) {
			// JWTが不要なやつ
			// レクリエーションの一覧を取得するためのAPI
			r.Get("/", rcon.GetRecreationsByID)

			// JWTが必要なやつ
			r.With(authrization.AuthMiddleware).Group(func(r chi.Router) {
				// レクリエーションを投稿するためのAPI
				r.Post("/", rcon.PostRecreations)
			})
		})

		// お問い合わせ用のAPI
		r.Route("/contact", func(r chi.Router) {
			r.Post("/", ccon.PostContact)
		})

		// 疎通確認用のAPI
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

		// Example
		r.Route("/example", func(r chi.Router) {
			r.Get("/nojwt", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode("This is NOT JWT protected API.")
			})

			r.With(authrization.AuthMiddleware).Group(func(r chi.Router) {
				r.Get("/jwt", func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode("This is JWT protected API.")
				})
			})
		})
	})

	return r
}
