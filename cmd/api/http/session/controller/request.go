package controller

func checkEmpty(values ...string) bool {
	for _, value := range values {
		if value == "" {
			return true
		}
	}

	return false
}

type SignUpRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (body *SignUpRequestBody) Validate() bool {
	return !checkEmpty(body.Email, body.Password)
}

type SignInRequestBody struct {
	Email    string
	Password string
}

func (body *SignInRequestBody) Validate() bool {
	return !checkEmpty(body.Email, body.Password)
}
