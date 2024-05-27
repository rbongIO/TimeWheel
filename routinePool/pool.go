package main

import (
	"sync"
	"sync/atomic"
)

type Pool struct {
	RunningWorker int32
	Capacity      int32
	JobChan       chan *Task
	sync.Mutex
}

func NewPool(capacity int32) *Pool {
	return &Pool{
		Capacity: capacity,
		JobChan:  make(chan *Task),
	}
}

func (p *Pool) GetCap() int32 {
	return p.Capacity
}

func (p *Pool) incRunning() {
	atomic.AddInt32(&p.RunningWorker, 1)
}

func (p *Pool) decRunning() {
	atomic.AddInt32(&p.RunningWorker, -1)
}

func (p *Pool) getRunningWorker() int32 {
	return atomic.LoadInt32(&p.RunningWorker)
}

func (p *Pool) Run() {
	p.incRunning()
	go func() {
		defer p.decRunning()
		for task := range p.JobChan {
			task.f()
		}
	}()
}

func (p *Pool) AddTask(task *Task) {
	p.Lock()
	defer p.Unlock()
	if p.getRunningWorker() < p.GetCap() {
		p.Run()
	}
	p.JobChan <- task
}
