package workgroup

import (
	"errors"
	"testing"
)

func TestServerStopped(t *testing.T) {
	var g Group

	stop := make(chan error)
	g.Add(Server(
			func () error {
				<- stop
				return nil
			},
			func () error {
				close(stop)
				return nil
			},
		))

	wait := make(chan struct{})
	err := errors.New("err")
	g.Add(func(<-chan struct{}) error {
		<-wait
		return err
	})
	close(wait)
	assert(t, err, g.Run())
}

