package main

import (
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

const (
	endpoint        string = "172.29.1.113:31211"
	accessKeyID     string = "minioadmin"
	secretAccessKey string = "minioadmin"
	useSSL                 = false
)

var (
	client *minio.Client
	err    error
)

func initMinio() {
	client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		panic(err)
	}
	log.Println(client)
}
