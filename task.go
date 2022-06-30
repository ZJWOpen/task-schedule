package schedule

import (
	"context"
	"sync"
)

type task struct {
	sync.Mutex

	id        string             // id of the task
	taskState int                // task state
	taskType  int                // type of the task
	runFunc   func(string) error // run the task
	onError   func(error)
	onSucces  func()             // success handler
	ctx       context.Context    // context
	cancel    context.CancelFunc // cancel
}
