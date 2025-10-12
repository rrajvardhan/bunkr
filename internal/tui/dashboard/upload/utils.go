package upload

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
)

func NewFilePicker() filepicker.Model {
	picker := filepicker.New()

	picker.CurrentDirectory, _ = os.UserHomeDir()
	picker.AllowedTypes = []string{}
	picker.ShowPermissions = false
	picker.SetHeight(10)

	picker.Styles.Directory = Directory
	picker.Styles.Selected = Selected
	picker.Styles.Cursor = Cursor
	picker.Styles.File = File

	return picker
}

func UploadFile(filePath, uploadURL string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	writer.Close()

	req, err := http.NewRequest("POST", uploadURL, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("upload failed: %s", resp.Status)
	}

	return nil
}

func UploadCmd(filePath string) tea.Cmd {
	return func() tea.Msg {
		err := UploadFile(filePath, shared.SState.URL+"/upload")
		if err != nil {
			return shared.UploadErrorMsg{Err: err}
		}
		return shared.UploadSuccessMsg{FileName: filePath}
	}
}
