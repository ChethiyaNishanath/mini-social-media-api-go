package services

import (
	"time"

	"github.com/ChethiyaNishanath/social-media-api/src/models"
	"github.com/ChethiyaNishanath/social-media-api/src/repository"
	"github.com/google/uuid"
)

func CreatePost(post models.Post) (*models.Post, error) {
	currentTime := time.Now()
	post.CreatedAt = currentTime
	post.UpdatedAt = currentTime
	post.Deleted = false

	createdPost, err := repository.CreatePost(&post)
	if err != nil {
		return nil, err
	}
	return createdPost, nil
}

func ListPosts() ([]models.Post, error) {
	return repository.ListPosts()
}

func GetPost(postID uuid.UUID) (*models.Post, error) {
	post, err := repository.GetPostByID(postID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func UpdatePostContent(postId uuid.UUID, newContent string) (*models.Post, error) {
	post, err := GetPost(postId)

	if err != nil {
		return nil, err
	}

	post.Content = newContent
	updatedPost, _ := repository.UpdatePost(postId, post)

	return updatedPost, nil
}

func DeletePost(postID uuid.UUID) (*models.Post, error) {
	_, err := repository.DeletePost(postID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func LikePost(postId uuid.UUID) (*models.Post, error) {

	post, err := GetPost(postId)

	if err != nil {
		return nil, err
	}

	post.Likes++
	updatedPost, _ := repository.UpdatePost(postId, post)

	return updatedPost, nil
}
