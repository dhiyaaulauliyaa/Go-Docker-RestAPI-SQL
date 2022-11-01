package db

import (
	"context"
	"database/sql"
	"testing"

	util "github.com/dhiyaaulauliyaa/learn-go/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomUsername(),
		Name:     util.RandomName(),
		Password: "secret",
		Phone:    util.RandomPhone(),
		Gender:   int32(util.RandomGender()),
		Age:      int32(util.RandomAge()),
		Avatar:   sql.NullString{},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Gender, user.Gender)
	require.Equal(t, arg.Age, user.Age)
	require.Equal(t, arg.Avatar, user.Avatar)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username: util.RandomUsername(),
		Name:     "Test User",
		Password: "secret",
		Phone:    util.RandomPhone(),
		Gender:   1,
		Age:      30,
		Avatar:   sql.NullString{},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Gender, user.Gender)
	require.Equal(t, arg.Age, user.Age)
	require.Equal(t, arg.Avatar, user.Avatar)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}

func TestGetUser(t *testing.T) {
	testUser := createRandomUser(t)

	user, err := testQueries.GetUser(context.Background(), testUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, testUser.Name, user.Name)
	require.Equal(t, testUser.Avatar, user.Avatar)
	require.Equal(t, testUser.Gender, user.Gender)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}

func TestListUsers(t *testing.T) {
	createRandomUser(t)
	createRandomUser(t)

	user, err := testQueries.ListUsers(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, user)
}

func TestUpdateUser(t *testing.T) {
	testUser := createRandomUser(t)

	arg := UpdateUserParams{
		ID:       testUser.ID,
		Username: testUser.Username,
		Name:     "Test User Update",
		Phone:    testUser.Phone,
		Gender:   0,
		Age:      25,
		Avatar:   sql.NullString{},
	}

	user, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Avatar, user.Avatar)
	require.Equal(t, arg.Gender, user.Gender)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}

func TestDeleteUser(t *testing.T) {
	testUser := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), testUser.ID)
	require.NoError(t, err)

}
