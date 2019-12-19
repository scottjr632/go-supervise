package jwt

import "errors"

// errors for the JWT package
var (
	ErrTokenNotValid        = errors.New("Token is invalid")
	ErrUnableToExtracttoken = errors.New("Unable to extract token")
	ErrUnableToCreateToken  = errors.New("Unable to create token")
	ErrClaimsNotInContext   = errors.New("Unable to get claims from context")
)
