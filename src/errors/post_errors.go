package errors

import "fmt"

type PostNotFoundError struct {
	Message string
}

func NewPostNotFoundError(message string) *PostNotFoundError {
	return &PostNotFoundError{Message: message}
}

func (e *PostNotFoundError) Error() string {
	return fmt.Sprintf("PostNotFoundError: %s", e.Message)
}

type PostAlreadyDeletedError struct {
	Message string
}

func NewPostAlreadyDeletedError(message string) *PostAlreadyDeletedError {
	return &PostAlreadyDeletedError{Message: message}
}

func (e *PostAlreadyDeletedError) Error() string {
	return fmt.Sprintf("PostAlreadyDeletedError: %s", e.Message)
}

type CommentNotFoundError struct {
	Message string
}

func NewCommentNotFoundError(message string) *CommentNotFoundError {
	return &CommentNotFoundError{Message: message}
}

func (e *CommentNotFoundError) Error() string {
	return fmt.Sprintf("CommentNotFoundError: %s", e.Message)
}
