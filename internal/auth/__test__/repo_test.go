package __test__

import (
	"errors"
	"go-auth-server/internal/auth"
	"go-auth-server/internal/auth/models"
	"go-auth-server/internal/auth/repo/user"
	"testing"
)

func testAddNewUser(t *testing.T, userRepository user.UserRepository, email string) {
	err := userRepository.Add(&models.User{
		Id:        "0",
		Name:      "okaabe",
		Email:     email,
		Password:  "a",
		CreatedAt: 0,
		UpdatedAt: 0,
	})

	if err != nil {
		t.Errorf("Not expected to receive an error, but instead got %s", err.Error())
	}
}

func testDuplicatedUser(t *testing.T, userRepository user.UserRepository, email string) {
	err := userRepository.Add(&models.User{
		Id:        "0",
		Name:      "okaabe",
		Email:     email,
		Password:  "a",
		CreatedAt: 0,
		UpdatedAt: 0,
	})

	if !errors.Is(err, auth.ErrDuplicatedUser) {
		t.Errorf("Expected to receive an error of duplicated item, but instead got %s", err.Error())
	}
}

func TestMockedRepository(t *testing.T) {
	userRepository := user.NewMockedUserRepository()

	testAddNewUser(t, userRepository, "okaabe@okaabe")
	testDuplicatedUser(t, userRepository, "okaabe@okaabe")
}
