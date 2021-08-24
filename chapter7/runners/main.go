package main

import (
	"log"
	"os"
	"runners/runner"
	"time"
)

func main() {
	r := runner.New(3 * time.Second)
	r.Add(task, task, task)
	switch err := r.Start(); err {
	case runner.ErrInterrupt:
		log.Println("Terminating due to interrupt")
		os.Exit(1)
	case runner.ErrTimeout:
		log.Println("Terminating due to timeout")
		os.Exit(1)
	}
}

func task(num int) {
	println(num)
	time.Sleep(time.Duration(num) * time.Second)
}
