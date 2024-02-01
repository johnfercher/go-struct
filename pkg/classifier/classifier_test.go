package classifier_test

import (
	"github.com/johnfercher/go-pkg-struct/pkg/classifier"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClassifier_IsGo(t *testing.T) {
	t.Run("when file has header with package + string, should return is go file", func(t *testing.T) {
		// Arrange
		sut := classifier.New()

		// Act
		isGo := sut.IsGo("package classifier")

		// Assert
		assert.True(t, isGo)
	})
	t.Run("when file hasn´t header with package + string, should return is not go file", func(t *testing.T) {
		// Arrange
		sut := classifier.New()

		// Act
		isGo := sut.IsGo("module github.com/johnfercher/go-pkg-struct")

		// Assert
		assert.False(t, isGo)
	})
}
