package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func AuthenticateWithServer(url, password string) error {
	reqBody := map[string]string{"password": password}
	jsonBody, _ := json.Marshal(reqBody)

	resp, err := http.Post(url+"/login", "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to reach server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("login failed: %s", string(body))
	}

	return nil
}
