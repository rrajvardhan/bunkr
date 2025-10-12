package upload

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
	"github.com/rrajvardhan/bunkr/internal/tui/shared/style"
)

type State struct {
	fp           filepicker.Model
	selectedFile string
	err          error
	status       string
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg { return clearErrorMsg{} })
}

func Start() State {
	m := State{
		fp: NewFilePicker(),
	}

	return m
}

func (m State) Init() tea.Cmd { return m.fp.Init() }

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg { return shared.NavigateMsg{Target: shared.Dashboard} }
		}
	case clearErrorMsg:
		m.err = nil

	case shared.UploadSuccessMsg:
		m.err = nil
		m.selectedFile = ""
		m.status = fmt.Sprintf("uploaded %s successfully!", filepath.Base(msg.FileName))

	case shared.UploadErrorMsg:
		m.err = msg.Err
		return m, clearErrorAfter(3 * time.Second)
	}

	var cmd tea.Cmd
	m.fp, cmd = m.fp.Update(msg)

	if didSelect, path := m.fp.DidSelectFile(msg); didSelect {
		m.selectedFile = path
	}

	if didSelect, path := m.fp.DidSelectFile(msg); didSelect {
		m.selectedFile = path
		return m, UploadCmd(path)
	}

	return m, cmd
}

func (m State) View() string {
	header := style.Subtle.Render(shared.SState.URL)

	body := m.fp.View()

	divider := style.Divider.Render(strings.Repeat("─", 74))
	controls := Info.Render("[Esc] Back to Dashboard")

	status := ""
	if m.err != nil {
		status = style.Error.Render(m.err.Error())
	} else if info := m.status; info != "" {
		status = Succes.Render(info)
	}

	content := fmt.Sprintf(
		"%s\n\n%s\n%s\n%s\n%s",
		header,
		body,
		status,
		divider,
		controls,
	)

	return style.TermCenter(style.BoxStyle.Render(content))
}
