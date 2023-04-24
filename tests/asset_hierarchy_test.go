package tests

import (
	"testing"

	"github.com/personal-asset-manager/src/ds"
)

func TestLTA(t *testing.T) {
	lta := ds.LongTermAsset{}
	if lta.Id != "" {
		t.Fatal("issue")
	}
}
