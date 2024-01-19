package story

import (
	"context"
)

type CreateStory struct {
}

func NewCreateStory() (*CreateStory, error) {
	return &CreateStory{}, nil
}

func (c *CreateStory) Execute(_ context.Context, req CreateStoryRequest) error {
	if err := c.validateWorldID(req.WorldID); err != nil {
		return err
	}

	return c.validatePartySize(req.PartySize)
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
