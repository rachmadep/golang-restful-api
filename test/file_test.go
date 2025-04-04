package test

import (
	"rachmadep/golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	conn, cleanup := simple.InitializedConnection("test.txt")
	defer cleanup()

	assert.NotNil(t, conn)
	assert.NotNil(t, conn.File)
	assert.Equal(t, "test.txt", conn.File.Name)
}