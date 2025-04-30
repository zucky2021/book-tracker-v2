package infrastructure

import (
	"backend/domain"
	"os"
)

type EnvVarProviderImpl struct{}

func NewEnvVarProvider() domain.EnvVarProvider {
	return &EnvVarProviderImpl{}
}

func (evp *EnvVarProviderImpl) GetDBHost() string {
	return os.Getenv("DB_HOST")
}

func (evp *EnvVarProviderImpl) GetDBPort() string {
	return os.Getenv("DB_PORT")
}

func (evp *EnvVarProviderImpl) GetDBUser() string {
	return os.Getenv("DB_USER")
}

func (evp *EnvVarProviderImpl) GetDBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func (evp *EnvVarProviderImpl) GetDBName() string {
	return os.Getenv("DB_NAME")
}

func (evp *EnvVarProviderImpl) GetDBReaderHost() string {
	return os.Getenv("DB_READER_HOST")
}

func (evp *EnvVarProviderImpl) GetDBReaderPort() string {
	return os.Getenv("DB_READER_PORT")
}

func (evp *EnvVarProviderImpl) GetDBReaderUser() string {
	return os.Getenv("DB_READER_USER")
}

func (evp *EnvVarProviderImpl) GetDBReaderPassword() string {
	return os.Getenv("DB_READER_PASSWORD")
}

func (evp *EnvVarProviderImpl) GetS3Endpoint() string {
	return os.Getenv("S3_ENDPOINT")
}

func (evp *EnvVarProviderImpl) GetAWSRegion() string {
	return os.Getenv("AWS_REGION")
}

func (evp *EnvVarProviderImpl) GetAWSAccessKeyID() string {
	return os.Getenv("AWS_ACCESS_KEY_ID")
}

func (evp *EnvVarProviderImpl) GetAWSSecretAccessKey() string {
	return os.Getenv("AWS_SECRET_ACCESS_KEY")
}

func (evp *EnvVarProviderImpl) GetS3BucketName() string {
	return os.Getenv("S3_BUCKET_NAME")
}

func (evp *EnvVarProviderImpl) GetGoogleBooksEndpoint() string {
	return os.Getenv("GOOGLE_BOOKS_ENDPOINT")
}
