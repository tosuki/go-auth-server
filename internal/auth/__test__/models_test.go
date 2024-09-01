package __test__

import (
	"go-auth-server/internal/auth/models"
	"testing"
)

func TestNewUser(t *testing.T) {
	_, err := models.NewUser("tosuki", "kdkddkdk", "ddd")

	if err != nil {
		t.Errorf("Not expected an error to create the user model")
	}
}
