package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccountUser(t *testing.T) {

	arg := CreateAccountParams{
		Email: "vthinh2512agmail.com",

		Password: "Dvanthinh06@",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.TypeUser, account.TypeUser)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.UpdateAt)
}
