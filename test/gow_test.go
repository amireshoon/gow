package gow

import (
	"amireshoon/gow/gow"
	"testing"
)

func TestRun(t *testing.T) {

	got := gow.Run()
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}
