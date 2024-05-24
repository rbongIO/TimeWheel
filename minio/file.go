package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"log"
)

func fileUploader(ctx context.Context) {
	bucketName := "mytable"
	objectName := "table"
	filePath := "../go.mod"
	contentType := "application/text"
	object, err := client.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Println("upload file failed,err: ", err)
		return
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, object.Size)
}
