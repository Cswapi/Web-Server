package models

import (
	"errors"
	"net/http"
)

// 中间件
type MiddleWare func(http.ResponseWriter, *http.Request, NextFunc) error
type NextFunc func() error

func MixHandler(middlewares ...MiddleWare) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		idx := -1
		dispatch := func(i int) error {
			var dispatch func(int) error
			dispatch = func(i int) error {
				if i <= idx {
					return errors.New("next called mutiple times")
				}
				idx = i
				if i == len(middlewares) {
					return nil
				}
				fn := middlewares[i]
				return fn(w, req, func(i int) func() error {
					return func() error {
						return dispatch(i)
					}
				}(i+1))
			}
			return dispatch(i)
		}
		dispatch(0)
	}
}
