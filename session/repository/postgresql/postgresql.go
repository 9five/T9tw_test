package postgresql

import (
	"T9tw/domain"
	"context"

	"gorm.io/gorm"
)

type postgresqlSessionRepo struct {
	db *gorm.DB
}

func NewPostgresqlSessionRepo(db *gorm.DB) domain.SessionRepository {
	return &postgresqlSessionRepo{db}
}

func (p *postgresqlSessionRepo) New(ctx context.Context, s *domain.Session) (*domain.Session, error) {
	result := p.db.Create(&s)
	return s, result.Error
}

func (p *postgresqlSessionRepo) Get(ctx context.Context, s *domain.Session) (*domain.Session, error) {
	var result *domain.Session
	err := p.db.Model(&domain.Session{}).Where(s).First(&result).Error
	return result, err
}
