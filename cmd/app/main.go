package main

import (
	"context"
	"net/http"
	"time"
)

func main() {

}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	timeout, err := time.ParseDuration(r.FormValue("timeout"))
	if err == nil {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel := context.WithCancel(context.Background())
	}
	defer cancel()

	query := r.FormValue("q")
	if query == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
