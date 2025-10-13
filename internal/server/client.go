package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/rrajvardhan/bunkr/internal/tui/shared"
)

type FileInfo struct {
	Name string  `json:"name"`
	Size float64 `json:"size_kb"`
}

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("uploads")
	if err != nil {
		http.Error(w, "failed to read uploads directory", http.StatusInternalServerError)
		return
	}

	var list []FileInfo
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		info, err := f.Info()
		if err != nil {
			continue
		}
		list = append(list, FileInfo{
			Name: f.Name(),
			Size: float64(info.Size()) / 1024,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(list)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := filepath.Base(r.URL.Path[len("/download/"):])
	if filename == "" {
		http.Error(w, "filename required", http.StatusBadRequest)
		return
	}

	path := filepath.Join("uploads", filename)
	f, err := os.Open(path)
	if err != nil {
		http.Error(w, fmt.Sprintf("file not found: %v", err), http.StatusNotFound)
		return
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		http.Error(w, "unable to read file info", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeContent(w, r, filename, info.ModTime(), f)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds struct {
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if shared.SState.Password != "" && creds.Password != shared.SState.Password {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status":"ok"}`)
}
