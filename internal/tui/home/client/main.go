package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/rrajvardhan/bunkr/internal/tui/shared"
	"github.com/rrajvardhan/bunkr/internal/tui/shared/style"
)

type fileItem struct {
	title, desc string
}

func (f fileItem) Title() string       { return f.title }
func (f fileItem) Description() string { return f.desc }
func (f fileItem) FilterValue() string { return f.title }

type clearMsg struct{}

type State struct {
	list   list.Model
	status string
	err    error
}

func Start() State {
	items, err := loadServerFiles()
	m := State{
		list: list.New(items, list.NewDefaultDelegate(), 60, 1),
	}
	m.list.Title = ""
	m.list.SetHeight(20)
	m.list.SetShowHelp(false)
	m.list.SetShowTitle(false)

	if err != nil {
		m.err = err
		m.list.SetItems(nil) // hide list if error
	}

	return m
}

func (m State) Init() tea.Cmd { return nil }

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg { return shared.NavigateMsg{Target: shared.Home} }

		case "r":
			items, err := loadServerFiles()
			if err != nil {
				m.err = err
				m.list.SetItems(nil)
			} else {
				m.err = nil
				m.list.SetItems(items)
				m.status = "refreshed"
			}
			return m, clearAfter(1 * time.Second)

		case "enter":
			selected, ok := m.list.SelectedItem().(fileItem)
			if !ok {
				break
			}

			fileURL := fmt.Sprintf("%s/download/%s", shared.SState.URL, selected.title)
			savePath := filepath.Join("downloads", selected.title)

			if err := downloadFile(fileURL, savePath); err != nil {
				m.err = err
				return m, clearAfter(3 * time.Second)
			}

			m.status = fmt.Sprintf("downloaded %s", selected.title)
			return m, clearAfter(2 * time.Second)
		}

	case clearMsg:
		m.err = nil
		m.status = ""
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m State) View() string {
	header := style.Subtle.Render(shared.SState.URL)
	body := m.list.View()

	divider := style.Divider.Render(strings.Repeat("─", 74))
	controls := Info.Render("[Enter] Download  [R] Refresh  [Esc] Back to Dashboard")

	status := ""
	if m.err != nil {
		status = style.Error.Render(m.err.Error())
	} else if m.status != "" {
		status = Succes.Render(m.status)
	}

	content := fmt.Sprintf(
		"%s\n%s\n\n%s\n%s\n%s",
		header,
		body,
		status,
		divider,
		controls,
	)

	return style.TermCenter(style.BoxStyle.Render(content))
}

func loadServerFiles() ([]list.Item, error) {
	url := fmt.Sprintf("%s/files", shared.SState.URL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch files: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %s", resp.Status)
	}

	var serverFiles []struct {
		Name string  `json:"name"`
		Size float64 `json:"size_kb"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&serverFiles); err != nil {
		return nil, fmt.Errorf("invalid server response: %w", err)
	}

	items := make([]list.Item, 0, len(serverFiles))
	for _, f := range serverFiles {
		items = append(items, fileItem{
			title: f.Name,
			desc:  fmt.Sprintf("%.1f KB", f.Size),
		})
	}

	return items, nil
}

func downloadFile(url, savePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned %s", resp.Status)
	}

	if err := os.MkdirAll(filepath.Dir(savePath), 0755); err != nil {
		return err
	}

	out, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func clearAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(_ time.Time) tea.Msg { return clearMsg{} })
}
