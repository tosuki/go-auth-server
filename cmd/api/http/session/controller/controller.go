package controller

import (
	"go-auth-server/internal/auth"
	"go-auth-server/internal/auth/usecase"
	"go-auth-server/pkg/utils"

	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
	AuthUsecase *usecase.AuthUsecase
}

func (controller *SessionController) SignUp(context *gin.Context) {
	var signUpRequestBody SignUpRequestBody

	if err := context.ShouldBindJSON(&signUpRequestBody); err != nil && !signUpRequestBody.Validate() {
		utils.RefuseHTTPRequestWithMessage(context, http.StatusBadRequest, "bad-request")
		return
	}

	authorizationToken, signUpErr := controller.AuthUsecase.SignUp(
		signUpRequestBody.Name,
		signUpRequestBody.Email,
		signUpRequestBody.Password,
	)

	if signUpErr != nil {
		switch {
		case errors.Is(signUpErr, auth.ErrDuplicatedUser):
			utils.RefuseHTTPRequestWithError(context, http.StatusConflict, signUpErr)
		case errors.Is(signUpErr, auth.ErrInvalidPasswordFormat):
			utils.RefuseHTTPRequestWithError(context, http.StatusBadRequest, signUpErr)
		default:
			slog.Error("Internal Server Error", "error", signUpErr.Error())
			utils.RefuseHTTPRequestWithMessage(context, http.StatusInternalServerError, "internal-server-error")
		}
		return
	}

	utils.AnswerHTTPRequestWithData(context, http.StatusCreated, authorizationToken)
}

func (controller *SessionController) SignIn(context *gin.Context) {
	var signInRequestBody SignInRequestBody

	if err := context.ShouldBindJSON(&signInRequestBody); err != nil && !signInRequestBody.Validate() {
		utils.RefuseHTTPRequestWithMessage(context, http.StatusBadRequest, "bad-request")
		return
	}

	authorizationToken, signInErr := controller.AuthUsecase.SignIn(
		signInRequestBody.Email,
		signInRequestBody.Password,
	)

	if signInErr != nil {
		switch {
		case errors.Is(signInErr, auth.ErrInvalidEmail):
			utils.RefuseHTTPRequestWithError(context, http.StatusNotFound, signInErr)
		case errors.Is(signInErr, auth.ErrInvalidPassword):
			utils.RefuseHTTPRequestWithError(context, http.StatusUnauthorized, signInErr)
		case errors.Is(signInErr, auth.ErrInvalidPasswordFormat):
			utils.RefuseHTTPRequestWithError(context, http.StatusBadRequest, signInErr)
		default:
			slog.Error("Unhandled error", "err", signInErr.Error())
			utils.RefuseHTTPRequestWithMessage(context, http.StatusInternalServerError, "internal-server-error")
		}
		return
	}

	utils.AnswerHTTPRequestWithData(context, http.StatusAccepted, authorizationToken)
}

func (controller *SessionController) Rewoke(context *gin.Context) {
	authorizationToken := context.GetHeader("authorization")

	if authorizationToken == "" {
		utils.RefuseHTTPRequestWithMessage(context, http.StatusBadRequest, "bad-request")
		return
	}

	session, rewokeErr := controller.AuthUsecase.Rewoke(authorizationToken)

	if rewokeErr != nil {
		switch {
		case errors.Is(rewokeErr, auth.ErrInvalidToken) || errors.Is(rewokeErr, auth.ErrExpiredToken) || errors.Is(rewokeErr, auth.ErrInvalidEmail):
			utils.RefuseHTTPRequestWithError(context, http.StatusUnauthorized, rewokeErr)
		default:
			slog.Error("Unhandled error", "error", rewokeErr.Error())
			utils.RefuseHTTPRequestWithMessage(context, http.StatusInternalServerError, "internal-server-error")
		}
		return
	}

	utils.AnswerHTTPRequestWithData(context, http.StatusFound, session)
}
