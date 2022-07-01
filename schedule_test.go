package schedule

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	sched := New()

	sched.WithLimit(3)
	tasks := make([]*task, 0)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for _, s := range a {

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		task := NewTask(fmt.Sprintf("%d", s)).
			WithRunFunc(runtestTaskfunc).
			WithOnError(func(id string, err error) {
				fmt.Println("task run failed ", id)
				// t.Error(err)
			}).WithOnComplete(func(id string) {
			t.Logf("task succes,%s", id)
		}).WithContext(ctx).
			WithCancel(cancel)
		tasks = append(tasks, task)

	}
	for _, k := range tasks {
		err := sched.Add(k)
		if err != nil {
			t.Error(err)
		}
	}
	time.Sleep(20 * time.Second)
}

func runtestTaskfunc(ctx context.Context, cancel context.CancelFunc, id string) error {
	fmt.Println("task run ", id)
	select {
	case <-ctx.Done():
		fmt.Println("task cancelled")
	default:
		time.Sleep(10)
	}
	return nil
}
