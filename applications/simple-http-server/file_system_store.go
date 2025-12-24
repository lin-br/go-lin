package simple_http_server

import (
	"io"

	"github.com/lin-br/go-lin/applications/simple-http-server/model"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (fs FileSystemPlayerStore) GetLeague() []model.Player {
	league, _ := model.NewLeague(fs.database)
	return league
}
