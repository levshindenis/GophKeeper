package main

import (
	"database/sql"
	"fmt"
	"github.com/levshindenis/GophKeeper/internal/app/consts"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"log"
	"os"
	"slices"
	"strings"

	input "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-resty/resty/v2"

	"github.com/levshindenis/GophKeeper/internal/app/cloud"
	"github.com/levshindenis/GophKeeper/internal/app/database"
	"github.com/levshindenis/GophKeeper/internal/app/tools"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatalf(err.Error())
		os.Exit(1)
	}

	p.Quit()
}

type model struct {
	cursor         int
	state          string
	helpStr        string
	userId         string
	secretKey      string
	choices        []string
	encChoices     []string
	nextState      map[string]string
	currentChoices map[string][]string
	client         *resty.Client
	textInput      input.Model
	db             database.Database
	err            models.TeaErr
	cloud          cloud.Cloud
	textItem       models.Text
	fileItem       models.File
	cardItem       models.Card
}

func initialModel() model {
	var (
		cl cloud.Cloud
	)

	newModel := model{
		choices:        []string{"Регистрация", "Вход", "Выйти из программы"},
		cursor:         0,
		state:          "start",
		helpStr:        "",
		userId:         "",
		secretKey:      "",
		encChoices:     []string{},
		textInput:      input.New(),
		nextState:      consts.GetNextStatesMap(),
		currentChoices: consts.GetStatesMap(),
		client:         resty.New(),
		err:            models.TeaErr{},
		textItem:       models.Text{},
		fileItem:       models.File{},
		cardItem:       models.Card{},
	}

	newModel.textInput.Focus()

	if err := tools.MakeBaseDirectories(); err != nil {
		log.Fatalf(err.Error())
	}

	db, err := sql.Open("sqlite3", "/tmp/keeper/db/keeper.db")
	if err != nil {
		log.Fatalf(err.Error())
	}

	sd := database.Database{DB: db}

	if err = sd.MakeTables(""); err != nil {
		log.Fatalf(err.Error())
	}

	if err = cl.Init(); err != nil {
		log.Fatalf(err.Error())
	}

	newModel.db = sd
	newModel.cloud = cl

	return newModel
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if _, ok := m.nextState[m.state]; ok {
		return m.InputUpdate(msg)
	}

	switch {
	case m.state == "repeat":
		return m.RepeatUpdate(msg)
	case m.state == "menu":
		return m.MenuUpdate(msg)
	case m.state == "favourites":
		return m.FavouritesUpdate(msg)
	case slices.Contains([]string{"texts", "files", "cards"}, m.state):
		return m.ListsUpdate(msg)
	case strings.Contains(m.state, "view"):
		return m.ItemUpdate(msg)

	}

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			switch m.cursor {
			case 0:
				m.state = "reg_input_login"
				return m, nil
			case 1:
				m.cursor = 0
				m.state = "log_input_login"
				return m, nil
			case 2:
				if err := m.db.Close(); err != nil {
					log.Fatalf(err.Error())
				}
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if _, ok := m.nextState[m.state]; ok {
		return m.InputView()
	}

	if m.state == "repeat" {
		return m.RepeatView()
	}

	s := ""
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	return s

}
