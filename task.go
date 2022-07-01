package schedule

import (
	"context"
	"sync"
)

const (
	taskState = iota
	taskStateRunning
	taskSucceeded
	taskFailed
)

type task struct {
	sync.Mutex

	id        string                                                  // id of the task
	taskState int                                                     // task state
	runFunc   func(context.Context, context.CancelFunc, string) error // run the task
	onError   func(string, error)
	onSucces  func(string)       // success handler
	ctx       context.Context    // context
	cancel    context.CancelFunc // cancel
}

func NewTask(id string) *task {
	return &task{
		id:        id,
		taskState: taskStateRunning,
	}
}

func (t *task) WithContext(ctx context.Context) *task {
	t.ctx = ctx
	return t
}

func (t *task) WithCancel(cancel context.CancelFunc) *task {
	t.cancel = cancel
	return t
}

func (t *task) WithRunFunc(runFunc func(context.Context, context.CancelFunc, string) error) *task {
	t.runFunc = runFunc
	return t
}

func (t *task) WithOnComplete(onComplete func(string)) *task {
	t.onSucces = onComplete
	t.taskState = taskSucceeded
	return t
}

func (t *task) WithOnError(onError func(string, error)) *task {
	t.onError = onError
	t.taskState = taskFailed
	return t
}

func (t *task) Stop() {
	if t.cancel != nil {
		t.cancel()
		t.taskState = taskFailed
	}
}
