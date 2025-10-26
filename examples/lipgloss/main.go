package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	gruvbox "github.com/mpstewart/go-gruvbox"
)

type entry struct {
	name string
	hex  string
}

type model struct {
	entries []entry
	cursor  int
}

func newModel() model {
	return model{
		entries: paletteEntries(gruvbox.Lipgloss()),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "j", "down":
			if len(m.entries) == 0 {
				return m, nil
			}
			m.cursor++
			if m.cursor >= len(m.entries) {
				m.cursor = len(m.entries) - 1
			}
		case "k", "up":
			if len(m.entries) == 0 {
				return m, nil
			}
			m.cursor--
			if m.cursor < 0 {
				m.cursor = 0
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if len(m.entries) == 0 {
		return "no colors"
	}

	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor >= len(m.entries) {
		m.cursor = len(m.entries) - 1
	}

	var b strings.Builder

	header := lipgloss.NewStyle().Bold(true).Render("Gruvbox Palette (j/k to navigate, q to exit)")
	b.WriteString(header)
	b.WriteRune('\n')

	table := m.tableView()
	detail := m.detailView()

	body := lipgloss.JoinHorizontal(lipgloss.Top, table, detail)
	b.WriteString(body)

	return b.String()
}

func (m model) tableView() string {
	nameStyle := lipgloss.NewStyle().Width(12).Align(lipgloss.Left).MarginRight(2).Faint(true)
	hexStyle := lipgloss.NewStyle().Width(10).Align(lipgloss.Left).MarginRight(2).Faint(true)
	swatchLabel := lipgloss.NewStyle().Faint(true).Render("swatch")

	headerRow := lipgloss.JoinHorizontal(lipgloss.Left,
		nameStyle.Render("name"),
		hexStyle.Render("hex"),
		swatchLabel,
	)

	var rows []string
	rows = append(rows, headerRow)

	rowNameStyle := lipgloss.NewStyle().Width(12).Align(lipgloss.Left).MarginRight(2)
	rowHexStyle := lipgloss.NewStyle().Width(10).Align(lipgloss.Left).MarginRight(2)
	swatchStyle := lipgloss.NewStyle().Padding(0, 4)
	selectedStyle := lipgloss.NewStyle().Bold(true)

	cursorStyle := lipgloss.NewStyle().Width(2).Align(lipgloss.Center).Faint(true)
	selectedCursorStyle := cursorStyle.Copy().Bold(true)

	for i, entry := range m.entries {
		currentCursor := cursorStyle.Render(" ")
		hexCell := rowHexStyle.Render(entry.hex)
		swatchCell := swatchStyle.Background(lipgloss.Color(entry.hex)).Render(" ")
		nameCell := rowNameStyle.Render(entry.name)

		if i == m.cursor {
			currentCursor = selectedCursorStyle.Render(">")
			lineBody := lipgloss.JoinHorizontal(lipgloss.Left, currentCursor, nameCell, hexCell, swatchCell)
			rows = append(rows, selectedStyle.Render(lineBody))
			continue
		}

		line := lipgloss.JoinHorizontal(lipgloss.Left, currentCursor, nameCell, hexCell, swatchCell)
		rows = append(rows, line)
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (m model) detailView() string {
	if m.cursor < 0 || m.cursor >= len(m.entries) {
		return ""
	}

	selected := m.entries[m.cursor]
	padding := lipgloss.NewStyle().PaddingLeft(4)

	title := lipgloss.NewStyle().Bold(true).Render("Details")
	name := lipgloss.NewStyle().Bold(true).Render(selected.name)
	hex := lipgloss.NewStyle().Faint(false).Render(selected.hex)
	swatch := lipgloss.NewStyle().
		Width(20).
		Height(5).
		Align(lipgloss.Center).
		Background(lipgloss.Color(selected.hex)).
		Render(" ")

	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		name,
		hex,
		swatch,
	)
	return padding.Render(content)
}

func paletteEntries(p gruvbox.LipglossPalette) []entry {
	val := reflect.ValueOf(p)
	typ := val.Type()

	entries := make([]entry, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		entries = append(entries, entry{
			name: field.Name,
			hex:  val.Field(i).String(),
		})
	}
	return entries
}

func main() {
	if err := tea.NewProgram(newModel()).Start(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
