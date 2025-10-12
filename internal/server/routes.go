package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bunkr is running")
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form, limit 50MB
	if err := r.ParseMultipartForm(50 << 20); err != nil {
		http.Error(w, fmt.Sprintf("Could not parse multipart form: %v", err), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not get uploaded file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	dstPath := "./uploads/" + handler.Filename
	if err := os.MkdirAll("./uploads", 0755); err != nil {
		http.Error(w, fmt.Sprintf("Could not create upload directory: %v", err), http.StatusInternalServerError)
		return
	}

	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not create file: %v", err), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, fmt.Sprintf("Could not save file: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File %s uploaded successfully", handler.Filename)
}
