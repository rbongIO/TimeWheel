package main

import (
	"context"
	minio "github.com/minio/minio-go/v7"
	"log"
)

func createBucket(ctx context.Context) {
	bucketName := "mytable"
	err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "cn-south-1", ObjectLocking: false})
	if err != nil {
		log.Println("make bucket failed, err:", err)
		exites, _ := client.BucketExists(ctx, bucketName)
		if exites {
			log.Println("bucket yi")
		}
		return
	}
	log.Println("创建bucket成功")
}

func listBuckets(ctx context.Context) []minio.BucketInfo {
	buckets, err := client.ListBuckets(ctx)
	if err != nil {
		log.Println("list buckets failed,err:", err)
		return nil
	}
	return buckets
}
