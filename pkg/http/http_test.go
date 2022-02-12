package http_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/assert.v1"
)

func TestPetRoutes(t *testing.T) {
	tt := []struct {
		desc       string
		method     string
		uri        string
		payload    []byte
		headers    map[string]string
		StatusCode int
	}{
		{
			desc:   "create pet",
			method: "POST",
			uri:    "/pets/",
			payload: []byte(`{"category": "dog", "breed": "labour", "age": 2, "gender": "male", 
								"colors": "black", "contact": {"owner": "bigshow", "phone": "123456789"}, "price": 10000}`),
			headers:    nil,
			StatusCode: 200,
		},
		{
			desc:       "list pets",
			method:     "GET",
			uri:        "/pets?category=dog",
			payload:    nil,
			headers:    nil,
			StatusCode: 200,
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			var payload io.Reader
			if tc.payload != nil {
				payload = bytes.NewBuffer(tc.payload)
			}
			req, err := http.NewRequest(tc.method, "localhost:6000"+tc.uri, payload)
			assert.NoError(t, err)

			for k, v := range tc.headers {
				req.Header.Set(k, v)
			}

			res, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)

			assert.Equal(t, tc.StatusCode, res.StatusCode)
		})
	}
}
