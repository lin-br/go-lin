package utils

import (
	"reflect"
	"testing"

	"github.com/lin-br/go-lin/applications/simple-http-server/model"
)

func AssertLeague(t testing.TB, got, want []model.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
