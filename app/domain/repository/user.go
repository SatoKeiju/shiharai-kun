package repository

import (
	"context"

	"github.com/SatoKeiju/shiharai-kun/app/domain/model/user"
)

// User : ユーザーに関するリポジトリ
type User interface {
	FetchByID(ctx context.Context, userID string) (user.User, error)
}
