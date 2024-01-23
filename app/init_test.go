package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewServiceReturnsErrorIfPublicAPIKeyNotConfigured(t *testing.T) {
	_, err := NewService(nil)
	require.ErrorIs(t, err, ErrEmailAPIKeyNotConfigured)
}

func TestNewServiceReturnsErrorIfPrivateAPIKeyNotConfigured(t *testing.T) {
	t.Setenv(EmailPublicAPIKeyEnvVar, "something")

	_, err := NewService(nil)
	require.ErrorIs(t, err, ErrEmailAPIKeyNotConfigured)
}

func TestNewServiceWorksIfBothKeysAreConfigured(t *testing.T) {
	t.Setenv(EmailPublicAPIKeyEnvVar, "something_public")
	t.Setenv(EmailPrivateAPIKeyEnvVar, "something_private")

	_, err := NewService(nil)
	require.NoError(t, err)
}

func TestNewDBConnectionReturnErrorIfEnvVarsNotConfigured(t *testing.T) {
	_, err := NewDBConnection()
	require.ErrorIs(t, err, ErrDatabaseNotConfigured)
}
