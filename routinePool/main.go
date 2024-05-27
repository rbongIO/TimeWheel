package main

import (
	"fmt"
	"time"
)

func main() {
	pool := NewPool(10)
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		task := NewTask(func() error {
			defer func() {
				fmt.Printf("task[%d] is finished\n", i)
			}()
			for putNum := range ch {
				fmt.Printf("task[%d] is running,Put num:%d\n", i, putNum)
			}
			return nil
		})
		pool.AddTask(task)
	}
	for i := 0; i < 100; i++ {
		ch <- i
		time.Sleep(300 * time.Millisecond)
	}
	close(ch)
	for {
	}
}
