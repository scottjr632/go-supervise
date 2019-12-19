package jwt

import (
	"context"
	"net/http"
)

// JWTMiddleware checks for a valid JWT; If valid, adds claims to the request context
// If it is not valid returns an errors
func (j *JWT) JWTMiddleware(r *http.Request, setInContext func(string, interface{})) error {
	token, err := j.TokenExtractor(j.TokenName, r)
	if err != nil {
		return ErrUnableToExtracttoken
	}

	claims, err := j.ValidateToken(token)
	if err != nil {
		return ErrTokenNotValid
	}

	setInContext(string(j.Key), claims)
	return nil
}

// GetClaimsFromContext returns the claims from the context
func (j *JWT) GetClaimsFromContext(c context.Context) (*Claims, error) {
	if claims := c.Value(string(j.Key)); claims != nil {
		return claims.(*Claims), nil
	}

	return nil, ErrClaimsNotInContext
}
