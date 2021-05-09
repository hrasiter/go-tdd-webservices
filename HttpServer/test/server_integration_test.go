package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.go.com/HttpServer/data"
	"example.go.com/HttpServer/handler"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "")
	defer cleanDatabase()
	store := &data.FileSystemPlayerStore{Database: database}
	server := handler.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()

		server.ServeHTTP(response, newGetScoreRequest(player))

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()

		server.ServeHTTP(response, newLeagueRequest())

		assertResponseStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []data.Player{
			{Name: "Pepper", Wins: 3},
		}

		assertLeague(t, got, want)
	})

}
