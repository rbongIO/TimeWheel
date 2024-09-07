package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sync"
	"time"
)

var chanRR = &cobra.Command{
	Use:   "chanRR",
	Short: "展示 channel 随机竞争",
	Run:   ChanRR,
}

func ChanRR(cmd *cobra.Command, args []string) {
	c := make(chan int, 10)
	wg := sync.WaitGroup{}
	f := func(idx int) {
		wg.Add(1)
		defer wg.Done()
		for ix := range c {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("[%d] get v:%d", idx, ix)
		}
	}
	go f(1)
	go f(2)
	for i := 0; i < 50; i++ {
		c <- i
	}
	close(c)
	wg.Wait()
}
