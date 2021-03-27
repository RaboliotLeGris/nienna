package objectStorage

import (
	"context"
	"errors"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type ObjectStorage struct {
	minio *minio.Client
}

func NewStorageClient(uri, accessKey, secretKey, bucketName string, ssl bool) (*ObjectStorage, error) {
	minioClient, err := minio.New(uri, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: ssl,
	})
	if err != nil {
		return nil, err
	}

	storage := ObjectStorage{
		minio: minioClient,
	}

	// Checking connection with bucket by ensuring the bucker exists
	storage.EnsureBuckerExist(bucketName)
	return &storage, nil
}

func (s *ObjectStorage) EnsureBuckerExist(bucketName string) error {
	ok, err := s.minio.BucketExists(context.Background(), bucketName)
	if err != nil {
		return err
	}
	if !ok {
		err = s.minio.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "eu-west-1"})
		if err != nil {
			return err
		}
	}
	return errors.New("unable to create bucker")
}

func (s *ObjectStorage) PutObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64) error {
	_, err := s.minio.PutObject(ctx, bucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	return err
}
