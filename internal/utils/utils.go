package utils

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gustafer/linkord/internal/middleware"
)

type MyClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func GetUserId(w http.ResponseWriter, r *http.Request) string {
	ctx := r.Context()
	userId := ctx.Value(middleware.UserIdCtx).(string)
	return userId
}
