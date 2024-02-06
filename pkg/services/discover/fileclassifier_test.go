package discover_test

import (
	"github.com/johnfercher/go-struct/pkg/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClassifier_IsGoFile(t *testing.T) {
	t.Run("when file has header with package + string, should return is go file", func(t *testing.T) {
		// Arrange
		sut := services.New()

		// Act
		isGo := sut.IsGoFile("package classifiers")

		// Assert
		assert.True(t, isGo)
	})
	t.Run("when file hasn´t header with package + string, should return is not go file", func(t *testing.T) {
		// Arrange
		sut := services.New()

		// Act
		isGo := sut.IsGoFile("module github.com/johnfercher/go-struct")

		// Assert
		assert.False(t, isGo)
	})
}
