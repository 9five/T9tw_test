package domain

import (
	"context"
	"time"
)

type Session struct {
	Token  string
	Expiry time.Time
}

type SessionRepository interface {
	New(ctx context.Context, s *Session) (*Session, error)
	Get(ctx context.Context, s *Session) (*Session, error)
}

type SessionUsecase interface {
	NewSession(ctx context.Context, s *Session) (result *Session, err error)
	GetSession(ctx context.Context, session *Session) (result *Session, err error)
	AuthSession(ctx context.Context, session *Session) error
}
