package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vikshith-hg-c/simpleBank/utils"
)

func CreateRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	args := GetAccountParams{
		ID: account1.ID,
	}
	account2, err := testQueries.GetAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.NotZero(t, account2.ID)
	require.WithinDuration(t, account2.CreatedAt, account1.CreatedAt, time.Second)
}

func TestUpadateAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	args := UpdateAccountParams{
		Balance: utils.RandomMoney(),
		ID:      account1.ID,
	}
	getArgs := GetAccountParams{
		ID: account1.ID,
	}
	err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), getArgs)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, args.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.NotZero(t, account2.ID)
	require.WithinDuration(t, account2.CreatedAt, account1.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	testQueries.DeleteAccount(context.Background(), account1.ID)
	getArgs := GetAccountParams{
		ID: account1.ID,
	}
	account2, err := testQueries.GetAccount(context.Background(), getArgs)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2.ID)

}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}
	args := ListAcountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAcounts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}
