package main

import (
	"context"
	"log"
	"time"
)

func main() {
	initMinio()
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	//createBucket(ctx)
	list := listBuckets(ctx)
	for _, info := range list {
		log.Println(info)
	}
	fileUploader(ctx)
}
