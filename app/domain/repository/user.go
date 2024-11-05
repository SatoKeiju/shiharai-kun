package repository

import (
	"context"

	"github.com/SatoKeiju/shiharai-kun/app/domain/model"
)

// User : ユーザーに関するリポジトリ
type User interface {
	FetchByID(ctx context.Context, userID string) (model.User, error)
}
