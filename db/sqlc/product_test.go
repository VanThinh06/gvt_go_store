package db

import (
	"context"
	"log"

	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"
)

func CreateProduct(t *testing.T) Product {

	arg := CreateProductParams{
		IDCategory:  uuid.Must(uuid.Parse("95d51658-12b7-4dbd-8649-16ce98dccb85")),
		Name:        "Vsmart V2",
		Price:       null.IntFrom(0),
		Image:       null.StringFrom(""),
		ListImage:   []string{"https://m.media-amazon.com/images/I/61u0X2mvgjL._AC_SS450_.jpg", ""},
		Description: null.StringFrom(""),
		Sold:        null.IntFrom(0),
		Status:      null.IntFrom(0),
		Sale:        null.IntFrom(0),
	}
	product, err := testQueries.CreateProduct(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Name, product.Name)
	// require.Equal(t, arg.National, product.National)

	// require.NotZero(t, category.ID)
	// require.NotZero(t, category.CreatedAt)
	// require.NotZero(t, category.UpdateAt)
	return product
}

func TestCreateProduct(t *testing.T) {
	CreateProduct(t)
}

// func TestGetProduct(t *testing.T) {

// 	// category1 := CreateCategory(t)
// 	product, err := testQueries.GetProductById(context.Background(), "c7443723-e127-46f4-9a82-dc53c2fae1bf")

// 	require.NoError(t, err)
// 	require.NotEmpty(t, product)

// }

func TestGetLimitProduct(t *testing.T) {

	// category1 := CreateCategory(t)
	listProduct := ListProductParams{
		Limit:  1,
		Offset: 2,
	}
	product2, err := testQueries.ListProduct(context.Background(), listProduct)
	require.NoError(t, err)
	require.NotEmpty(t, product2)
}

func TestGetAllProduct(t *testing.T) {
	product2, err := testQueries.GetAllProduct(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, product2)

	log.Println("product 2  ", product2)
}
