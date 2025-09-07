package simple_http_server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const JsonContentType = "application/json"

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeagueTable() []Player
}

type PlayerServer struct {
	Store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	ps := new(PlayerServer)
	ps.Store = store
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(ps.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(ps.playersHandler))
	ps.Handler = router
	return ps
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Handler.ServeHTTP(w, r)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", JsonContentType)
	json.NewEncoder(w).Encode(p.Store.GetLeagueTable())
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	_, _ = fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
