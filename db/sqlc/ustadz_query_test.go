package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUstadz(t *testing.T) Ustadz {
	arg := CreateUstadzParams{
		Name:   "Test Ustadz",
		Avatar: sql.NullString{},
		Gender: 0,
	}

	ustadz, err := testQueries.CreateUstadz(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ustadz)

	require.Equal(t, arg.Name, ustadz.Name)
	require.Equal(t, arg.Avatar, ustadz.Avatar)
	require.Equal(t, arg.Gender, ustadz.Gender)

	require.NotZero(t, ustadz.ID)
	require.NotZero(t, ustadz.CreatedAt)

	return ustadz
}

func TestCreateUstadz(t *testing.T) {
	arg := CreateUstadzParams{
		Name:   "Test Ustadz",
		Avatar: sql.NullString{},
		Gender: 0,
	}

	ustadz, err := testQueries.CreateUstadz(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ustadz)

	require.Equal(t, arg.Name, ustadz.Name)
	require.Equal(t, arg.Avatar, ustadz.Avatar)
	require.Equal(t, arg.Gender, ustadz.Gender)

	require.NotZero(t, ustadz.ID)
	require.NotZero(t, ustadz.CreatedAt)
}

func TestGetUstadz(t *testing.T) {
	testUstadz := createRandomUstadz(t)

	ustadz, err := testQueries.GetUstadz(context.Background(), testUstadz.ID)
	require.NoError(t, err)
	require.NotEmpty(t, ustadz)

	require.Equal(t, testUstadz.Name, ustadz.Name)
	require.Equal(t, testUstadz.Avatar, ustadz.Avatar)
	require.Equal(t, testUstadz.Gender, ustadz.Gender)

	require.NotZero(t, ustadz.ID)
	require.NotZero(t, ustadz.CreatedAt)
}

func TestListUstadzs(t *testing.T) {
	createRandomUstadz(t)
	createRandomUstadz(t)

	ustadz, err := testQueries.ListUstadz(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, ustadz)
}

func TestUpdateUstadz(t *testing.T) {
	testUstadz := createRandomUstadz(t)

	arg := UpdateUstadzParams{
		ID:     testUstadz.ID,
		Name:   "Test Ustadz",
		Avatar: sql.NullString{},
		Gender: 0,
	}

	ustadz, err := testQueries.UpdateUstadz(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ustadz)

	require.Equal(t, arg.Name, ustadz.Name)
	require.Equal(t, arg.Avatar, ustadz.Avatar)
	require.Equal(t, arg.Gender, ustadz.Gender)

	require.NotZero(t, ustadz.ID)
	require.NotZero(t, ustadz.CreatedAt)
}

func TestDeleteUstadz(t *testing.T) {
	testUstadz := createRandomUstadz(t)

	err := testQueries.DeleteUstadz(context.Background(), testUstadz.ID)
	require.NoError(t, err)

}
