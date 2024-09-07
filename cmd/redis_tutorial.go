package cmd

import (
	"fmt"
	"github.com/rbongIO/TimeWheel/internal/redis"
	"github.com/spf13/cobra"
)

var redis_cmd = &cobra.Command{
	Use:   "redis",
	Short: "连接 redis 进行操作",
	Long:  "将操作封装到命令行工具当中",
	Run:   rediGo,
}

func rediGo(cmd *cobra.Command, args []string) {
	ops := []interface{}{}
	for i := 1; i < len(args); i++ {
		ops = append(ops, args[i])
	}
	do, err := redis.Pool.Get().Do(args[0], ops...)
	if err != nil {
		panic(fmt.Sprintln("Do failed,err:", err))
	}
	fmt.Printf("%s\n", do)
}
