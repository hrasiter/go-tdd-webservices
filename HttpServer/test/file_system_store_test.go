package test

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"example.go.com/HttpServer/data"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo", "Wins":10},
			{"Name": "Chris", "Wins":11}]`)

		defer cleanDatabase()

		store := data.FileSystemPlayerStore{Database: database}

		got := store.GetLeague()

		want := []data.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 11},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo", "Wins":10},
			{"Name": "Chris", "Wins":11}]`)

		defer cleanDatabase()

		store := data.FileSystemPlayerStore{Database: database}

		got := store.GetPlayerScore("Chris")
		want := 11
		assertScoresEquals(t, want, got)

	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store := data.FileSystemPlayerStore{Database: database}

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoresEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store := data.FileSystemPlayerStore{Database: database}

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoresEquals(t, got, want)
	})
}

func assertScoresEquals(t testing.TB, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
