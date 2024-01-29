package domain

type StoryID string

type Story struct {
	id    StoryID
	world World
	party Party
}

func (s *Story) ID() StoryID {
	return s.id
}

func (s *Story) WorldID() WorldID {
	return s.world.ID()
}

func (s *Story) PartySize() int {
	return s.party.Size()
}

func (s *Story) Characters() []Character {
	return s.party.Characters()
}

type StoryBuilder struct {
	id    *StoryID
	world *World
	party *Party
}

func (sb *StoryBuilder) WithID(id StoryID) *StoryBuilder {
	sb.id = &id
	return sb
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
	if sb.id == nil {
		return nil, NewUnexpectedError("story id is required to build a story")
	}
	if sb.world == nil {
		return nil, NewUnexpectedError("world is required to build a story")
	}
	if sb.party == nil {
		return nil, NewUnexpectedError("party is required to build a story")
	}
	return &Story{
		id:    *sb.id,
		world: *sb.world,
		party: *sb.party,
	}, nil
}
