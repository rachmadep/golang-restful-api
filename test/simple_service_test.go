package test

import (
	"fmt"
	"rachmadep/golang-restful-api/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializedService()
	fmt.Println(simpleService.SimpleRepository)
}