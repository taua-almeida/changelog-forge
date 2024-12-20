package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/rand"
)

const listHeight = 9
const leftPadding = 1

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(0)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(leftPadding)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(leftPadding)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(leftPadding)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(leftPadding).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 0, 2)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list     list.Model
	version  string
	step     int
	quitting bool
	date     string
}

func initialModel() model {
	items := []list.Item{
		item("major"),
		item("minor"),
		item("patch"),
	}

	const defaultWidth = 20
	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Select the version type"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return model{
		list: l,
		step: 0,
		date: time.Now().Format("2006-01-02"),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			if m.step == 0 { // Step 0: Select version
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.version = string(i)
				}
				m.step = 1
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return quitTextStyle.Render("Exiting program.")
	}
	if m.step == 0 {
		return "\n" + m.list.View()
	}
	return ""
}

var adjectives = []string{
	"brilliant", "lazy", "energetic", "sleepy", "grumpy",
	"cheerful", "mighty", "tiny", "fancy", "witty",
}

var nouns = []string{
	"butterfly", "unicorn", "kitten", "wizard", "robot",
	"phoenix", "dragon", "penguin", "octopus", "yeti",
}

func randomName() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	adj := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	return fmt.Sprintf("%s-%s", adj, noun)
}

func GenerateJSON() {
	// Ensure the .changeset directory exists
	changesetDir := ".changeset"
	if _, err := os.Stat(changesetDir); os.IsNotExist(err) {
		if err := os.Mkdir(changesetDir, 0755); err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			os.Exit(1)
		}
	}

	// Run the Bubble Tea program for selecting version type
	p := tea.NewProgram(initialModel())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}

	m := finalModel.(model)

	if m.quitting {
		fmt.Println("Program exited without generating JSON.")
		os.Exit(0)
	}

	// Prompt user for descriptions
	fmt.Print("Enter descriptions (comma-separated): ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	if input == "\n" || input == "" {
		fmt.Println("No descriptions entered. Exiting program.")
		os.Exit(0)
	}

	// Process descriptions
	descriptions := strings.Split(input, ",")
	for i, desc := range descriptions {
		descriptions[i] = strings.TrimSpace(desc)
	}

	// Create JSON object
	entry := struct {
		Version      string   `json:"version"`
		Date         string   `json:"date"`
		Descriptions []string `json:"descriptions"`
	}{
		Version:      m.version,
		Date:         m.date,
		Descriptions: descriptions,
	}

	// Write to file
	fileName := fmt.Sprintf("%s.json", randomName())
	filePath := filepath.Join(changesetDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(entry); err != nil {
		fmt.Printf("Error writing JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nchangelog.json created successfully!")
}
