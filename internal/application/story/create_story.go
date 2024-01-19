package story

import "context"

type CreateStory struct {
}

func NewCreateStory() (*CreateStory, error) {
	return &CreateStory{}, nil
}

func (c *CreateStory) Execute(ctx context.Context, request CreateStoryRequest) error {
	return ErrWorldIDIsRequired
}
