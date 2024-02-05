package services

import (
	"fmt"
	"github.com/johnfercher/go-pkg-struct/internal/samples"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceInterpreter_ParseAll(t *testing.T) {
	// Arrange
	sut := NewInterfaceInterpreter()

	// Act
	interfaces := sut.ParseAll(samples.InterfaceFile)

	for _, interf := range interfaces {
		fmt.Println(interf.String())
	}

	// Assert
	assert.NotNil(t, interfaces)
}
