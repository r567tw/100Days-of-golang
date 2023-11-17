package main

import (
	"fmt"
    "log"
    "net/http"
    _ "net/http/pprof"
)

func main() {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	// 啟動一個 HTTP 服務器，提供性能剖析端點

	log.Println(http.ListenAndServe("localhost:6060", nil))
}

// go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
// go tool pprof http://localhost:6060/debug/pprof/heap
