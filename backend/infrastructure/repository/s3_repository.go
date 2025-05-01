package repository

import (
	"backend/domain"
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3RepositoryImpl struct {
	Client     *s3.Client
	BucketName string
}

func NewS3Repository(client *s3.Client, bucketName string) domain.Storage {
	return &S3RepositoryImpl{
		Client:     client,
		BucketName: bucketName,
	}
}

func (s *S3RepositoryImpl) Upload(ctx context.Context, key string, data []byte) error {
	_, err := s.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.BucketName,
		Key:    &key,
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		return err
	}
	return nil
}
