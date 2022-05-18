// Copyright (c) 2022 Orlov Boris onlycodergod@gmail.com.
// This file hello.go is subject to the terms and
// conditions defined in file 'LICENSE', which is part of this project source code.
package main

import (
	"fmt"
	"net/http"
	"runtime"

	log "github.com/inconshreveable/log15"
)

var goVersion string

func main() {
	port := 8000
	goVersion = runtime.Version()

	log.Info("starting server", "port", port)
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func add(a int, b int) int {
	return a + b
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling request", "method", r.Method, "remote_address", r.RemoteAddr, "path", r.URL)
	fmt.Fprintf(w, "Hello from %s", goVersion)
}
