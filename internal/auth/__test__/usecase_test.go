package __test__

import (
	"errors"
	"go-auth-server/internal/auth"
	"go-auth-server/internal/auth/repo/user"
	"go-auth-server/internal/auth/usecase"
	"testing"
)

func checkEmptyAuthorizationToken(t *testing.T, authorizationToken string) {
	if authorizationToken != "" {
		t.Errorf("Expected to receive an empty authorization token, but instead got %s", authorizationToken)
	}
}

func checkError(t *testing.T, err, expectedErr error) {
	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected to receive %s as error, but instead got %s", expectedErr.Error(), err)
	}
}

func testSucessfulSignUp(t *testing.T, authUsecase *usecase.AuthUsecase) (email, password string) {
	email = "okaabe@okaabe"
	password = "okaabe"

	authorizationToken, signUpError := authUsecase.SignUp("test", email, password)

	if signUpError != nil {
		t.Errorf("Not expected to receive an error of %s", signUpError.Error())
	}

	if authorizationToken == "" {
		t.Errorf("Expected to receive an authorization token, but instead received an empty string")
	}

	return
}

func testDuplicatedSignUp(t *testing.T, authUsecase *usecase.AuthUsecase, email, password string) {
	_, signUpErr := authUsecase.SignUp("aa", email, password)

	checkError(t, signUpErr, auth.ErrDuplicatedUser)
}

func testSucessfulSignIn(t *testing.T, authUsecase *usecase.AuthUsecase, email, password string) string {
	authorizationToken, signInErr := authUsecase.SignIn(email, password)

	if signInErr != nil {
		t.Errorf("Not expected to receive an error, but instead got %s", signInErr.Error())
	}

	if authorizationToken == "" {
		t.Error("Once error is nil, it was expected to receive an authorization token, but instead got an empty string")
	}

	return authorizationToken
}

func testWrongPasswordSignIn(t *testing.T, authUsecase *usecase.AuthUsecase, email string) {
	authorizationToken, signInErr := authUsecase.SignIn(email, "akaka2323@")

	checkError(t, signInErr, auth.ErrInvalidPassword)
	checkEmptyAuthorizationToken(t, authorizationToken)
}

func testInvalidEmailSignIn(t *testing.T, authUsecase *usecase.AuthUsecase) {
	authorizationToken, signInErr := authUsecase.SignIn("aa", "22")

	checkError(t, signInErr, auth.ErrInvalidEmail)
	checkEmptyAuthorizationToken(t, authorizationToken)
}

func testSucessfulRewoke(t *testing.T, authUsecase *usecase.AuthUsecase, authorizationToken string) {
	_, rewokeErr := authUsecase.Rewoke(authorizationToken)

	if rewokeErr != nil {
		t.Errorf("Not expected to receive an error, but instead received %s", rewokeErr.Error())
	}
}

func testInvalidTokenRewoke(t *testing.T, authUsecase *usecase.AuthUsecase) {
	_, rewokeErr := authUsecase.Rewoke("akakakak")

	checkError(t, rewokeErr, auth.ErrInvalidToken)
}

func TestAuthUsecase(t *testing.T) {
	authUsecase := &usecase.AuthUsecase{
		UserRepository: user.NewMockedUserRepository(),
	}

	email, password := testSucessfulSignUp(t, authUsecase)
	testDuplicatedSignUp(t, authUsecase, email, password)

	authorizationToken := testSucessfulSignIn(t, authUsecase, email, password)

	testWrongPasswordSignIn(t, authUsecase, email)
	testInvalidEmailSignIn(t, authUsecase)

	testSucessfulRewoke(t, authUsecase, authorizationToken)
	testInvalidTokenRewoke(t, authUsecase)
}
