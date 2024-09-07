package main

import "github.com/rbongIO/TimeWheel/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		return
	}
}
