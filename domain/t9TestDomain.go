package domain

import (
	"context"
	"time"
)

type T9Test struct {
}

type T9TestRepository interface {
	UserCheck(ctx context.Context, user *UserTable) (result *UserTable, err error)
}

type T9TestUsecase interface {
	SortBySliceInt(ctx context.Context, sInt []int) []int
	Login(ctx context.Context, input LoginInput) (token string, err error)
}

type UserTable struct {
	Id       int
	Name     string
	Password string
	Created  time.Time
	Updated  time.Time
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
