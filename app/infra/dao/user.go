package dao

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	model "github.com/SatoKeiju/shiharai-kun/app/domain/model/user"
	"github.com/SatoKeiju/shiharai-kun/app/domain/repository"
	dto "github.com/SatoKeiju/shiharai-kun/app/infra/dto/user"
)

type user struct {
	rdb *sqlx.DB
}

// NewUser : userリポジトリを生成
func NewUser(rdb *sqlx.DB) repository.User {
	return user{rdb: rdb}
}

// FetchByID : ユーザーIDからユーザー情報を取得
func (u user) FetchByID(ctx context.Context, userID string) (model.User, error) {
	// TODO: "*"をやめて取得するカラムを全て書く
	q := "SELECT * FROM users WHERE id = ?"
	var d dto.DTO
	if err := u.rdb.SelectContext(ctx, d, q, userID); err != nil {
		return model.User{}, fmt.Errorf("rdb.SelectContext(userID: %s): %w", userID, err)
	}

	return dto.ConvertFromDTO(d), nil
}
