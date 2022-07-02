package schedule

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	runSleep(ctx)
	time.Sleep(2 * time.Second)
}

func runSleep(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("task run sleep")
		time.Sleep(5 * time.Second)
	case <-ctx.Done():
		fmt.Println("sleep timed out")
		return
	}
}
