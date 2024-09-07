package cmd

import (
	"fmt"
	"github.com/rbongIO/TimeWheel/internal/ants"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func init() {}

var ant = &cobra.Command{
	Use:   "ants",
	Short: "使用线程池实现打印任务",
	Long:  "创建一个线程池，并执行指定数目的任务",
	Run:   CreateFunc,
}

func CreateFunc(cmd *cobra.Command, args []string) {
	n, _ := strconv.Atoi(args[0])
	wait := make(chan struct{})
	retry := make(chan func(), n)
	err := ants.W.Submit(func() {
		for {
			time.Sleep(time.Second)
			select {
			case f := <-retry:
				err := ants.W.Submit(f)
				if err != nil {
					log.Println(err)
					retry <- f
				}
			}
		}
	})
	if err != nil {
		panic(err)
	}
	for i := 0; i < n; i++ {
		f := func() {
			for j := 0; j < 20; j++ {
				fmt.Printf("job[%d] is running, output:%v\n", i, j)
				r := int(rand.Float32()*200) + 500
				time.Sleep(time.Duration(r) * time.Millisecond)
			}
			wait <- struct{}{}
		}
		now := time.Now()
		err := ants.W.Submit(f)
		fmt.Println(time.Since(now))
		if err != nil {
			log.Println(err)
			retry <- f
		}
	}
	for i := 0; i < n; i++ {
		<-wait
	}
}
