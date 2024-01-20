package domain

type Story struct {
	world World
}

func (s *Story) WorldID() WorldID {
	return s.world.ID()
}

type StoryBuilder struct {
	world *World
}

func (sb *StoryBuilder) WithWorld(world *World) *StoryBuilder {
	sb.world = world
	return sb
}

func (sb *StoryBuilder) Build() (*Story, error) {
	if sb.world == nil {
		return nil, NewUnexpectedError("world is required to build a story")
	}
	return &Story{
		world: *sb.world,
	}, nil
}
