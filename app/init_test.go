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

func TestNewServiceReturnsErrorIfDatabaseInfoIsNotConfigured(t *testing.T) {
	t.Setenv(EmailPublicAPIKeyEnvVar, "something_public")
	t.Setenv(EmailPrivateAPIKeyEnvVar, "something_private")

	_, err := NewService(nil)
	require.ErrorIs(t, err, ErrDatabaseNotConfigured)
}
