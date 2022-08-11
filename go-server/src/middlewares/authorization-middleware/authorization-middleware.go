package authorization_middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fyllekanin/go-server/src/common/error-interface"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

type AuthClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var authorizationHeader = r.Header.Get("authorization")
		var parts = strings.Split(authorizationHeader, " ")
		var errorResponse = &error_interface.RestError{}

		if len(parts) != 2 {
			errorResponse.Message = "Incorrect authorization header"
		} else if strings.ToLower(parts[0]) != "basic" {
			errorResponse.Message = fmt.Sprintf("Scope %s is not supported", parts[0])
		} else if len(parts[1]) == 0 {
			errorResponse.Message = "Invalid token"
		}

		if len(errorResponse.Message) > 0 {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		token, err := jwt.ParseWithClaims(parts[1], &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("AllYourBase"), nil
		})

		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode("Refresh")
			return
		}

		if err != nil || !token.Valid {
			fmt.Println(err)
			errorResponse.Message = "Something went wrong"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		next.ServeHTTP(w, r)
	})
}
