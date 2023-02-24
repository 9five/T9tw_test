package postgresql

import (
	"T9tw/domain"
	"context"

	"gorm.io/gorm"
)

type postgresqlT9Repo struct {
	db *gorm.DB
}

func NewPostgresqlT9Repo(db *gorm.DB) domain.T9TestRepository {
	return &postgresqlT9Repo{db}
}

func (t *postgresqlT9Repo) UserCheck(ctx context.Context, user *domain.UserTable) (result *domain.UserTable, err error) {
	err = t.db.Model(&domain.UserTable{}).Where(user).First(&result).Error
	return
}
