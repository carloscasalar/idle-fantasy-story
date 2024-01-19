package story

import "context"

type CreateStory struct {
}

func NewCreateStory() (*CreateStory, error) {
	return &CreateStory{}, nil
}

func (c *CreateStory) Execute(_ context.Context, _ CreateStoryRequest) error {
	return ErrWorldIDIsRequired
}
