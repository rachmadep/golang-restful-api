package main

import (
	"net/http"
	"rachmadep/golang-restful-api/app"
	"rachmadep/golang-restful-api/controller"
	"rachmadep/golang-restful-api/helper"
	"rachmadep/golang-restful-api/middleware"
	"rachmadep/golang-restful-api/repository"
	"rachmadep/golang-restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}