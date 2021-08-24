package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner在给定的超时时间内执行一组任务
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt通道报告从操作系统发送的信号
	interrupt chan os.Signal

	// complete通道报告处理任务已经完成
	complete chan error

	// timeout通道报告处理任务已超时
	timeout <-chan time.Time

	// tasks持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// ErrTimeout会在任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt会在接收到操作心态的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New返回一个新得准备使用的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add将一个任务附加到Runner上
// 任务是一个接受int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	// 接受所有的中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的goroutine执行不同任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 任务完成时发出的信号
	case err := <-r.complete:
		return err
	// 任务超时时发出的信号
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行已注册的任务
		task(id)
	}
	return nil
}

// gotInterrupt验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		return true
	default:
		return false
	}
}
