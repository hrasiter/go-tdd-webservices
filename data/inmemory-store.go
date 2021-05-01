package data

import "log"

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (im *InMemoryPlayerStore) GetPlayerScore(name string) int {
	log.Println("get score")
	return im.store[name]
}

func (im *InMemoryPlayerStore) RecordWin(name string) {
	im.store[name]++
	log.Println("increment win")
	log.Println(len(im.store))
}
