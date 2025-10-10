package server

import (
	"log"
	"net/http"
)

func Start(port string) {
	go func() {
		if err := http.ListenAndServe(port, nil); err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()
}
