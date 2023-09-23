package sqlc

import (
	"context"
	"ecom/token"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateSession(t *testing.T) {

	user := CreateRandomUser(t)
	uuid := RandomUUID(t)
	tokenmaker, err := token.NewPasetoMaker("AB5AB82DFBE16254CE9CE486F94FBE7C")
	require.NoError(t, err)
	token, payload, err := tokenmaker.CreateToken(user.Username, time.Minute*15)
	require.NoError(t, err)
	args := CreateSessionParams{
		ID:           uuid,
		Username:     user.Username,
		RefreshToken: token,
		ClientIp:     "127.0.0.1",
		ExpiredAt:    payload.ExpiredAt,
	}
	fmt.Println(token)
	session, err := testQueries.CreateSession(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, session)
	require.Equal(t, session.RefreshToken, token)
	require.Equal(t, session.Username, user.Username)
	require.Equal(t, session.IsBlocked, false)
	require.Equal(t, session.ClientIp, "127.0.0.1")
	require.WithinDuration(t, session.ExpiredAt, payload.ExpiredAt, time.Microsecond)
	require.WithinDuration(t, session.CreatedAt, payload.IssuedAt, time.Second)
}

func RandomUUID(t *testing.T) uuid.UUID {
	uuid, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, uuid)
	require.Greater(t, len(uuid.String()), 0)
	return uuid
}
