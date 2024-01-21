package domain

type Story struct {
	world World
	party Party
}

func (s *Story) WorldID() WorldID {
	return s.world.ID()
}

func (s *Story) Party() Party {
	return s.party
}

type StoryBuilder struct {
	world *World
	party *Party
}

func (sb *StoryBuilder) WithWorld(world *World) *StoryBuilder {
	sb.world = world
	return sb
}

func (sb *StoryBuilder) WithParty(party *Party) *StoryBuilder {
	sb.party = party
	return sb
}

func (sb *StoryBuilder) Build() (*Story, error) {
	if sb.world == nil {
		return nil, NewUnexpectedError("world is required to build a story")
	}
	if sb.party == nil {
		return nil, NewUnexpectedError("party is required to build a story")
	}
	return &Story{
		world: *sb.world,
		party: *sb.party,
	}, nil
}
