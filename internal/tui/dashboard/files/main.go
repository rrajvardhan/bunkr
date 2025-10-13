package files

import (
	"fmt"
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
	items := loadFiles()
	m := State{list: list.New(items, list.NewDefaultDelegate(), 50, 1)}
	m.list.Title = ""
	m.list.Styles.Title = style.Header
	m.list.SetShowHelp(false)
	m.list.SetHeight(20)

	return m
}

func (m State) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg { return shared.NavigateMsg{Target: shared.Dashboard} }

		case "enter":
			selected, ok := m.list.SelectedItem().(fileItem)
			if !ok {
				break
			}

			path := filepath.Join("uploads", selected.title)
			if err := os.Remove(path); err != nil {
				m.err = err
				return m, clearAfter(3 * time.Second)
			}

			m.status = fmt.Sprintf("deleted %s", selected.title)
			m.list.SetItems(loadFiles())
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
	controls := Info.Render("[Enter] Delete  [Esc] Back to Dashboard")

	status := ""
	if m.err != nil {
		status = style.Error.Render(m.err.Error())
	} else if m.status != "" {
		status = Succes.Render(m.status)
	}

	content := fmt.Sprintf(
		"%s%s\n\n%s\n%s\n%s",
		header,
		body,
		status,
		divider,
		controls,
	)

	return style.TermCenter(style.BoxStyle.Render(content))
}

func loadFiles() []list.Item {
	files, err := os.ReadDir("uploads")
	if err != nil {
		return []list.Item{}
	}

	items := []list.Item{}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		info, err := f.Info()
		desc := "unknown size"
		if err == nil {
			desc = fmt.Sprintf("%.1f KB", float64(info.Size())/1024)
		}
		items = append(items, fileItem{
			title: f.Name(),
			desc:  fmt.Sprintf("%s", desc),
		})
	}
	return items
}

func clearAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(_ time.Time) tea.Msg { return clearMsg{} })
}
