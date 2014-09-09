package taskq_test

import (
	"time"
)

type Ctx struct {
	Task1 string
	Task2 int
	Task3 []string
	Task4 float64
}

func sleep() {
	time.Sleep(10 * time.Millisecond)
}

// Simple tasks.
func Task1(ctx interface{}) error {
	ctx.(*Ctx).Task1 = "task1"
	return nil
}

func Task2(ctx interface{}) error {
	ctx.(*Ctx).Task2 = 2
	return nil
}

func Task3(ctx interface{}) error {
	ctx.(*Ctx).Task3 = []string{"task3"}
	return nil
}

func Task4(ctx interface{}) error {
	ctx.(*Ctx).Task4 = 4.4
	return nil
}

// Long running tasks.
func Task1Sleep(ctx interface{}) error {
	sleep()
	ctx.(*Ctx).Task1 = "task1"
	return nil
}

func Task2Sleep(ctx interface{}) error {
	sleep()
	ctx.(*Ctx).Task2 = 2
	return nil
}

func Task3Sleep(ctx interface{}) error {
	sleep()
	ctx.(*Ctx).Task3 = []string{"task3"}
	return nil
}

func Task4Sleep(ctx interface{}) error {
	sleep()
	ctx.(*Ctx).Task4 = 4.4
	return nil
}
