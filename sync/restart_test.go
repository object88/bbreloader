package sync

import (
	"context"
	"testing"
	"time"
)

func Test_WithOneInvocation(t *testing.T) {
	count := 0

	f := func(ctx context.Context) {
		time.Sleep(10 * time.Millisecond)
		count++
	}

	r := NewRestarter(f)
	r.Invoke()

	if count != 1 {
		t.Fail()
	}
}

func Test_WithTwoSynchInvocations(t *testing.T) {
	count := 0

	f := func(ctx context.Context) {
		time.Sleep(10 * time.Millisecond)
		select {
		case <-ctx.Done():
			return
		default:
			count++
		}
	}

	r := NewRestarter(f)
	r.Invoke()
	r.Invoke()

	if count != 2 {
		t.Fail()
	}
}

func Test_WithTwoAsyncInvocations(t *testing.T) {
	count := 0

	f := func(ctx context.Context) {
		time.Sleep(10 * time.Millisecond)
		select {
		case <-ctx.Done():
			return
		default:
			count++
		}
	}

	r := NewRestarter(f)
	go func() {
		time.Sleep(1 * time.Millisecond)
		r.Invoke()
	}()
	r.Invoke()

	if count != 1 {
		t.Fail()
	}
}
