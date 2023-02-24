package usecase

import (
	"T9tw/domain"
	"context"
	"errors"
	"time"
)

type sessionUsecase struct {
	sessionRepo domain.SessionRepository
}

func NewSessionUsecase(sessionRepo domain.SessionRepository) domain.SessionUsecase {
	return &sessionUsecase{
		sessionRepo: sessionRepo,
	}
}

func (s *sessionUsecase) NewSession(ctx context.Context, session *domain.Session) (result *domain.Session, err error) {
	return s.sessionRepo.New(ctx, session)
}

func (s *sessionUsecase) GetSession(ctx context.Context, session *domain.Session) (result *domain.Session, err error) {
	return s.sessionRepo.Get(ctx, session)
}

func (s *sessionUsecase) AuthSession(ctx context.Context, session *domain.Session) error {
	if time.Now().Sub(session.Expiry).Hours() > 168 {
		return errors.New("token expired")
	}

	return nil
}
