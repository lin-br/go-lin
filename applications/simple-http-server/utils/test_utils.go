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

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
