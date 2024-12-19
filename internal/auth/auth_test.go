package auth_test

import (
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetApiKey(t *testing.T) {
	t.Run("no authorization header included", func(t *testing.T) {
		header := make(map[string][]string)
		_, err := auth.GetAPIKey(header)
		if err != auth.ErrNoAuthHeaderIncluded {
			t.Errorf("expected %v, got %v", auth.ErrNoAuthHeaderIncluded, err)
		}
	})
}
