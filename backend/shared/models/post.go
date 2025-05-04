package models

import (
	"time"

	"github.com/google/uuid"
)

type PostCreatorType string

const (
	PostUser      PostCreatorType = "user"
	PostCommunity PostCreatorType = "community"
)

type Post struct {
	Id           uuid.UUID
	CreatorId    uuid.UUID
	CreatorType  PostCreatorType
	Desc         string
	Images       []*File
	ImagesURL    []string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LikeCount    int
	RepostCount  int
	CommentCount int
	IsRepost     bool
}

type PostUpdate struct {
	Id    uuid.UUID
	Desc  string
	Files []*File
}
