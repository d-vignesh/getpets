package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
)

func ValidateURLParam(params ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			for _, param := range params {
				paramVal := chi.URLParam(r, param)
				if paramVal == "" {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(fmt.Sprintf("url param %s no provided", param)))
					return
				}
				ctx = context.WithValue(ctx, param, paramVal)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ValidateQueryParam(dst interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := reflect.New(reflect.TypeOf(dst)).Interface()
			err := json.NewDecoder(r.Body).Decode(t)
			if err != nil {
				resp := Resp{
					Code: http.StatusBadRequest,
					Msg:  fmt.Sprintf("invalid query param. %s", err.Error()),
				}
				respond(w, r, &resp)
				return
			}
			err = validator.New().Struct(t)
			if err != nil {
				resp := Resp{
					Code: http.StatusBadRequest,
					Msg:  fmt.Sprintf("invalid query param. %s", err.Error()),
				}
				respond(w, r, &resp)
				return
			}
			ctx := context.WithValue(r.Context(), QUERY, t)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ValidateBody(dst interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := reflect.New(reflect.TypeOf(dst)).Interface()
			err := json.NewDecoder(r.Body).Decode(t)
			if err != nil {
				resp := Resp{
					Code: http.StatusBadRequest,
					Msg:  fmt.Sprintf("invalid request body. %s", err.Error()),
				}
				respond(w, r, &resp)
				return
			}
			err = validator.New().Struct(t)
			if err != nil {
				resp := Resp{
					Code: http.StatusBadRequest,
					Msg:  fmt.Sprintf("invalid request body. %s", err.Error()),
				}
				respond(w, r, &resp)
				return
			}
			ctx := context.WithValue(r.Context(), BODY, t)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
