package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gustafer/linkord/configs"
)

type MyClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}
type ContextKey string

var UserIdCtx ContextKey = "userId"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Auth")
		fmt.Println(cookie)
		fmt.Println(err)
		if err != nil {
			http.Error(w, fmt.Sprintf("cookies not set or err reason: %v", err.Error()), http.StatusUnauthorized)
			return
		}
		tokenFromHeader := cookie.Value
		decodedToken, err := jwt.ParseWithClaims(tokenFromHeader, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(configs.LoadAuthKey()), nil
		})
		if err != nil {
			http.Error(w, "Could not authenticate user", http.StatusUnauthorized)
			return
		}
		myClaims := decodedToken.Claims.(*MyClaims)
		ctx := context.WithValue(r.Context(), UserIdCtx, myClaims.UserId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SetHeader(key, value string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(key, value)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
