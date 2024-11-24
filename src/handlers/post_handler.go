package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ChethiyaNishanath/social-media-api/src/errors"
	"github.com/ChethiyaNishanath/social-media-api/src/models"
	"github.com/ChethiyaNishanath/social-media-api/src/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		slog.Error("Failed to decode request body", "error", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdPost, err := services.CreatePost(post)
	if err != nil {
		slog.Error("Error creating post", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	slog.Debug("Post created", "post_id", createdPost.Id, "author", createdPost.Author)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdPost)
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := services.ListPosts()
	if err != nil {
		slog.Error("Error retrieving posts", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	slog.Debug("Listed all posts", "total_posts", len(posts))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	postId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		slog.Warn("Invalid post ID", "id", chi.URLParam(r, "id"))
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	post, err := services.GetPost(postId)
	if err != nil {
		slog.Warn("Post not found", "post_id", postId)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	slog.Debug("Post retrieved", "post_id", postId)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func UpdatePostContent(w http.ResponseWriter, r *http.Request) {
	postId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		slog.Warn("Invalid post ID", "id", chi.URLParam(r, "id"))
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var payload struct {
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		slog.Error("Failed to decode request body", "error", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	post, err := services.UpdatePostContent(postId, payload.Content)
	if err != nil {
		slog.Warn("Post not found or cannot be updated", "post_id", postId)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	slog.Debug("Post updated", "post_id", postId)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	postId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		slog.Warn("Invalid post ID", "id", chi.URLParam(r, "id"))
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = services.DeletePost(postId)
	if err != nil {
		slog.Warn("Post not found or cannot be deleted", "post_id", postId)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	slog.Debug("Post deleted", "post_id", postId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	postId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		slog.Warn("Invalid post ID", "id", chi.URLParam(r, "id"))
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	post, err := services.LikePost(postId)

	if err != nil {
		switch e := err.(type) {
		case *errors.PostNotFoundError:
			slog.Warn("Post not found", "post_id", postId)
			http.Error(w, e.Error(), http.StatusNotFound)
		case *errors.PostAlreadyDeletedError:
			slog.Warn("Post already deleted", "post_id", postId)
			http.Error(w, e.Error(), http.StatusGone)
		default:
			slog.Error("Unknown error occurred", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	slog.Debug("Post liked", "post_id", postId, "total_likes", post.Likes)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
