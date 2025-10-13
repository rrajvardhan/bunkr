package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rrajvardhan/bunkr/internal/tui/shared"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	shared.SState.Password = creds.Password

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status":"ok"}`)
}
