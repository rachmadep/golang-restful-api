package main

import (
	"net/http"
	"rachmadep/golang-restful-api/helper"
	"rachmadep/golang-restful-api/middleware"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func ProvideValidatorOptions() []validator.Option {
    return []validator.Option{}
}

func main() {
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}