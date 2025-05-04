package repository

import (
	"backend/domain"
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3RepositoryImpl struct {
	Client *s3.Client
	env    domain.EnvVarProvider
}

func NewS3Repository(client *s3.Client, env domain.EnvVarProvider) domain.StorageRepository {
	return &S3RepositoryImpl{
		Client: client,
		env:    env,
	}
}

func (repo *S3RepositoryImpl) Upload(ctx context.Context, key string, data []byte) error {
	bucketName := repo.env.GetS3BucketName()
	contentType := http.DetectContentType(data)
	_, err := repo.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &key,
		Body:        bytes.NewReader(data),
		ContentType: &contentType,
	})
	if err != nil {
		return fmt.Errorf("failed to s3 upload (bucket: %s, key=%s): %w", bucketName, key, err)
	}
	return nil
}
