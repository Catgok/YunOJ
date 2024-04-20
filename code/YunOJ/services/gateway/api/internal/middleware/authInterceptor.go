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
		uToken := r.Header.Get("U-Token")
		token, err := jwt.Parse(uToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("6B$)*Us(nvOWLx2C"), nil
		})
		if err != nil {
			http.Error(w, "Invalid or missing U-Token", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid or missing U-Token", http.StatusUnauthorized)
			return
		}

		var userType, userId int64
		userType, userId = int64(claims["user_type"].(float64)), int64(claims["user_id"].(float64))
		ctx := context.WithValue(r.Context(), "user_type", userType)
		ctx = context.WithValue(ctx, "user_id", userId)
		next(w, r.WithContext(ctx))
	}
}
