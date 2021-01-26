package gow

import (
	"amireshoon/gow/gow"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	got := gow.GetVersion()

	if reflect.TypeOf(got).String() != "string" {
		t.Errorf("got %s, wanted %s", got, reflect.TypeOf(got).String())
	}

}
