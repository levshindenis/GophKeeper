package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"slices"

	input "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"github.com/levshindenis/GophKeeper/internal/app/storages/server_database"
	"github.com/levshindenis/GophKeeper/internal/app/tools"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := tools.MakeBaseDirectories(); err != nil {
		log.Fatalf(err.Error())
	}

	db, err := sql.Open("sqlite3", "/tmp/keeper/db/keeper.db")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	p := tea.NewProgram(initialModel(server_database.ServerDatabase{DB: db}))
	if _, err = p.Run(); err != nil {
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
	choices        []string
	nextState      map[string]string
	currentChoices map[string][]string
	textInput      input.Model
	db             server_database.ServerDatabase
	err            models.TeaErr
}

func initialModel(db server_database.ServerDatabase) model {
	newModel := model{
		choices:        []string{"Регистрация", "Вход", "Выйти из программы"},
		cursor:         0,
		state:          "start",
		helpStr:        "",
		textInput:      input.New(),
		nextState:      make(map[string]string),
		currentChoices: make(map[string][]string),
		db:             db,
	}
	newModel.textInput.Focus()
	newModel.nextState["reg_input_login"] = "reg_input_password"
	newModel.nextState["reg_input_password"] = "reg_input_word"
	newModel.nextState["reg_input_word"] = "registration"
	newModel.nextState["log_input_login"] = "log_input_password"
	newModel.nextState["log_input_password"] = "login"
	newModel.nextState["add_text_name"] = "add_text_description"
	newModel.nextState["add_text_description"] = "add_text_comment"
	newModel.nextState["add_text_comment"] = "add_text"
	newModel.nextState["add_card_name"] = "add_card_number"
	newModel.nextState["add_card_number"] = "add_card_month"
	newModel.nextState["add_card_month"] = "add_card_year"
	newModel.nextState["add_card_year"] = "add_card_cvv"
	newModel.nextState["add_card_cvv"] = "add_card_owner"
	newModel.nextState["add_card_owner"] = "add_card_comment"
	newModel.nextState["add_card_comment"] = "add_card"

	newModel.currentChoices["start"] = []string{"Регистрация", "Вход", "Выйти из программы"}
	newModel.currentChoices["repeat"] = []string{"Да", "Нет"}
	newModel.currentChoices["menu"] = []string{"Показать все тексты", "Показать все файлы", "Показать все карты",
		"Сменить пользователя", "Выйти из программы"}

	if err := db.MakeTables(); err != nil {
		panic(err)
	}

	return newModel
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if _, ok := m.nextState[m.state]; ok {
		return m.InputUpdate(msg)
	}

	if m.state == "repeat" {
		return m.RepeatUpdate(msg)
	}

	if m.state == "menu" {
		return m.MenuUpdate(msg)
	}

	if ok := slices.Contains([]string{"texts", "files", "cards"}, m.state); ok {
		return m.ListsUpdate(msg)
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
			if m.cursor == 0 {
				m.state = "reg_input_login"
				return m, nil
			}
			if m.cursor == 1 {
				m.cursor = 0
				m.state = "log_input_login"
				return m, nil
			}
			if m.cursor == 2 {
				m.cursor = 0
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
