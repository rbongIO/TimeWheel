package main

type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	return &Task{f: f}
}
