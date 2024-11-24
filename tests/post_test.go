package tests

import (
	"testing"
	"time"

	"github.com/ChethiyaNishanath/social-media-api/src/models"
	"github.com/ChethiyaNishanath/social-media-api/src/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreatePost(post *models.Post) (*models.Post, error) {
	args := m.Called(post)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (m *MockRepository) ListPosts() ([]models.Post, error) {
	args := m.Called()
	return args.Get(0).([]models.Post), args.Error(1)
}

func (m *MockRepository) GetPostByID(postID uuid.UUID) (*models.Post, error) {
	args := m.Called(postID)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (m *MockRepository) UpdatePost(postId uuid.UUID, post *models.Post) (*models.Post, error) {
	args := m.Called(postId, post)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (m *MockRepository) DeletePost(postID uuid.UUID) (*models.Post, error) {
	args := m.Called(postID)
	return args.Get(0).(*models.Post), args.Error(1)
}

func setupTestData_Post() *MockRepository {
	mockRepo := new(MockRepository)

	post := &models.Post{
		Id:        uuid.New(),
		Author:    "Test Author",
		Content:   "Test Content",
		Likes:     0,
		Comments:  []models.Comment{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}

	mockRepo.On("CreatePost", post).Return(post, nil)
	mockRepo.On("ListPosts").Return([]models.Post{*post}, nil)
	mockRepo.On("GetPostByID", post.Id).Return(post, nil)
	mockRepo.On("UpdatePost", post.Id, post).Return(post, nil)
	mockRepo.On("DeletePost", post.Id).Return(nil, nil)

	return mockRepo
}

func TestCreatePost(t *testing.T) {
	setupTestData_Post()

	post := models.Post{
		Author:  "Test Author",
		Content: "Test Content",
	}

	createdPost, err := services.CreatePost(post)

	assert.NoError(t, err)
	assert.NotNil(t, createdPost)
	assert.Equal(t, "Test Author", createdPost.Author)
	assert.Equal(t, "Test Content", createdPost.Content)
}

func TestListPosts(t *testing.T) {
	setupTestData_Post()

	posts, err := services.ListPosts()

	assert.NoError(t, err)
	assert.Len(t, posts, 1)
	assert.Equal(t, "Test Author", posts[0].Author)
}

func TestGetPost(t *testing.T) {
	setupTestData_Post()
	postID := uuid.New()

	post, err := services.GetPost(postID)

	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.Equal(t, "Test Author", post.Author)
}

func TestUpdatePostContent(t *testing.T) {
	setupTestData_Post()
	postID := uuid.New()

	newContent := "Updated Content"
	updatedPost, err := services.UpdatePostContent(postID, newContent)

	assert.NoError(t, err)
	assert.NotNil(t, updatedPost)
	assert.Equal(t, newContent, updatedPost.Content)
}

func TestDeletePost(t *testing.T) {
	setupTestData_Post()
	postID := uuid.New()

	deletedPost, err := services.DeletePost(postID)

	assert.NoError(t, err)
	assert.Nil(t, deletedPost)
}

func TestLikePost(t *testing.T) {
	setupTestData_Post()
	postID := uuid.New()

	updatedPost, err := services.LikePost(postID)

	assert.NoError(t, err)
	assert.NotNil(t, updatedPost)
	assert.Equal(t, 1, updatedPost.Likes)
}
