package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

func JwtMiddleware(actualHandlingFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
		if !ok || tokenString == "" {
			fmt.Fprintln(w, "Invalid Authentication Token Header", http.StatusUnauthorized)
		}

		token, err := jwt.Parse(tokenString, func(unsignedToken *jwt.Token) (interface{}, error) {
			ok := unsignedToken.Method.Alg()
			if ok != "HS256" {
				return nil, fmt.Errorf("invalid Signing method , Status:%v", http.StatusUnauthorized)
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil || !token.Valid {
			fmt.Fprintln(w, "Parsing Failed, Invalid Token", http.StatusUnauthorized)
			return
		}

		var userName string
		claims := token.Claims
		if claims != nil {
			exp, err := claims.GetExpirationTime()
			if err == nil {
				if exp.Unix() < time.Now().Unix() {
					fmt.Fprintln(w, "Expired Token")
					return
				}

			}
			userName, err = claims.GetSubject()
			if err != nil {
				fmt.Fprintln(w, "Invalid UserName", http.StatusUnauthorized)
				return
			}
			IssuedBy, err := claims.GetIssuer()
			if err != nil || IssuedBy != "Indian Library" {
				fmt.Fprintln(w, "Invalid Token", http.StatusUnauthorized)
				return
			}

		}

		var userNameKey = ContextKey("user_name")
		ctx := context.WithValue(r.Context(), userNameKey, userName)
		NewRequest := r.WithContext(ctx)
		actualHandlingFunc.ServeHTTP(w, NewRequest)

	})

}
