package tasks_test

import (
	"errors"
	"fmt"
	"github.com/miniriley2012/tasks"
	"testing"
	"time"
)

var task *tasks.Task

func TestNew(t *testing.T) {
	task = tasks.New(func() {
		time.Sleep(10 * time.Second)
		fmt.Println("Hello!")
	})
}

func TestTask_Run(t *testing.T) {
	task.Run()
	task.Run()
}

func TestTask_Wait(t *testing.T) {
	if err := task.Wait(); errors.Is(err, tasks.ErrTaskNotRunning) {
		t.Fatal(err)
	}
}
