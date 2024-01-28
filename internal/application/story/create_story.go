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
	SaveStory(ctx context.Context, story *domain.Story) error
}
type NameGenerator interface {
	GenerateCharacterName(species domain.Species) string
}

type CreateStory struct {
	repository    Repository
	nameGenerator NameGenerator
}

func NewCreateStory(repository Repository, nameGenerator NameGenerator) (*CreateStory, error) {
	return &CreateStory{repository, nameGenerator}, nil
}

func (c *CreateStory) Execute(ctx context.Context, req CreateStoryRequest) (*StoryDTO, error) {
	// TODO push into StoryFactory
	if err := c.validateWorldID(req.WorldID); err != nil {
		return nil, err
	}

	// TODO push into StoryFactory
	if err := c.validatePartySize(req.PartySize); err != nil {
		return nil, err
	}

	world, err := c.repository.GetWorldByID(context.Background(), domain.WorldID(req.WorldID))
	if err != nil {
		if errors.Is(err, domain.ErrWorldDoesNotExist) {
			return nil, application.ErrWorldDoesNotExist
		}
		log.WithContext(ctx).WithError(err).Error("unexpected error finding world by ID")
		return nil, application.ErrInternalServer
	}

	story, err := new(domain.StoryFactory).
		WithNameGenerator(c.nameGenerator).
		WithWorld(world).
		WithPartySize(req.PartySize).
		Build()

	if err := c.repository.SaveStory(ctx, story); err != nil {
		log.WithContext(ctx).WithError(err).Error("unexpected error saving story")
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
