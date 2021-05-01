package data

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}
