package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomMasjid(t *testing.T) Masjid {
	arg := CreateMasjidParams{
		Name:       "Test Masjid",
		Address:    "Merdeka",
		City:       "Bogor",
		Coordinate: "12345",
	}

	masjid, err := testQueries.CreateMasjid(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, masjid)

	require.Equal(t, arg.Name, masjid.Name)
	require.Equal(t, arg.Address, masjid.Address)
	require.Equal(t, arg.City, masjid.City)
	require.Equal(t, arg.Coordinate, masjid.Coordinate)

	require.NotZero(t, masjid.ID)
	require.NotZero(t, masjid.CreatedAt)
	return masjid
}

func TestCreateMasjid(t *testing.T) {
	arg := CreateMasjidParams{
		Name:       "Test Masjid",
		Address:    "Merdeka",
		City:       "Bogor",
		Coordinate: "12345",
	}

	masjid, err := testQueries.CreateMasjid(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, masjid)

	require.Equal(t, arg.Name, masjid.Name)
	require.Equal(t, arg.Address, masjid.Address)
	require.Equal(t, arg.City, masjid.City)
	require.Equal(t, arg.Coordinate, masjid.Coordinate)

	require.NotZero(t, masjid.ID)
	require.NotZero(t, masjid.CreatedAt)
}

func TestGetMasjid(t *testing.T) {
	testMasjid := createRandomMasjid(t)

	masjid, err := testQueries.GetMasjid(context.Background(), testMasjid.ID)
	require.NoError(t, err)
	require.NotEmpty(t, masjid)

	require.Equal(t, testMasjid.Name, masjid.Name)
	require.Equal(t, testMasjid.Address, masjid.Address)
	require.Equal(t, testMasjid.City, masjid.City)
	require.Equal(t, testMasjid.Coordinate, masjid.Coordinate)

	require.NotZero(t, masjid.ID)
	require.NotZero(t, masjid.CreatedAt)
	require.WithinDuration(t, testMasjid.CreatedAt, masjid.CreatedAt, time.Second)

}

func TestListMasjids(t *testing.T) {
	createRandomMasjid(t)
	createRandomMasjid(t)

	masjids, err := testQueries.ListMasjids(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, masjids)
}

func TestUpdateMasjid(t *testing.T) {
	testMasjid := createRandomMasjid(t)

	arg := UpdateMasjidParams{
		ID:         testMasjid.ID,
		Name:       "Test Update Masjid",
		Address:    "Merdeka Update",
		City:       "Bogor Update",
		Coordinate: "12345 Update",
	}

	masjid, err := testQueries.UpdateMasjid(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, masjid)

	require.Equal(t, arg.Name, masjid.Name)
	require.Equal(t, arg.Address, masjid.Address)
	require.Equal(t, arg.City, masjid.City)
	require.Equal(t, arg.Coordinate, masjid.Coordinate)

	require.NotZero(t, masjid.ID)
	require.NotZero(t, masjid.CreatedAt)
}

func TestDeleteMasjid(t *testing.T) {
	testMasjid := createRandomMasjid(t)

	err := testQueries.DeleteMasjid(context.Background(), testMasjid.ID)
	require.NoError(t, err)

}
