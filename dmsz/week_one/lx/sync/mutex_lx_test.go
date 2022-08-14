package sync

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestAll(t *testing.T) {
	a := NewArrayList[int](10)
	//a.Set(0, 0)
	//a.Set(1, 1)
	//a.Set(2, 2)
	//a.Set(3, 3)
	//
	//a.Delete(2)

	a.Append(10)
	spew.Dump(a)
	a.Append(11)
	spew.Dump(a)
}
