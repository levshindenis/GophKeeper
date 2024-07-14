package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	input "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"github.com/levshindenis/GophKeeper/internal/app/storages/cloud"
	"github.com/levshindenis/GophKeeper/internal/app/storages/server_database"
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
	choices        []string
	nextState      map[string]string
	currentChoices map[string][]string
	textInput      input.Model
	db             server_database.ServerDatabase
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
		textInput:      input.New(),
		nextState:      make(map[string]string),
		currentChoices: make(map[string][]string),
		err:            models.TeaErr{},
		textItem:       models.Text{},
		fileItem:       models.File{},
		cardItem:       models.Card{},
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
	newModel.nextState["add_file_comment"] = "add_file"
	newModel.nextState["change_text_name"] = "change_text_description"
	newModel.nextState["change_text_description"] = "change_text_comment"
	newModel.nextState["change_text_comment"] = "change_text"
	newModel.nextState["change_card_name"] = "change_card_number"
	newModel.nextState["change_card_number"] = "change_card_month"
	newModel.nextState["change_card_month"] = "change_card_year"
	newModel.nextState["change_card_year"] = "change_card_cvv"
	newModel.nextState["change_card_cvv"] = "change_card_owner"
	newModel.nextState["change_card_owner"] = "change_card_comment"
	newModel.nextState["change_card_comment"] = "change_card"
	newModel.nextState["change_file_name"] = "change_file"

	newModel.currentChoices["start"] = []string{"Регистрация", "Вход", "Выйти из программы"}
	newModel.currentChoices["repeat"] = []string{"Да", "Нет"}
	newModel.currentChoices["menu"] = []string{"Показать все записи", "Показать все файлы", "Показать все карты",
		"Посмотреть избранное", "Сменить пользователя", "Выйти из программы"}

	if err := tools.MakeBaseDirectories(); err != nil {
		log.Fatalf(err.Error())
	}

	db, err := sql.Open("sqlite3", "/tmp/keeper/db/keeper.db")
	if err != nil {
		log.Fatalf(err.Error())
	}

	sd := server_database.ServerDatabase{DB: db}

	if err = sd.MakeTables(); err != nil {
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

	if m.state == "repeat" {
		return m.RepeatUpdate(msg)
	}

	if m.state == "menu" {
		return m.MenuUpdate(msg)
	}

	if ok := slices.Contains([]string{"texts", "files", "cards"}, m.state); ok {
		return m.ListsUpdate(msg)
	}

	if strings.Contains(m.state, "view") {
		return m.ItemUpdate(msg)
	}

	if m.state == "favourites" {
		return m.FavouritesUpdate(msg)
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
