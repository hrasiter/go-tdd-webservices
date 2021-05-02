package main

import (
	"fmt"
	"net/http"
	"strings"

	"example.go.com/HttpServer/data"
)

type PlayerServer struct {
	store data.PlayerStore
}

func (p *PlayerServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processWins(rw, player)
	case http.MethodGet:
		p.showScore(rw, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWins(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
