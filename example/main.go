package main

import (
	"context"
	"fmt"
	"time"

	schedule "github.com/task-schedule"
)

func main() {
	sched := schedule.New()

	sched.WithLimit(3)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for _, s := range a {

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		task := schedule.NewTask(fmt.Sprintf("%d", s)).
			WithRunFunc(runtestTaskfunc).
			WithOnError(func(id string, err error) {
				fmt.Println("task run failed ", id)
				// t.Error(err)
			}).
			WithOnComplete(func(id string) {
				fmt.Printf("task succes,%s", id)
				fmt.Println()
			}).
			WithContext(ctx).
			WithCancel(cancel)
		err := sched.Add(task)
		if err != nil {
			fmt.Println(err)
		}

	}
	time.Sleep(20 * time.Second)
}

func runtestTaskfunc(ctx context.Context, id string) error {
	fmt.Println("task run ", id)
	select {
	case <-ctx.Done():
		fmt.Println("task cancelled")
		return nil
	default:
		time.Sleep(10 * time.Second)
	}
	return nil
}
