package schedule

import (
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
		tasks = append(tasks, &task{
			id: fmt.Sprintf("%d", s),
			runFunc: func(id string) error {
				t.Logf("task run %s", id)
				time.Sleep(1000000)
				return nil
			},
			onError: func(err error) {
				t.Error(err)
			},
			onSucces: func(id string) {
				t.Logf("task succes,%s", id)
			},
		})
	}
	for _, k := range tasks {
		err := sched.Add(k)
		if err != nil {
			t.Error(err)
		}
		// you can check the task condition
		//
		if k.id == "2" {
			sched.Stop(k.id)
		}
	}
}
