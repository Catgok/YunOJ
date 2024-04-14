package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

type AuthInterceptor struct {
}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

func (m *AuthInterceptor) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uToken := r.Header.Get("utoken")
		token, err := jwt.Parse(uToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("6B$)*Us(nvOWLx2C"), nil
		})
		if err != nil {
			http.Error(w, "Invalid or missing utoken", http.StatusUnauthorized)
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid or missing utoken", http.StatusUnauthorized)
		}

		var userType, userId int64
		userType, userId = int64(claims["user_type"].(float64)), int64(claims["user_id"].(float64))
		ctx := context.WithValue(r.Context(), "user_type", userType)
		ctx = context.WithValue(ctx, "user_id", userId)
		next(w, r.WithContext(ctx))
	}
}
