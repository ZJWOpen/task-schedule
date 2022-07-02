# task-schedule
golang task schedule async,timeout,can cancel and stop task

# usage

```

func main() {
	sched := schedule.New()

	sched.WithLimit(3)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for _, s := range a {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		task := schedule.NewTask(fmt.Sprintf("%d", s)).
			WithRunFunc(runtestTaskfunc).
			WithOnError(func(id string, err error) {
				fmt.Println("task run failed ", id, err)
			}).
			WithOnComplete(func(id string) {
				fmt.Println("task succes,", id)
			}).
			WithContext(ctx).
			WithCancel(cancel).
			WithTimeout(6 * time.Second)
		err := sched.Add(task)
		if err != nil {
			fmt.Println(err)
		}
	}
	time.Sleep(20 * time.Second)
}

func runtestTaskfunc(ctx context.Context, id string, timeout time.Duration) error {
	done := make(chan struct{}, 1)
	fmt.Println("task run ", id)
	go func() {
		time.Sleep(5 * time.Second)
		done <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		fmt.Println("task is cancelled", id)
		return errors.New("task cancelled")
	case <-time.After(timeout):
		fmt.Println("task timeout", id)
		return errors.New("task timeout")
	case <-done:
		fmt.Println("run test task finished", id)
		return nil
	}
}

```
