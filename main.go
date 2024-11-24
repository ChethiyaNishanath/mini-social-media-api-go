package main

import (
	"log/slog"
	"net/http"

	"github.com/ChethiyaNishanath/social-media-api/src/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Route("/posts", func(r chi.Router) {
		r.Post("/", handlers.CreatePost)
		r.Get("/", handlers.ListPosts)
		r.Get("/{id}", handlers.GetPost)
		r.Patch("/{id}", handlers.UpdatePostContent)
		r.Delete("/{id}", handlers.DeletePost)
		r.Patch("/{id}/like", handlers.LikePost)
		r.Post("/{id}/comment", handlers.AddComment)
		r.Patch("/{post_id}/comment/{comment_id}", handlers.UpdateComment)
		r.Delete("/{post_id}/comment/{comment_id}", handlers.DeleteComment)
	})

	slog.Info("Server starting", "port", 8080)
	http.ListenAndServe(":3000", r)
}
