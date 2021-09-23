package db

import (
	"context"
	"testing"
	"time"

	"github.com/keshavbhattad/bank/db/utils"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.GenerateHashPassword(utils.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       utils.RandomName(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomName(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}

func generateRandomPassword(t *testing.T, password string) (hashedPassword string) {
	if password == "" {
		password = utils.RandomString(6)
	}

	hashedPassword, err := utils.GenerateHashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	return
}
func TestPassword(t *testing.T) {
	password := utils.RandomString(6)
	hashedPassword := generateRandomPassword(t, password)

	err := utils.CheckPassword(password, hashedPassword)
	require.NoError(t, err)

	wrongPassword := utils.RandomString(6)
	err = utils.CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
