package schedule

import (
	"context"
	"sync"
	"time"
)

type task struct {
	sync.Mutex

	id       string                                             // id of the task
	runFunc  func(context.Context, string, time.Duration) error // run the task
	onError  func(string, error)
	onSucces func(string)       // success handler
	ctx      context.Context    // context
	cancel   context.CancelFunc // cancel
	timeout  time.Duration      // timeout for the task
}

func NewTask(id string) *task {
	return &task{
		id: id,
	}
}

func (t *task) WithTimeout(tt time.Duration) *task {
	t.timeout = tt
	return t
}

func (t *task) WithContext(ctx context.Context) *task {
	t.ctx = ctx
	return t
}

func (t *task) WithCancel(cancel context.CancelFunc) *task {
	t.cancel = cancel
	return t
}

func (t *task) WithRunFunc(runFunc func(context.Context, string, time.Duration) error) *task {
	t.runFunc = runFunc
	return t
}

func (t *task) WithOnComplete(onComplete func(string)) *task {
	t.onSucces = onComplete
	return t
}

func (t *task) WithOnError(onError func(string, error)) *task {
	t.onError = onError
	return t
}

func (t *task) stop() {
	if t.cancel != nil {
		t.cancel()
	}
}
