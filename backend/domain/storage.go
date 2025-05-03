package domain

import (
	"context"
)

type StorageRepository interface {
	// 単一のオブジェクトをアップロード
	Upload(ctx context.Context, key string, data []byte) error
}
