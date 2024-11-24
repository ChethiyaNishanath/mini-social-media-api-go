# Social Media API

This project contains API written with Go prgamming language, demostrating simple api for social media post creation,
commenting and likes.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Assumptions](#assumptions)
- [Improvements](#improvements)
- [Installation](#installation)
- [API Endpoints](#api-endpoints)
    - [Create Post](#create-post)
    - [List Posts](#list-posts)
    - [Get Post](#get-post)
    - [Update Post Content](#update-post-content)
    - [Delete Post](#delete-post)
    - [Like Post](#like-post)
    - [Add Comment](#add-comment)
    - [Update Comment](#update-comment)
    - [Delete Comment](#delete-comment)

## Prerequisites

Before running the project, make sure you have the following installed:

- Go 1.18 or later
- Git

## Assumptions

I have developed the API with strict to the following assumptions.

- Multiple posts can have the same content.
- User can create unlimited posts.
- A post should not be totally removed from the data store.
- Users can create unlimited comments on a post.

## Further Work

These features can be implemented in the future.

- Authentication and authorization.
- Limiting single like per a user for the post
- Enabling likes for comments.
- Nested comments.

## Installation

Follow these steps to install and run the project locally.

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/social-media-api.git

2. Install required modules

    ```bash
    go mod tidy

3. Run project

    ```bash
    go run main.go

## API Endpoints

The following endpoints are available for interacting with posts and comments:

### Create Post

- **Endpoint**: `POST /posts`
- **Description**: Create a new post.
- **Request Body**:
    ```json
    {
        "author": "Author Name",
        "content": "Post content"
    }
    ```
- **Response**:
    - **Status**: `200 Ok`
    - **Response Body**:
    ```json
    {
        "id": "uuid-of-the-created-post",
        "author": "uuid-of-the-author",
        "content": "Post content",
        "likes": 0,
        "comments": [],
        "createdAt": "2024-11-24T12:34:56Z",
        "updatedAt": "2024-11-24T12:34:56Z",
        "deleted": false
    }
    ```

### List all Posts

- **Endpoint**: `GET /posts`
- **Description**: Get all posts which are either deleted or not
- **Response**:
    - **Status**: `200 Ok`
    - **Response Body**:
    ```json
    {
        "id": "uuid-of-the-created-post",
        "author": "uuid-of-the-author",
        "content": "Post content",
        "likes": 0,
        "comments": [],
        "createdAt": "2024-11-24T12:34:56Z",
        "updatedAt": "2024-11-24T12:34:56Z",
        "deleted": false
    }
    ```

### Get post by post Id

- **Endpoint**: `GET /posts/{post_id}`
- **Description**: Get post by the provided post_id
- **Response**:
    - **Status**: `200 Ok`
    - **Response Body**:
    ```json
    {
        "id": "uuid-of-the-created-post",
        "author": "uuid-of-the-author",
        "content": "Post content",
        "likes": 0,
        "comments": [],
        "createdAt": "2024-11-24T12:34:56Z",
        "updatedAt": "2024-11-24T12:34:56Z",
        "deleted": false
    }
    ```

### Update Post

- **Endpoint**: `PATCH /posts`
- **Description**: Update content of a post created by a user.
- **Request Body**:
    ```json
    {
        "content": "Post content"
    }
    ```
- **Response**:
    - **Status**: `200 Ok`
    - **Response Body**:
    ```json
    {
        "id": "uuid-of-the-created-post",
        "author": "uuid-of-the-author",
        "content": "Post content [Updated]",
        "likes": 0,
        "comments": [],
        "createdAt": "2024-11-24T12:34:56Z",
        "updatedAt": "2024-11-24T12:34:56Z",
        "deleted": false
    }
    ```

### Delete post by post Id

- **Endpoint**: `DELETE /posts/{post_id}`
- **Description**: Delete a post (post will not be totally removed from the memory only deleted flag will be set
  to `true`)
- **Response**:
    - **Status**: `204 No Content`

### Like a Post

- **Endpoint**: `PATCH /posts/{post_id}/like`
- **Description**: Likes will be incremented by this endpoint
- **Response**:
    - **Status**: `200 Ok`
    - **Response Body**:
    ```json
    {
        "id": "uuid-of-the-created-post",
        "author": "uuid-of-the-author",
        "content": "Post content [Updated]",
        "likes": 1,
        "comments": [],
        "createdAt": "2024-11-24T12:34:56Z",
        "updatedAt": "2024-11-24T12:34:56Z",
        "deleted": false
    }
    ```

### Create a comment

- **Endpoint**: `POST /posts/{post_id}/comment`
- **Description**: Create a new comment for an existing Post.
  - **Request Body**:
      ```json
     {
        "postId": "31b94a1e-3d2c-49fd-9c18-10fe11e17fdc",
        "author": "Jane Doe",
        "content": "This is the updated comment"
     }
    ```
- **Response**:
    - **Status**: `200 Ok`
    - **Response Body**:
    ```json
    {
        "id": "uuid-of-the-post",
        "author": "uuid-of-the-author",
        "content": "This is a post not updated",
        "likes": 0,
        "comments": [
            {
                "id": "uuid-of-the-comment",
                "postId": "uuid-of-the-post",
                "author": "uuid-of-the-author",
                "content": "This is the updated comment",
                "createdAt": "2024-11-24T17:22:00.1922884+05:30",
                "updatedAt": "2024-11-24T17:22:00.1922884+05:30"
            }
        ],
        "createdAt": "2024-11-24T17:21:51.9498475+05:30",
        "updatedAt": "2024-11-24T17:22:11.5465625+05:30",
        "deleted": false
  }
  ```

### update a comment

- **Endpoint**: `PATCH /posts/{post_id}/comment/{comment_id}`
- **Description**: Create a new comment for an existing Post.
    - **Request Body**:
        ```json
       {
          "postId": "uuid-of-the-post",
          "author": "uuid-of-the-author",
          "content": "This is the updated comment"
       }
      ```
- **Response**:
    - **Status**: `200 Ok`
    - **Response Body**:
    ```json
    {
        "id": "dc5aed30-8097-4bc9-80a5-4121581ae022",
        "author": "Chethiya",
        "content": "This is a post not updated",
        "likes": 0,
        "comments": [
            {
                "id": "uuid-of-the-comment",
                "postId": "uuid-of-the-post",
                "author": "uuid-of-the-author",
                "content": "This is the updated comment",
                "createdAt": "2024-11-24T17:22:00.1922884+05:30",
                "updatedAt": "2024-11-24T17:22:00.1922884+05:30"
            }
        ],
        "createdAt": "2024-11-24T17:21:51.9498475+05:30",
        "updatedAt": "2024-11-24T17:22:11.5465625+05:30",
        "deleted": false
  }
  ```

### Delete a comment

- **Endpoint**: `DELETE /posts/{post_id}/comment/{comment_id}`
- **Description**: Delete a comment (comment will be removed from the memory)
- **Response**:
    - **Status**: `204 No Content`