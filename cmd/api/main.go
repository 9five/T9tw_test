package main

import (
	_t9TestHttp "T9tw/T9Test/delivery/http"
	_t9TestRepo "T9tw/T9Test/repository/postgresql"
	_t9TestUsecase "T9tw/T9Test/usecase"
	"T9tw/config"
	_sessionRepo "T9tw/session/repository/postgresql"
	_sessionUsecase "T9tw/session/usecase"

	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	r := gin.Default()
	setupRouter(r)
	return r
}

func setupRouter(router *gin.Engine) {
	setupT9Test(router)
}

func setupT9Test(router *gin.Engine) {
	sessionRepo := _sessionRepo.NewPostgresqlSessionRepo(config.DB)
	sessionUsecase := _sessionUsecase.NewSessionUsecase(sessionRepo)
	t9TestRepo := _t9TestRepo.NewPostgresqlT9Repo(config.DB)
	t9TestUsecase := _t9TestUsecase.NewT9TwUsecase(config.aesKey, t9TestRepo, sessionRepo)
	_t9TestHttp.NewT9Handler(router, t9TestUsecase, sessionUsecase)
}

func main() {
	r := newRouter()
	r.Run(":8000")
}
