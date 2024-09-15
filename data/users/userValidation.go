package users

import "movies_api/internal/validation"

func ValidateEmail(v *validation.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validation.Matches(email, validation.EmailRX), "email", "not a valid email")
}

func ValidatePasswordPlainText(v *validation.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 charachters long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 charachters long")
}

func ValidateUser(v *validation.Validator, user User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not be longer than 500 charachters")

	ValidateEmail(v, user.Email)

	if user.Password.plaintext != nil {
		ValidatePasswordPlainText(v, *user.Password.plaintext)
	}
}
