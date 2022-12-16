package db

import (
	"context"

	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateCategory(t *testing.T) Category {

	arg := CreateCategoryParams{
		Name:     "Vsmart",
		National: "",
	}
	category, err := testQueries.CreateCategory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Name, category.Name)
	require.Equal(t, arg.National, category.National)

	require.NotZero(t, category.ID)
	require.NotZero(t, category.CreatedAt)
	require.NotZero(t, category.UpdateAt)
	return category
}

func TestCreateCategory(t *testing.T) {
	CreateCategory(t)
}

func TestGetCategory(t *testing.T) {

	// category1 := CreateCategory(t)
	category2, err := testQueries.GetCategoryById(context.Background(), uuid.Must(uuid.Parse("b1c9385c-b0a7-4b60-8ac0-f99b373a3cef")))

	require.NoError(t, err)
	require.NotEmpty(t, category2)

}

func TestGetLimitCategory(t *testing.T) {

	// category1 := CreateCategory(t)
	listCategory := ListCategoryParams{
		Limit:  1,
		Offset: 2,
	}
	category2, err := testQueries.ListCategory(context.Background(), listCategory)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

}

func TestGetAllCategory(t *testing.T) {

	// category1 := CreateCategory(t)
	category2, err := testQueries.GetAllCategory(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, category2)

}
