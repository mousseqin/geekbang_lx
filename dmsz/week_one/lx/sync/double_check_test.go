package sync

import "testing"

func TestSaveMap_LoadOrStoreHeavy(t *testing.T) {
	s := &SaveMap[string, string]{m: make(map[string]string)}
	s.LoadOrStoreHeavy("name", func() string { return "abc" })
}
