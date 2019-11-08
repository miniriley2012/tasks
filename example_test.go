package tasks_test

import (
	"errors"
	"fmt"
	"github.com/miniriley2012/tasks"
	"log"
	"time"
)

func Example() {
	task := tasks.New(func() {
		time.Sleep(10 * time.Second)
		fmt.Println("Task finished!")
	})

	task.Run()

	if err := task.Wait(); errors.Is(err, tasks.ErrTaskNotRunning) {
		log.Fatalln(err)
	}
}
