package infrastructure

import (
	"backend/domain"
	"testing"
)

func TestEnvVarProvider(t *testing.T) {
	tests := []struct {
		name         string
		envVar       string
		envValue     string
		methodToCall func(domain.EnvVarProvider) string
		want         string
	}{
		{
			name:     "DBホストの取得",
			envVar:   "DB_HOST",
			envValue: "test_db_host",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetDBHost()
			},
			want: "test_db_host",
		},
		{
			name:     "DBポートの取得",
			envVar:   "DB_PORT",
			envValue: "3306",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetDBPort()
			},
			want: "3306",
		},
		{
			name:     "DBユーザーの取得",
			envVar:   "DB_USER",
			envValue: "test_user",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetDBUser()
			},
			want: "test_user",
		},
		{
			name:     "DBパスワードの取得",
			envVar:   "DB_PASSWORD",
			envValue: "test_password",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetDBPassword()
			},
			want: "test_password",
		},
		{
			name:     "DB名の取得",
			envVar:   "DB_NAME",
			envValue: "test_db_name",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetDBName()
			},
			want: "test_db_name",
		},
		{
			name:     "読み込み専用のDBホストの取得",
			envVar:   "DB_READER_HOST",
			envValue: "test_db_reader_host",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetDBReaderHost()
			},
			want: "test_db_reader_host",
		},
		{
			name:     "読み込み専用のDBポートの取得",
			envVar:   "DB_READER_PORT",
			envValue: "3306",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetDBReaderPort()
			},
			want: "3306",
		},
		{
			name:     "読み込み専用のDBユーザーの取得",
			envVar:   "DB_READER_USER",
			envValue: "test_reader_user",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetDBReaderUser()
			},
			want: "test_reader_user",
		},
		{
			name:     "読み込み専用のDBパスワードの取得",
			envVar:   "DB_READER_PASSWORD",
			envValue: "test_reader_password",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetDBReaderPassword()
			},
			want: "test_reader_password",
		},
		{
			name:     "S3エンドポイントの取得",
			envVar:   "S3_ENDPOINT",
			envValue: "test_s3_endpoint",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetS3Endpoint()
			},
			want: "test_s3_endpoint",
		},
		{
			name:     "AWSリージョンの取得",
			envVar:   "AWS_REGION",
			envValue: "ap-northeast-1",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetAWSRegion()
			},
			want: "ap-northeast-1",
		},
		{
			name:     "AWSアクセスキーIDの取得",
			envVar:   "AWS_ACCESS_KEY_ID",
			envValue: "test_access_key_id",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetAWSAccessKeyID()
			},
			want: "test_access_key_id",
		},
		{
			name:     "AWSシークレットアクセスキーの取得",
			envVar:   "AWS_SECRET_ACCESS_KEY",
			envValue: "test_secret_access_key",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetAWSSecretAccessKey()
			},
			want: "test_secret_access_key",
		},
		{
			name:     "S3バケット名の取得",
			envVar:   "S3_BUCKET_NAME",
			envValue: "test_s3_bucket_name",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetS3BucketName()
			},
			want: "test_s3_bucket_name",
		},
		{
			name:     "Google Books APIエンドポイントの取得",
			envVar:   "GOOGLE_BOOKS_API_ENDPOINT",
			envValue: "https://www.googleapis.com/books/v1",
			methodToCall: func(evp domain.EnvVarProvider) string {
				return evp.GetGoogleBooksEndpoint()
			},
			want: "https://www.googleapis.com/books/v1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Setenv(tt.envVar, tt.envValue)
			provider := NewEnvVarProvider()

			if got := tt.methodToCall(provider); got != tt.want {
				t.Errorf("期待される値: %v, 実際の値: %v", tt.want, got)
			}
		})
	}
}
