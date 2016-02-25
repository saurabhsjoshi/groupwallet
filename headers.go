package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AddHeader(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			defer func() {
				if r := recover(); r != nil {
					//Catch & Return error
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(r)
					fmt.Println(r)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			}()
			inner.ServeHTTP(w, r)
		})
}
