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

func AddComment(w http.ResponseWriter, r *http.Request) {
	var newComment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil {
		slog.Error("Failed to decode request body", "error", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	postId, err := uuid.Parse(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid Post Id", http.StatusBadRequest)
		return
	}

	updatedPost, err := services.AddComment(postId, newComment)

	if err != nil {
		slog.Error("Failed to add comment", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	slog.Info("Comment added", "post_id", updatedPost.Id, "comment_author", newComment.Author)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedPost)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {

	postId, err := uuid.Parse(chi.URLParam(r, "post_id"))

	if err != nil {
		slog.Warn("Invalid post ID", "id", chi.URLParam(r, "id"))
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	commentId, err := uuid.Parse(chi.URLParam(r, "comment_id"))

	if err != nil {
		slog.Warn("Invalid comment ID", "id", chi.URLParam(r, "id"))
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var updateRequest struct {
		Content string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updatedComment, err := services.UpdateCommentContent(postId, updateRequest.Content, commentId)
	if err != nil {
		switch e := err.(type) {
		case *errors.CommentNotFoundError:
			slog.Warn("Post not found", "post_id", postId)
			http.Error(w, e.Error(), http.StatusNotFound)
		default:
			slog.Error("Unknown error occurred", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	slog.Debug("Comment updated", "post_id", postId, "comment_id", commentId)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedComment)

}

func DeleteComment(w http.ResponseWriter, r *http.Request) {

	postId, err := uuid.Parse(chi.URLParam(r, "post_id"))

	if err != nil {
		slog.Warn("Invalid post ID", "id", chi.URLParam(r, "id"))
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	commentId, err := uuid.Parse(chi.URLParam(r, "comment_id"))

	if err != nil {
		slog.Warn("Invalid comment ID", "id", chi.URLParam(r, "id"))
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	_, err = services.DeleteComment(postId, commentId)

	if err != nil {
		slog.Warn("Post not found or cannot be deleted", "post_id", postId)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	slog.Debug("Comment deleted", "post_id", postId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
