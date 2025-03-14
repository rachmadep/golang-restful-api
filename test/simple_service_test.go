package test

import (
	"fmt"
	"rachmadep/golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	fmt.Println(err)
	fmt.Println(simpleService.SimpleRepository)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	fmt.Println(err)
	fmt.Println(simpleService)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}