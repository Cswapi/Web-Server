package models

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const secret = "Cswapi"

type token struct {
	Name string `json:"name"`
	Exp  string `json:"exp"`
}

// 将中间件转换为token进行验证
func ValidateMid(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

	if err == nil {
		if token.Valid {
			next(w, req)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Not authorized")
	}
}
