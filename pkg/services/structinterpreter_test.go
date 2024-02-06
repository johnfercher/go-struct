package services_test

import (
	"fmt"
	"github.com/johnfercher/go-pkg-struct/internal/samples"
	"github.com/johnfercher/go-pkg-struct/pkg/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructInterpreter_ParseAll(t *testing.T) {
	// Arrange
	sut := services.NewStructInterpreter()

	// Act
	structs := sut.ParseAll(samples.StructFile, "path")

	for _, struc := range structs {
		fmt.Println(struc.String())
	}

	// Assert
	assert.NotNil(t, structs)
}
