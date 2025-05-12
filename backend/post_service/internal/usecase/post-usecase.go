package usecase

import (
	"context"
	"errors"
	"fmt"
	"path"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"

	post_errors "quickflow/post_service/internal/errors"
	"quickflow/post_service/utils/validation"
	"quickflow/shared/models"
	shared_models "quickflow/shared/models"
)

type PostValidator interface {
	ValidateFeedParams(numPosts int, timestamp time.Time) error
}

type PostRepository interface {
	AddPost(ctx context.Context, post models.Post) error
	UpdatePostText(ctx context.Context, postId uuid.UUID, text string) error
	UpdatePostFiles(ctx context.Context, postId uuid.UUID, fileURLs []string) error
	DeletePost(ctx context.Context, postId uuid.UUID) error
	BelongsTo(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (bool, error)
	GetPost(ctx context.Context, postId uuid.UUID) (models.Post, error)
	GetPostsForUId(ctx context.Context, uid uuid.UUID, numPosts int, timestamp time.Time) ([]models.Post, error)
	GetUserPosts(ctx context.Context, id uuid.UUID, requesterId uuid.UUID, numPosts int, timestamp time.Time) ([]models.Post, error)
	GetRecommendationsForUId(ctx context.Context, uid uuid.UUID, numPosts int, timestamp time.Time) ([]models.Post, error)
	GetPostFiles(ctx context.Context, postId uuid.UUID) ([]string, error)
	CheckIfPostLiked(ctx context.Context, postId uuid.UUID, userId uuid.UUID) (bool, error)
	UnlikePost(ctx context.Context, postId uuid.UUID, userId uuid.UUID) error
	LikePost(ctx context.Context, postId uuid.UUID, userId uuid.UUID) error
}

type FileService interface {
	UploadFile(ctx context.Context, file *shared_models.File) (string, error)
	UploadManyFiles(ctx context.Context, files []*shared_models.File) ([]string, error)
	DeleteFile(ctx context.Context, filename string) error
}

type PostUseCase struct {
	postRepo  PostRepository
	fileRepo  FileService
	validator PostValidator
}

// NewPostUseCase creates new post service.
func NewPostUseCase(postRepo PostRepository, fileRepo FileService, validator PostValidator) *PostUseCase {
	return &PostUseCase{
		postRepo:  postRepo,
		fileRepo:  fileRepo,
		validator: validator,
	}
}

// AddPost adds post to the repository.
func (p *PostUseCase) AddPost(ctx context.Context, post models.Post) (*models.Post, error) {
	post.Id = uuid.New()

	var err error
	// Upload files to storage
	post.ImagesURL, err = p.fileRepo.UploadManyFiles(ctx, post.Images)
	if err != nil {
		return nil, fmt.Errorf("p.fileRepo.UploadManyFiles: %w", err)
	}

	// Update post images with urls
	err = p.postRepo.AddPost(ctx, post)
	if err != nil {
		return nil, fmt.Errorf("p.postRepo.AddPost: %w", err)
	}

	return &post, nil
}

// DeletePost removes post from the repository.
func (p *PostUseCase) DeletePost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) error {
	//belongsTo, err := p.postRepo.BelongsTo(ctx, userId, postId)
	//if err != nil {
	//	return post_errors.ErrPostNotFound
	//}

	// TODO user_service
	//if !belongsTo && user.Username != "Nikita" && user.Username != "rvasutenko" {
	//	return post_errors.ErrPostDoesNotBelongToUser
	//}

	// retrieve post files
	postFiles, err := p.postRepo.GetPostFiles(ctx, postId)
	if err != nil {
		return fmt.Errorf("p.postRepo.GetPostFiles: %w", err)
	}

	err = p.postRepo.DeletePost(ctx, postId)
	if err != nil {
		return fmt.Errorf("p.postRepo.DeletePost: %w", err)
	}

	// delete post files
	for _, pic := range postFiles {
		err = p.fileRepo.DeleteFile(ctx, path.Base(pic))
		if err != nil {
			return fmt.Errorf("p.fileRepo.DeleteFile: %w", err)
		}
	}

	return nil
}

// FetchFeed returns feed for user.
func (p *PostUseCase) FetchFeed(ctx context.Context, userId uuid.UUID, numPosts int, timestamp time.Time) ([]models.Post, error) {
	// validate params
	err := p.validator.ValidateFeedParams(numPosts, timestamp)
	if errors.Is(err, validation.ErrInvalidNumPosts) {
		return nil, post_errors.ErrInvalidNumPosts
	} else if errors.Is(err, validation.ErrInvalidTimestamp) {
		return nil, post_errors.ErrInvalidTimestamp
	} else if err != nil {
		return nil, fmt.Errorf("validation.ValidateFeedParams: %w", err)
	}

	// fetch posts
	posts, err := p.postRepo.GetPostsForUId(ctx, userId, numPosts, timestamp)
	if err != nil {
		return nil, fmt.Errorf("p.repo.GetPostsForUId: %w", err)
	}

	return posts, nil
}

// FetchRecommendations returns recommendations for user.
func (p *PostUseCase) FetchRecommendations(ctx context.Context, userId uuid.UUID, numPosts int, timestamp time.Time) ([]models.Post, error) {
	// validate params
	err := p.validator.ValidateFeedParams(numPosts, timestamp)
	if errors.Is(err, validation.ErrInvalidNumPosts) {
		return []models.Post{}, post_errors.ErrInvalidNumPosts
	} else if errors.Is(err, validation.ErrInvalidTimestamp) {
		return []models.Post{}, post_errors.ErrInvalidTimestamp
	} else if err != nil {
		return []models.Post{}, fmt.Errorf("validation.ValidateFeedParams: %w", err)
	}

	// fetch posts
	posts, err := p.postRepo.GetRecommendationsForUId(ctx, userId, numPosts, timestamp)
	if err != nil {
		return []models.Post{}, fmt.Errorf("p.repo.GetRecommendationsForUId: %w", err)
	}

	return posts, nil
}

func (p *PostUseCase) FetchUserPosts(ctx context.Context, userId, requesterId uuid.UUID, numPosts int, timestamp time.Time) ([]models.Post, error) {
	// validate params
	err := p.validator.ValidateFeedParams(numPosts, timestamp)
	if errors.Is(err, validation.ErrInvalidNumPosts) {
		return []models.Post{}, post_errors.ErrInvalidNumPosts
	} else if errors.Is(err, validation.ErrInvalidTimestamp) {
		return []models.Post{}, post_errors.ErrInvalidTimestamp
	} else if err != nil {
		return []models.Post{}, fmt.Errorf("validation.ValidateFeedParams: %w", err)
	}

	// fetch posts
	posts, err := p.postRepo.GetUserPosts(ctx, userId, requesterId, numPosts, timestamp)
	if err != nil {
		return []models.Post{}, fmt.Errorf("p.repo.GetPostsForUId: %w", err)
	}

	return posts, nil
}

func (p *PostUseCase) UpdatePost(ctx context.Context, postUpdate models.PostUpdate, userId uuid.UUID) (*models.Post, error) {
	//// check if user owns the post
	//belongsTo, err := p.postRepo.BelongsTo(ctx, userId, postUpdate.Id)
	//if err != nil {
	//	return nil, fmt.Errorf("p.postRepo.BelongsTo: %w", err)
	//}
	//
	//

	//if !belongsTo {
	//	return nil, post_errors.ErrPostDoesNotBelongToUser
	//}

	oldPost, err := p.postRepo.GetPost(ctx, postUpdate.Id)
	if err != nil {
		return nil, fmt.Errorf("p.postRepo.GetPost: %w", err)
	}

	if oldPost.CreatorId != userId && oldPost.CreatorType == models.PostUser {
		return nil, post_errors.ErrPostDoesNotBelongToUser
	}
	// TODO validate community update

	// retrieve old post photos
	oldPics, err := p.postRepo.GetPostFiles(ctx, postUpdate.Id)
	if err != nil {
		return nil, fmt.Errorf("p.postRepo.GetPostFiles: %w", err)
	}

	// Upload files to storage

	var g errgroup.Group
	fileURLChan := make(chan []string, 1)

	g.Go(func() error {
		var urls []string
		var err error
		if len(postUpdate.Files) > 0 {
			urls, err = p.fileRepo.UploadManyFiles(ctx, postUpdate.Files)
			if err != nil {
				return fmt.Errorf("p.fileRepo.UploadManyFiles: %w", err)
			}
		}
		fileURLChan <- urls
		return nil
	})

	g.Go(func() error {
		if err := p.postRepo.UpdatePostText(ctx, postUpdate.Id, postUpdate.Desc); err != nil {
			return fmt.Errorf("p.postRepo.UpdatePostText: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		fileURLs := <-fileURLChan
		if err := p.postRepo.UpdatePostFiles(ctx, postUpdate.Id, fileURLs); err != nil {
			return fmt.Errorf("p.postRepo.UpdatePostFiles: %w", err)
		}
		return nil
	})

	if err = g.Wait(); err != nil {
		return nil, err
	}
	close(fileURLChan)

	// delete old photos
	for _, pic := range oldPics {
		err = p.fileRepo.DeleteFile(ctx, path.Base(pic))
		if err != nil {
			return nil, fmt.Errorf("p.fileRepo.DeleteFile: %w", err)
		}
	}

	post, err := p.postRepo.GetPost(ctx, postUpdate.Id)
	if err != nil {
		return nil, fmt.Errorf("p.postRepo.GetPost: %w", err)
	}

	return &post, nil
}

func (p *PostUseCase) LikePost(ctx context.Context, postId uuid.UUID, userId uuid.UUID) error {
	if userId == uuid.Nil || postId == uuid.Nil {
		return errors.New("userId or postId is empty")
	}

	liked, err := p.postRepo.CheckIfPostLiked(ctx, postId, userId)
	if err != nil {
		return fmt.Errorf("p.postRepo.CheckIfPostLiked: %w", err)
	}

	if liked {
		return nil // идемпотентность лайка
	}

	err = p.postRepo.LikePost(ctx, postId, userId)
	if err != nil {
		return fmt.Errorf("p.postRepo.LikePost: %w", err)
	}

	return nil
}

func (p *PostUseCase) UnlikePost(ctx context.Context, postId uuid.UUID, userId uuid.UUID) error {
	if userId == uuid.Nil || postId == uuid.Nil {
		return errors.New("userId or postId is empty")
	}

	liked, err := p.postRepo.CheckIfPostLiked(ctx, postId, userId)
	if err != nil {
		return fmt.Errorf("p.postRepo.CheckIfPostLiked: %w", err)
	}

	if !liked {
		return nil // идемпотентность дизлайка
	}

	err = p.postRepo.UnlikePost(ctx, postId, userId)
	if err != nil {
		return fmt.Errorf("p.postRepo.UnlikePost: %w", err)
	}

	return nil
}

func (p *PostUseCase) GetPost(ctx context.Context, postId uuid.UUID) (*models.Post, error) {
	if postId == uuid.Nil {
		return nil, fmt.Errorf("postId is empty")
	}

	post, err := p.postRepo.GetPost(ctx, postId)
	return &post, err
}
