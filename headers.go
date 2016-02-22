package main

import "net/http"

func AddHeader(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			inner.ServeHTTP(w, r)
		})
}
