package sync

import (
	"context"
	"sync"
)

// RestartableFn is a function that will be cancelled by Restarter
// if a new invocation is started.
type RestartableFn func(ctx context.Context)

// Restarter is a synchronizing structure that acts as a retriggering
// gate for another method.  The method will be started with Invoke,
// and must accept a cancellable context.  If a second call to Invoke
// is made while one is already running, the first one will be cancelled,
// and then the second will be started.
type Restarter struct {
	f RestartableFn
	m *sync.Mutex
	c *context.CancelFunc
}

// NewRestarter creatnes a new restarter struct
func NewRestarter(f RestartableFn) *Restarter {
	return &Restarter{f, &sync.Mutex{}, nil}
}

// Invoke calls the associated method, first cancelling any existing
// call
func (r *Restarter) Invoke() {
	// Ensure that we are the only running method
	ctx, cancelFn := r.spinUp()

	// Do the work
	r.f(ctx)

	// Clean up
	r.spinDown(cancelFn)
}

func (r *Restarter) spinUp() (context.Context, context.CancelFunc) {
	r.m.Lock()

	if r.c != nil {
		(*r.c)()
	}

	ctx, cancelFn := context.WithCancel(context.Background())
	r.c = &cancelFn

	r.m.Unlock()

	return ctx, cancelFn
}

func (r *Restarter) spinDown(cancelFn context.CancelFunc) {
	r.m.Lock()
	cancelFn()
	r.c = nil
	r.m.Unlock()
}
