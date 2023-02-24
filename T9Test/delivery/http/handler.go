package http

import (
	"T9tw/domain"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// "T9tw/domain"

type T9Handler struct {
	T9Usecase      domain.T9TestUsecase
	SessionUsecase domain.SessionUsecase
}

func NewT9Handler(router *gin.Engine, t9Usecase domain.T9TestUsecase, sessionUsecase domain.SessionUsecase) {
	handler := &T9Handler{
		T9Usecase:      t9Usecase,
		SessionUsecase: sessionUsecase,
	}

	router.GET("/hello", handler.Hello)
	router.POST("/sortnum", handler.SortNum)
	router.POST("/login", handler.Login)
	router.Get("/is_auth", handler.IsAuth)
}

func (t *T9Handler) Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World!")
}

func (t *T9Handler) SortNum(ctx *gin.Context) {
	var input []int
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, t.T9Usecase.SortBySliceInt(ctx, input))
}

func (t *T9Handler) Login(ctx *gin.Context) {
	var input domain.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := t.T9Usecase.Login(ctx, input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (t *T9Handler) IsAuth(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	session, err := t.SessionUsecase.GetSession(ctx, &domain.Session{Token: token})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("session not found"),
		})
		return
	}

	if err = t.SessionUsecase.AuthSession(ctx, session); err != nil {
		ctx.JSON(http.StatusOK, false)
		return
	}

	ctx.JSON(http.StatusOK, true)
	return

}
