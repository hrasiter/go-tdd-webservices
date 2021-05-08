package data

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (im *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return im.store[name]
}

func (im *InMemoryPlayerStore) RecordWin(name string) {
	im.store[name]++
}

func (im *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player

	for name, wins := range im.store {
		league = append(league, Player{Name: name, Wins: wins})
	}
	return league
}
