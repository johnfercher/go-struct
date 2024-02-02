package classifiers_test

import (
	"github.com/johnfercher/go-pkg-struct/pkg/services/classifiers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGoPackageName(t *testing.T) {
	t.Run("when file has header with package + string, should return is go file", func(t *testing.T) {
		// Act
		name, err := classifiers.GetGoPackageName("package classifiers")

		// Assert
		assert.Equal(t, "classifiers", name)
		assert.Nil(t, err)
	})
	t.Run("when file hasn´t header with package + string, should return is not go file", func(t *testing.T) {
		// Act
		name, err := classifiers.GetGoPackageName("module github.com/johnfercher/go-pkg-struct")

		// Assert
		assert.Empty(t, name)
		assert.NotNil(t, err)
	})
}

func TestGetGoInterfacesName(t *testing.T) {
	t.Run("when file have interfaces, should return them", func(t *testing.T) {
		// Act
		interfaces, err := classifiers.GetGoInterfacesName(classifiers.InterfaceFileSample)

		// Assert
		assert.Equal(t, 1, len(interfaces))
		assert.Equal(t, "ProductService", interfaces[0])
		assert.Nil(t, err)
	})
	t.Run("when file doesn´t have interfaces, should return false", func(t *testing.T) {
		// Act
		interfaces, err := classifiers.GetGoInterfacesName(classifiers.StructFileSample)

		// Assert
		assert.Empty(t, interfaces)
		assert.NotNil(t, err)
	})
}

func TestGetGoStructsName(t *testing.T) {
	t.Run("when file have structs, should return them", func(t *testing.T) {
		// Act
		structs, err := classifiers.GetGoStructsName(classifiers.StructFileSample)

		// Assert
		assert.Equal(t, 1, len(structs))
		assert.Equal(t, "productService", structs[0])
		assert.Nil(t, err)
	})
	t.Run("when file doesn´t have structs, should return false", func(t *testing.T) {
		// Act
		structs, err := classifiers.GetGoStructsName(classifiers.InterfaceFileSample)

		// Assert
		assert.Empty(t, structs)
		assert.NotNil(t, err)
	})
}

func TestGetGoInterface(t *testing.T) {
	t.Run("When is interface, should return correctly", func(t *testing.T) {
		// Act
		file := classifiers.GetGoInterface(classifiers.InterfaceFileSample)

		// Assert
		assert.Equal(t, "type ProductService interface {\n\tGetByID(ctx context.Context, id string) (*productentities.Product, error)\n\tSearch(ctx context.Context, productType RAW_INTERFACE) ([]*productentities.Product, error)\n\tCreate(ctx context.Context, product *productentities.Product) (*productentities.Product, error)\n\tUpdate(ctx context.Context, product *productentities.Product) (*productentities.Product, error)\n\tDelete(ctx context.Context, id ...string) error\n\tFunc(closure func(RAW_INTERFACE) bool) bool\n}", file)
	})
}
