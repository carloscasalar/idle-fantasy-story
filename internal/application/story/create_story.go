package story

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/carloscasalar/idle-fantasy-story/internal/application"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
)

type Repository interface {
	GetWorldByID(ctx context.Context, worldID domain.WorldID) (*domain.World, error)
}

type CreateStory struct {
	worldRepository Repository
}

func NewCreateStory(worldRepository Repository) (*CreateStory, error) {
	return &CreateStory{worldRepository}, nil
}

func (c *CreateStory) Execute(ctx context.Context, req CreateStoryRequest) (*StoryDTO, error) {
	if err := c.validateWorldID(req.WorldID); err != nil {
		return nil, err
	}

	if err := c.validatePartySize(req.PartySize); err != nil {
		return nil, err
	}

	_, err := c.worldRepository.GetWorldByID(context.Background(), domain.WorldID(req.WorldID))
	if err != nil {
		if errors.Is(err, domain.ErrWorldDoesNotExist) {
			return nil, application.ErrWorldDoesNotExist
		}
		log.WithContext(ctx).WithError(err).Error("unexpected error finding world by ID")
		return nil, application.ErrInternalServer
	}
	return nil, nil
}

func (c *CreateStory) validateWorldID(id string) error {
	if id == "" {
		return ErrWorldIDIsRequired
	}

	return nil
}

func (c *CreateStory) validatePartySize(size *uint8) error {
	if size == nil {
		return nil
	}

	if *size < 3 || *size > 6 {
		return ErrInvalidPartySize
	}

	return nil
}
