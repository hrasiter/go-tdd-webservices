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
