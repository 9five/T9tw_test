package usecase

import (
	"T9tw/domain"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"sort"
	"strconv"
	"time"
)

type t9Usecase struct {
	aesKey      string
	t9Repo      domain.T9TestRepository
	sessionRepo domain.SessionRepository
}

func NewT9TwUsecase(aesKey string, t9Repo domain.T9TestRepository, sessionRepo domain.SessionRepository) domain.T9TestUsecase {
	return &t9Usecase{
		aesKey:      aesKey,
		t9Repo:      t9Repo,
		sessionRepo: sessionRepo,
	}
}

func (t *t9Usecase) SortBySliceInt(ctx context.Context, sInt []int) []int {
	sort.Ints(sInt)
	return sInt
}

func (t *t9Usecase) Login(ctx context.Context, input domain.LoginInput) (token string, err error) {
	user, err := t.t9Repo.UserCheck(ctx, &domain.UserTable{
		Name:     input.Username,
		Password: input.Password,
	})
	if err != nil {
		return "", err
	}

	if input.Username == "admin" && input.Password == "Admin&8181" {
		text := user.Name + user.Password + strconv.Itoa(int(time.Now().UnixNano()))

		c, err := aes.NewCipher([]byte(t.aesKey))
		if err != nil {
			return "", err
		}

		cfb := cipher.NewCFBEncrypter(c, []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f})
		ciphertext := make([]byte, len(text))
		cfb.XORKeyStream(ciphertext, []byte(text))
		token = string(ciphertext)

		t.sessionRepo.New(ctx, &domain.Session{
			Token:  token,
			Expiry: time.Now().AddDate(0, 0, 7),
		})

		return token, nil
	}

	return "", errors.New("not admin")
}
