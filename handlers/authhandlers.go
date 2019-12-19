package handlers

import (
	"context"
	"errors"
	"go-supervise/jwt"
	"go-supervise/services/password"
	"net/http"
	"time"
)

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *handlers) buildAuthHandlers(unprotectedGroup Routable, protectedGroup Routable) {
	// group.GET("/", h.getWorkersHealth)
	unprotectedGroup.POST("/", convertToGinHandler(h.doBasicAuth))
	protectedGroup.GET("/")
}

func (h *handlers) doBasicAuth(c context.Context, w http.ResponseWriter, r *http.Request) {
	req := &authRequest{}
	if err := readJSON(r, req); err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}
	pass, err := password.GetBasicPassword()
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	if req.Password == pass && req.Username == password.GetBasicUser() {
		if token, err := h.jwt.NewToken(password.GetBasicUser()); err != nil {
			writeError(w, errors.New("User is not authorized"), http.StatusUnauthorized)
		} else {
			cookie := &http.Cookie{
				Name:     h.jwt.TokenName,
				Value:    token,
				Path:     h.jwt.ProtectedPath,
				Expires:  time.Now().AddDate(0, 0, 1),
				HttpOnly: true,
			}
			http.SetCookie(w, cookie)
			writeJSON(w, &struct{ Message string }{"Authorized"})
		}
	} else {
		writeError(w, errors.New("Password was incorrect"), http.StatusUnauthorized)
	}
}

func (h *handlers) checkAuthentications(c context.Context, w http.ResponseWriter, r *http.Request) {
	claims := c.Value(h.jwt.Key).(jwt.Claims)
	writeJSON(w, &struct{ Username string }{claims.Username})
}
