package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 模拟接力跑步

type runner struct {
	name string
}

func (r runner) running() {
	// Yes, I'm running
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
}

func (r runner) gotBaton(baton <-chan baton) {
	<-baton
}

func (r runner) passedBaton(baton chan<- baton) {
	baton <- struct{}{}
}

type baton struct{}

var gotBaton = make(chan baton)
var pang = make(chan struct{})

func main() {

	runners := runnersInit()
	for _, runner := range runners {
		go runner.running()
	}

	pang <- struct{}{}

}

func runnersInit() []runner {
	return []runner{
		{
			name: "aoa",
		},
		{
			name: "bob",
		},
		{
			name: "coc",
		},
	}
}

func running(gotBaton chan baton, runner runner) {
	select {
	// game beginning
	case <-pang:

	case <-gotBaton:
		fmt.Println(runner.name, "got the baton.")
	}
	fmt.Println(runnerName, "is running.")
	run()
	fmt.Println(runnerName, "passed the baton.")
	gotBaton <- baton{}
}

func run() {

}
