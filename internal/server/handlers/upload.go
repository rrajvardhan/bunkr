package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rrajvardhan/bunkr/internal/crypto"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not get uploaded file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not read uploaded file: %v", err), http.StatusInternalServerError)
		return
	}

	passphrase := shared.SState.Password

	var blob []byte
	if passphrase != "" {
		blob, err = crypto.EncryptBytes(passphrase, data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not encrypt file: %v", err), http.StatusInternalServerError)
			return
		}
	} else {
		blob = data
	}

	dstPath := "./uploads/" + handler.Filename
	if err := os.MkdirAll("./uploads", 0755); err != nil {
		http.Error(w, fmt.Sprintf("Could not create upload directory: %v", err), http.StatusInternalServerError)
		return
	}

	if err := os.WriteFile(dstPath, blob, 0600); err != nil {
		http.Error(w, fmt.Sprintf("Could not save file: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File %s uploaded successfully", handler.Filename)
}
