package simple_http_server

import (
	"io"

	"github.com/lin-br/go-lin/applications/simple-http-server/model"
)

type FileSystemPlayerStore struct {
	// update the reader, and now we can read one more time
	database io.ReadSeeker
}

func (fs FileSystemPlayerStore) GetLeague() []model.Player {
	_, _ = fs.database.Seek(0, io.SeekStart)
	league, _ := model.NewLeague(fs.database)
	return league
}
