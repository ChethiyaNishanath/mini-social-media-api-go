package services

import (
	"fmt"
	"time"

	"github.com/ChethiyaNishanath/social-media-api/src/errors"
	"github.com/ChethiyaNishanath/social-media-api/src/models"
	"github.com/ChethiyaNishanath/social-media-api/src/repository"
	"github.com/google/uuid"
)

func AddComment(postId uuid.UUID, comment models.Comment) (*models.Post, error) {
	post, err := GetPost(postId)

	if err != nil {
		return nil, err
	}

	comment.ID = uuid.New()
	comment.PostID = postId

	curerrentTime := time.Now()
	comment.CreatedAt = curerrentTime
	comment.UpdatedAt = curerrentTime

	post.Comments = append(post.Comments, comment)
	updatedPost, _ := repository.UpdatePost(postId, post)

	return updatedPost, nil
}

func UpdateCommentContent(postId uuid.UUID, updatedContent string, commentId uuid.UUID) (*models.Comment, error) {

	post, err := GetPost(postId)

	if err != nil {
		return nil, err
	}

	repository.PostsMutex.RLock()
	defer repository.PostsMutex.RUnlock()

	var commentToUpdate *models.Comment
	updatedComments := make([]models.Comment, 0, len(post.Comments))

	for _, comment := range post.Comments {
		if comment.ID == commentId {
			comment.Content = updatedContent
			comment.UpdatedAt = time.Now()
			commentToUpdate = &comment
		}
		updatedComments = append(updatedComments, comment)
	}

	if commentToUpdate == nil {
		return nil, errors.NewCommentNotFoundError(fmt.Sprintf("comment not found with the given id: %s", commentId))
	}

	post.Comments = updatedComments

	return commentToUpdate, nil

}

func DeleteComment(postId uuid.UUID, commentId uuid.UUID) ([]models.Comment, error) {

	post, err := GetPost(postId)

	if err != nil {
		return nil, err
	}

	repository.PostsMutex.RLock()
	defer repository.PostsMutex.RUnlock()

	var found bool
	updatedComments := make([]models.Comment, 0, len(post.Comments))

	for _, comment := range post.Comments {
		if comment.ID == commentId {
			found = true
			continue
		}
		updatedComments = append(updatedComments, comment)
	}

	if !found {
		return nil, errors.NewCommentNotFoundError(fmt.Sprintf("comment not found with the given id: %s", commentId))
	}

	post.Comments = updatedComments
	post.UpdatedAt = time.Now()

	return updatedComments, nil
}
