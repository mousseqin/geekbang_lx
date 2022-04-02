package kratos

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"sync"
	"time"
)

func main(){
	hs := http.NewServer()
	gs := grpc.NewServer()

	app := New(
		Name("kratos"),
		Version("v1.0.0"),
		Server(hs, gs),
		Registrar(&mockRegistry{service: make(map[string]*registry.ServiceInstance)}),
	)
	time.AfterFunc(time.Second, func() {
		_ = app.Stop()
	})
	if err := app.Run(); err != nil {
		fmt.Printf("finish err:%v \n",err)
	}
}

type mockRegistry struct {
	lk      sync.Mutex
	service map[string]*registry.ServiceInstance
}

func (r *mockRegistry) Register(ctx context.Context, service *registry.ServiceInstance) error {
	if service == nil || service.ID == "" {
		return fmt.Errorf("no service id")
	}
	r.lk.Lock()
	defer r.lk.Unlock()
	r.service[service.ID] = service
	return nil
}

// Deregister the registration.
func (r *mockRegistry) Deregister(ctx context.Context, service *registry.ServiceInstance) error {
	r.lk.Lock()
	defer r.lk.Unlock()
	if r.service[service.ID] == nil {
		return fmt.Errorf("deregister service not found")
	}
	delete(r.service, service.ID)
	return nil
}
