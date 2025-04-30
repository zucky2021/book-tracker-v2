package config

import (
	"backend/domain"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
)

func NewS3Client(env domain.EnvVarProvider) *s3.Client {
	endpoint := env.GetS3Endpoint()
	region := env.GetAWSRegion()

	// 基本設定の読み込み
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("AWS設定の読み込みに失敗: %v", err)
	}

	// カスタムエンドポイントの設定
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		if endpoint != "" {
			o.BaseEndpoint = aws.String(endpoint)
			o.UsePathStyle = true // ローカル環境用にパススタイルを強制
		}
	})

	// 接続検証用コンテキスト
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 接続テスト実行
	if err := verifyConnection(ctx, client); err != nil {
		log.Printf("AWS接続検証失敗: %v", err)
	}

	return client
}

func verifyConnection(ctx context.Context, client *s3.Client) error {
	buckets, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) {
			switch apiErr.ErrorCode() {
			case "AccessDenied":
				return fmt.Errorf("アクセス拒否: IAMポリシーを確認してください (コード: %s)", apiErr.ErrorCode())
			case "InvalidAccessKeyId":
				return fmt.Errorf("不正なアクセスキー: AWS_ACCESS_KEY_IDを確認 (コード: %s)", apiErr.ErrorCode())
			case "SignatureDoesNotMatch":
				return fmt.Errorf("署名不一致: AWS_SECRET_ACCESS_KEYを確認 (コード: %s)", apiErr.ErrorCode())
			default:
				return fmt.Errorf("S3 APIエラー[%s]: %s", apiErr.ErrorCode(), apiErr.ErrorMessage())
			}
		}
		return fmt.Errorf("未知のエラー: %w", err)
	}

	log.Println("=== Check S3 connection - Dump S3 buckets: ===")
	for _, bucket := range buckets.Buckets {
		log.Printf("バケット名: %s", *bucket.Name)
	}
	log.Println("===============================================")

	return nil
}
