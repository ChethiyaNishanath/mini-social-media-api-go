package repository

import (
	"fmt"
	"sync"
	"time"

	"github.com/ChethiyaNishanath/social-media-api/src/errors"
	"github.com/ChethiyaNishanath/social-media-api/src/models"
	"github.com/google/uuid"
)

var (
	Posts      = make(map[uuid.UUID]*models.Post)
	PostsMutex sync.RWMutex
)

func CreatePost(post *models.Post) (*models.Post, error) {
	PostsMutex.Lock()
	defer PostsMutex.Unlock()

	post.Id = uuid.New()

	Posts[post.Id] = post
	return post, nil
}

func GetPostByID(postId uuid.UUID) (*models.Post, error) {

	post, err := checkPostExistById(postId)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func ListPosts() ([]models.Post, error) {
	PostsMutex.RLock()
	defer PostsMutex.RUnlock()

	var allPosts []models.Post
	for _, post := range Posts {
		allPosts = append(allPosts, *post)
	}
	return allPosts, nil
}

func UpdatePost(postId uuid.UUID, updatedPost *models.Post) (*models.Post, error) {
	PostsMutex.Lock()
	defer PostsMutex.Unlock()

	post, err := checkPostExistById(postId)
	if err != nil {
		return nil, err
	}

	post.Content = updatedPost.Content
	post.Likes = updatedPost.Likes
	post.UpdatedAt = time.Now()

	return post, nil
}

func DeletePost(postId uuid.UUID) (*models.Post, error) {
	PostsMutex.Lock()
	defer PostsMutex.Unlock()

	post, err := checkPostExistById(postId)
	if err != nil {
		return nil, err
	}

	post.Deleted = true
	post.UpdatedAt = time.Now()

	return post, nil
}

func checkPostExistById(postId uuid.UUID) (*models.Post, error) {
	post, exists := Posts[postId]
	if !exists {
		return nil, errors.NewPostNotFoundError(fmt.Sprintf("post not found for the given id: %s", postId))
	}

	if post.Deleted {
		return nil, errors.NewPostAlreadyDeletedError("post has been deleted")
	}

	return post, nil
}
