package myrouter

import (
	"net/http"
	"testing"
)

func TestRouter_GET(t *testing.T) {
	testcases := []struct {
		name     string
		endpoint string
		handler  http.Handler
	}{
		{
			name:     "/のエンドポイントにハンドラを追加する",
			endpoint: "/",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			}),
		},
		{
			name:     "/helloのエンドポインtにハンドラを追加する",
			endpoint: "/hello",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			}),
		},
		{
			name:     "/hogeのエンドポイントにハンドラを追加する",
			endpoint: "/hoge",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			}),
		},
		{
			name:     "/hoge/fugaのエンドポイントにハンドラを追加する",
			endpoint: "/hoge/fuga",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			}),
		},
	}
	r := NewRouter()

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			r.GET(testcase.endpoint, testcase.handler)
		})
	}
}
