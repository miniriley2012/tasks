// Package tasks allows for multiple instances of a function to be run and waited on.
package tasks

import (
	"errors"
	"reflect"
	"sync"
)

// ErrTaskNotRunning is returned from Wait if no tasks are running
var ErrTaskNotRunning = errors.New("no task running")

// Task contains the function to be run.
type Task struct {
	start       chan struct{}
	outstanding uint
	waitGroup   sync.WaitGroup
	function    interface{}
}

// New returns a new Task created from function.
func New(function interface{}) *Task {
	return &Task{start: make(chan struct{}), function: function}
}

// Runs the Task and passes args into the function.
func (task *Task) Run(args ...interface{}) {
	var callArgs []reflect.Value
	for _, arg := range args {
		callArgs = append(callArgs, reflect.ValueOf(arg))
	}
	fn := reflect.ValueOf(task.function)
	task.outstanding++
	go func() {
		task.start <- struct{}{}
		task.waitGroup.Add(1)
		fn.Call(callArgs)
		task.waitGroup.Done()
		task.outstanding--
	}()
}

// Waits for all instances of Task created from calling Run to finish.
// Wait returns ErrTaskNotRunning if there are no tasks running.
func (task *Task) Wait() error {
	if task.outstanding == 0 {
		return ErrTaskNotRunning
	}
	<-task.start
	task.waitGroup.Wait()
	return nil
}
