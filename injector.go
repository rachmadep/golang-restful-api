//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"

	"rachmadep/golang-restful-api/app"
	"rachmadep/golang-restful-api/controller"
	"rachmadep/golang-restful-api/middleware"
	"rachmadep/golang-restful-api/repository"
	"rachmadep/golang-restful-api/service"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

var validatorSet = wire.NewSet(
    ProvideValidatorOptions,
    validator.New,
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validatorSet,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}