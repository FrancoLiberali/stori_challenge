package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewServiceReturnsErrorIfPublicAPIKeyNotConfigured(t *testing.T) {
	_, err := NewService()
	require.ErrorIs(t, err, ErrEmailAPIKeyNotConfigured)
}

func TestNewServiceReturnsErrorIfPrivateAPIKeyNotConfigured(t *testing.T) {
	t.Setenv(emailPublicAPIKeyEnvVar, "something")

	_, err := NewService()
	require.ErrorIs(t, err, ErrEmailAPIKeyNotConfigured)
}

func TestNewServiceWorksIfBothAPIKeysAreConfigured(t *testing.T) {
	t.Setenv(emailPublicAPIKeyEnvVar, "something_public")
	t.Setenv(emailPrivateAPIKeyEnvVar, "something_private")

	_, err := NewService()
	require.NoError(t, err)
}
