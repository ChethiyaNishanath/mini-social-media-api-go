package tests

import (
	"testing"
	"time"

	"github.com/ChethiyaNishanath/social-media-api/src/models"
	"github.com/ChethiyaNishanath/social-media-api/src/repository"
	"github.com/ChethiyaNishanath/social-media-api/src/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupTestData() (uuid.UUID, uuid.UUID, *models.Post) {

	postID := uuid.New()
	commentID := uuid.New()

	comment := models.Comment{
		ID:        commentID,
		PostID:    postID,
		Author:    "Test Author",
		Content:   "Test Comment",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	post := &models.Post{
		Id:        postID,
		Author:    "Test Author",
		Content:   "Test Post",
		Likes:     0,
		Comments:  []models.Comment{comment},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}

	repository.Posts[postID] = post

	return postID, commentID, post
}

func TestDeleteComment_Success(t *testing.T) {
	postID, commentID, _ := setupTestData()

	updatedComments, err := services.DeleteComment(postID, commentID)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(updatedComments), "Expected comments list to be empty after deletion")
}

func TestDeleteComment_PostNotFound(t *testing.T) {
	postID := uuid.New()
	commentID := uuid.New()

	_, err := services.DeleteComment(postID, commentID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "post not found")
}

func TestDeleteComment_CommentNotFound(t *testing.T) {
	postID, _, _ := setupTestData()
	commentID := uuid.New()

	_, err := services.DeleteComment(postID, commentID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "comment not found")
}

func TestDeleteComment_DeletedPost(t *testing.T) {
	postID, commentID, post := setupTestData()

	post.Deleted = true

	_, err := services.DeleteComment(postID, commentID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Post has been deleted")
}
