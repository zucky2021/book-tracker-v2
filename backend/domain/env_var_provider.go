package domain

type EnvVarProvider interface {
	GetDBHost() string
	GetDBPort() string
	GetDBUser() string
	GetDBPassword() string
	GetDBName() string
	GetDBReaderHost() string
	GetDBReaderPort() string
	GetDBReaderUser() string
	GetDBReaderPassword() string
	GetS3Endpoint() string
	GetAWSRegion() string
	GetAWSAccessKeyID() string
	GetAWSSecretAccessKey() string
	GetS3BucketName() string
}