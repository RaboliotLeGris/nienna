package main

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewStorageClient(uri, accessKey, secretKey, bucketName string, ssl bool) (*minio.Client, error) {
	storage, err := minio.New(uri, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: ssl,
	})
	if err != nil {
		return nil, err
	}

	// Checking connection with bucket by ensuring the bucker exists
	ok, err := storage.BucketExists(context.Background(), bucketName)
	if err != nil {
		return nil, err
	}
	if !ok {
		err = storage.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "eu-west-1"})
		if err != nil {
			return nil, err
		}
	}
	return storage, nil
}
