package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

var (
	storage Database
)

func TestMain(m *testing.M) {
	if _, err := os.Create("./test.db"); err != nil {
		log.Fatalf("Create: %s", err.Error())
	}

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatalf("Open: %s", err.Error())
	}

	storage = Database{DB: db}

	exitCode := m.Run()

	storage.Close()

	if err = os.Remove("./test.db"); err != nil {
		log.Fatalf(err.Error())
	}

	os.Exit(exitCode)
}

func TestDatabase_MakeTables(t *testing.T) {
	var count int
	tests := []struct {
		name    string
		param   string
		wantErr bool
	}{
		{
			name:    "Make Local DB",
			param:   "",
			wantErr: true,
		},
		{
			name:    "Make Server DB",
			param:   "server",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if err := storage.MakeTables(tt.param); err != nil {
				assert.Fail(t, "Err with make tables")
			}
			err := storage.DB.QueryRowContext(ctx, `Select count(*) from users`).Scan(&count)
			assert.Equal(t, tt.wantErr, err != nil)

			err = storage.DB.QueryRowContext(ctx, `Select count(*) from updates`).Scan(&count)
			assert.NoError(t, err)

			err = storage.DB.QueryRowContext(ctx, `Select count(*) from texts`).Scan(&count)
			assert.NoError(t, err)

			err = storage.DB.QueryRowContext(ctx, `Select count(*) from binaries`).Scan(&count)
			assert.NoError(t, err)

			err = storage.DB.QueryRowContext(ctx, `Select count(*) from cards`).Scan(&count)
			assert.NoError(t, err)

			if tt.wantErr {
				_, err = storage.DB.ExecContext(ctx, `DROP TABLE updates`)
				assert.NoError(t, err)

				_, err = storage.DB.ExecContext(ctx, `DROP TABLE texts`)
				assert.NoError(t, err)

				_, err = storage.DB.ExecContext(ctx, `DROP TABLE binaries`)
				assert.NoError(t, err)

				_, err = storage.DB.ExecContext(ctx, `DROP TABLE cards`)
				assert.NoError(t, err)

			}
		})
	}
}

func TestDatabase_AddUser(t *testing.T) {
	var helpStr string

	tests := []struct {
		name     string
		login    string
		password string
		word     string
		repeat   bool
	}{
		{
			name:     "Good register",
			login:    "Dima",
			password: "12345",
			word:     "cat",
			repeat:   false,
		},
		{
			name:     "Repeat login",
			login:    "Dima",
			password: "98765",
			word:     "dog",
			repeat:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, flag, _ := storage.AddUser(models.Register{Login: tt.login, Password: tt.password, Word: tt.word})
			assert.Equal(t, flag, tt.repeat)

			if !tt.repeat {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				err := storage.DB.QueryRowContext(ctx,
					`select password from users where login = $1`, tt.login).Scan(&helpStr)
				assert.NoError(t, err)
			}
		})
	}
}

func TestDatabase_AddUpdateTime(t *testing.T) {
	var helpStr string

	tests := []struct {
		name    string
		login   string
		updTime string
	}{
		{
			name:    "Good update",
			login:   "Dima",
			updTime: time.Now().Format(time.RFC3339),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.AddUpdateTime(tt.login, tt.updTime)
			assert.NoError(t, err)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			err = storage.DB.QueryRowContext(ctx,
				`select update_time from updates where user_id = $1`, tt.login).Scan(&helpStr)
			assert.NoError(t, err)
		})
	}
}

func TestDatabase_AddTexts(t *testing.T) {
	var helpStr string

	tests := []struct {
		name  string
		login string
		items []models.Text
	}{
		{
			name:  "Good add texts",
			login: "Dima",
			items: []models.Text{
				{Name: "First", Description: "First", Comment: "First", Favourite: "Да"},
				{Name: "Second", Description: "Second", Comment: "Second", Favourite: "Нет"},
				{Name: "Third", Description: "Third", Comment: "Third", Favourite: "Нет"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.AddTexts(tt.login, tt.items)
			assert.NoError(t, err)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			for i := range tt.items {
				err = storage.DB.QueryRowContext(ctx,
					`select description from texts where user_id = $1 and name = $2`,
					tt.login, tt.items[i].Name).Scan(&helpStr)
				assert.NoError(t, err)
			}
		})
	}
}

func TestDatabase_AddFiles(t *testing.T) {
	var helpStr string

	tests := []struct {
		name  string
		login string
		items []models.File
	}{
		{
			name:  "Good add files",
			login: "Dima",
			items: []models.File{
				{Name: "First", Comment: "First", Favourite: "Да"},
				{Name: "Second", Comment: "Second", Favourite: "Нет"},
				{Name: "Third", Comment: "Third", Favourite: "Нет"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.AddFiles(tt.login, tt.items)
			assert.NoError(t, err)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			for i := range tt.items {
				err = storage.DB.QueryRowContext(ctx,
					`select comment from binaries where user_id = $1 and name = $2`,
					tt.login, tt.items[i].Name).Scan(&helpStr)
				assert.NoError(t, err)
			}
		})
	}
}

func TestDatabase_AddCards(t *testing.T) {
	var helpStr string

	tests := []struct {
		name  string
		login string
		items []models.Card
	}{
		{
			name:  "Good add cards",
			login: "Dima",
			items: []models.Card{
				{Bank: "First", Number: "1111 1111", Month: "1", Year: "1", CVV: "111",
					Owner: "First", Comment: "First", Favourite: "Да"},
				{Bank: "Second", Number: "2222 2222", Month: "2", Year: "2", CVV: "222",
					Owner: "Second", Comment: "Second", Favourite: "Нет"},
				{Bank: "Third", Number: "3333 3333", Month: "3", Year: "3", CVV: "333",
					Owner: "Third", Comment: "Third", Favourite: "Нет"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.AddCards(tt.login, tt.items)
			assert.NoError(t, err)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			for i := range tt.items {
				err = storage.DB.QueryRowContext(ctx,
					`select comment from cards where user_id = $1 and bank = $2 and number = $3`,
					tt.login, tt.items[i].Bank, tt.items[i].Number).Scan(&helpStr)
				assert.NoError(t, err)
			}
		})
	}
}

func TestDatabase_GetText(t *testing.T) {
	tests := []struct {
		name  string
		login string
		text  string
		flag  bool
	}{
		{
			name:  "Good Get Text",
			login: "Dima",
			text:  "Second",
			flag:  false,
		},
		{
			name:  "Bad login",
			login: "Dimas",
			text:  "Second",
			flag:  true,
		},
		{
			name:  "Bad value",
			login: "Dima",
			text:  "Seconddd",
			flag:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := storage.GetText(tt.login, tt.text)
			assert.Equal(t, tt.flag, err != nil)
		})
	}
}

func TestDatabase_GetFile(t *testing.T) {
	tests := []struct {
		name  string
		login string
		text  string
		flag  bool
	}{
		{
			name:  "Good Get File",
			login: "Dima",
			text:  "Second",
			flag:  false,
		},
		{
			name:  "Bad login",
			login: "Dimas",
			text:  "Second",
			flag:  true,
		},
		{
			name:  "Bad value",
			login: "Dima",
			text:  "Seconddd",
			flag:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := storage.GetFile(tt.login, tt.text)
			assert.Equal(t, tt.flag, err != nil)
		})
	}
}

func TestDatabase_GetCard(t *testing.T) {
	tests := []struct {
		name  string
		login string
		text  string
		flag  bool
	}{
		{
			name:  "Good Get Text",
			login: "Dima",
			text:  "Second///2222 2222",
			flag:  false,
		},
		{
			name:  "Bad login",
			login: "Dimas",
			text:  "Second///2222 2222",
			flag:  true,
		},
		{
			name:  "Bad value",
			login: "Dima",
			text:  "Seconddd///2222 2222",
			flag:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := storage.GetCard(tt.login, tt.text)
			assert.Equal(t, tt.flag, err != nil)
		})
	}
}

func TestDatabase_GetUserTexts(t *testing.T) {
	tests := []struct {
		name   string
		login  string
		flag   bool
		length int
	}{
		{
			name:   "Good Get User Texts",
			login:  "Dima",
			flag:   false,
			length: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items, err := storage.GetUserTexts(tt.login)
			assert.Equal(t, tt.flag, err != nil)

			assert.Equal(t, len(items), tt.length)
		})
	}
}

func TestDatabase_GetUserFiles(t *testing.T) {
	tests := []struct {
		name   string
		login  string
		flag   bool
		length int
	}{
		{
			name:   "Good Get User Files",
			login:  "Dima",
			flag:   false,
			length: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items, err := storage.GetUserFiles(tt.login)
			assert.Equal(t, tt.flag, err != nil)

			assert.Equal(t, len(items), tt.length)
		})
	}
}

func TestDatabase_GetUserCards(t *testing.T) {
	tests := []struct {
		name   string
		login  string
		flag   bool
		length int
	}{
		{
			name:   "Good Get User Cards",
			login:  "Dima",
			flag:   false,
			length: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items, err := storage.GetUserCards(tt.login)
			assert.Equal(t, tt.flag, err != nil)

			assert.Equal(t, len(items), tt.length)
		})
	}
}

func TestDatabase_ListTexts(t *testing.T) {
	tests := []struct {
		name   string
		login  string
		flag   bool
		length int
	}{
		{
			name:   "Good List Texts",
			login:  "Dima",
			flag:   false,
			length: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items, err := storage.ListTexts(tt.login)
			assert.Equal(t, tt.flag, err != nil)

			assert.Equal(t, len(items), tt.length)
		})
	}
}
func TestDatabase_ListFiles(t *testing.T) {
	tests := []struct {
		name   string
		login  string
		flag   bool
		length int
	}{
		{
			name:   "Good List Files",
			login:  "Dima",
			flag:   false,
			length: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items, err := storage.ListFiles(tt.login)
			assert.Equal(t, tt.flag, err != nil)

			assert.Equal(t, len(items), tt.length)
		})
	}
}
func TestDatabase_ListCards(t *testing.T) {
	tests := []struct {
		name   string
		login  string
		flag   bool
		length int
	}{
		{
			name:   "Good List Cards",
			login:  "Dima",
			flag:   false,
			length: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items, err := storage.ListCards(tt.login)
			assert.Equal(t, tt.flag, err != nil)

			assert.Equal(t, len(items), tt.length)
		})
	}
}

func TestDatabase_ListFavourites(t *testing.T) {
	tests := []struct {
		name   string
		login  string
		flag   bool
		length int
	}{
		{
			name:   "Good List Favourites",
			login:  "Dima",
			flag:   false,
			length: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items, err := storage.ListFavourites(tt.login)
			assert.Equal(t, tt.flag, err != nil)

			assert.Equal(t, len(items), tt.length)
		})
	}
}

func TestDatabase_ChangeText(t *testing.T) {
	var helpStr string

	tests := []struct {
		name  string
		login string
		items models.ChText
	}{
		{
			name:  "Good change text",
			login: "Dima",
			items: models.ChText{OldName: "First", NewName: "First_F",
				NewDescription: "First_F", NewComment: "First_F", NewFavourite: "Да"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.ChangeText(tt.login, tt.items)
			assert.NoError(t, err)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			err = storage.DB.QueryRowContext(ctx,
				`select description from texts where user_id = $1 and name = $2`,
				tt.login, tt.items.NewName).Scan(&helpStr)
			assert.NoError(t, err)
		})
	}
}

func TestDatabase_ChangeFile(t *testing.T) {
	var helpStr string

	tests := []struct {
		name  string
		login string
		items models.ChFile
	}{
		{
			name:  "Good change file",
			login: "Dima",
			items: models.ChFile{OldName: "First", NewName: "First_F", NewComment: "First_F", NewFavourite: "Да"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.ChangeFile(tt.login, tt.items)
			assert.NoError(t, err)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			err = storage.DB.QueryRowContext(ctx,
				`select comment from binaries where user_id = $1 and name = $2`,
				tt.login, tt.items.NewName).Scan(&helpStr)
			assert.NoError(t, err)
		})
	}
}

func TestDatabase_ChangeCard(t *testing.T) {
	var helpStr string

	tests := []struct {
		name  string
		login string
		items models.ChCard
	}{
		{
			name:  "Good change card",
			login: "Dima",
			items: models.ChCard{OldBank: "First", OldNumber: "1111 1111", NewBank: "First_F",
				NewNumber: "1111 1111 1111", NewMonth: "11", NewYear: "11", NewCVV: "1111",
				NewOwner: "First_F", NewComment: "First_F", NewFavourite: "Да"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.ChangeCard(tt.login, tt.items)
			assert.NoError(t, err)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			err = storage.DB.QueryRowContext(ctx,
				`select comment from cards where user_id = $1 and bank = $2 and number = $3`,
				tt.login, tt.items.NewBank, tt.items.NewNumber).Scan(&helpStr)
			assert.NoError(t, err)
		})
	}
}

func TestDatabase_CheckCookie(t *testing.T) {
	tests := []struct {
		name   string
		login  string
		cookie string
		flag   bool
	}{
		{
			name:   "Good check cookie",
			login:  "Dima",
			cookie: "",
			flag:   true,
		},
		{
			name:   "Bad check cookie",
			login:  "",
			cookie: "hgnglfgnlkmlfng",
			flag:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			if tt.cookie == "" {
				err := storage.DB.QueryRowContext(ctx, `select user_id from users where login = $1`,
					tt.login).Scan(&tt.cookie)
				assert.NoError(t, err)
			}

			assert.Equal(t, storage.CheckCookie(tt.cookie), tt.flag)
		})
	}
}

func TestDatabase_CheckUser(t *testing.T) {
	tests := []struct {
		name     string
		login    string
		password string
		flag     bool
	}{
		{
			name:     "Good check user",
			login:    "Dima",
			password: "12345",
			flag:     true,
		},
		{
			name:     "Bad check user",
			login:    "Dima",
			password: "54321",
			flag:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, storage.CheckUser(models.Login{Login: tt.login, Password: tt.password}), tt.flag)
		})
	}
}

func TestDatabase_DeleteTexts(t *testing.T) {
	var count int

	tests := []struct {
		name  string
		login string
		items []string
		param string
	}{
		{
			name:  "Good delete one text",
			login: "Dima",
			items: []string{"Second"},
			param: "",
		},
		{
			name:  "Good delete all texts",
			login: "Dima",
			items: nil,
			param: "all",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.DeleteTexts(tt.login, tt.items, tt.param)
			assert.NoError(t, err)

			if tt.param == "all" {
				ctx, cancel := context.WithCancel(context.Background())

				defer cancel()

				err = storage.DB.QueryRowContext(ctx, `select count(*) from texts`).Scan(&count)
				assert.NoError(t, err)

				assert.Equal(t, count, 0)
			}
		})
	}
}

func TestDatabase_DeleteFiles(t *testing.T) {
	var count int

	tests := []struct {
		name  string
		login string
		items []string
		param string
	}{
		{
			name:  "Good delete one file",
			login: "Dima",
			items: []string{"Second"},
			param: "",
		},
		{
			name:  "Good delete all files",
			login: "Dima",
			items: nil,
			param: "all",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.DeleteFiles(tt.login, tt.items, tt.param)
			assert.NoError(t, err)

			if tt.param == "all" {
				ctx, cancel := context.WithCancel(context.Background())

				defer cancel()

				err = storage.DB.QueryRowContext(ctx, `select count(*) from binaries`).Scan(&count)
				assert.NoError(t, err)

				assert.Equal(t, count, 0)
			}
		})
	}
}

func TestDatabase_DeleteCards(t *testing.T) {
	var count int

	tests := []struct {
		name  string
		login string
		items []string
		param string
	}{
		{
			name:  "Good delete one text",
			login: "Dima",
			items: []string{"Second///2222 2222"},
			param: "",
		},
		{
			name:  "Good delete all",
			login: "Dima",
			items: nil,
			param: "all",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.DeleteCards(tt.login, tt.items, tt.param)
			assert.NoError(t, err)

			if tt.param == "all" {
				ctx, cancel := context.WithCancel(context.Background())

				defer cancel()

				err = storage.DB.QueryRowContext(ctx, `select count(*) from cards`).Scan(&count)
				assert.NoError(t, err)

				assert.Equal(t, count, 0)
			}
		})
	}
}

func TestDatabase_GetLogin(t *testing.T) {
	tests := []struct {
		name   string
		login  string
		cookie string
		result string
		flag   bool
	}{
		{
			name:   "Good Get Login",
			login:  "Dima",
			cookie: "",
			result: "Dima",
			flag:   false,
		},
		{
			name:   "Bad Get Login",
			login:  "Dimas",
			cookie: "",
			result: "",
			flag:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			storage.DB.QueryRowContext(ctx, `select user_id from users where login = $1`, tt.login).Scan(&tt.cookie)

			helpStr, err := storage.GetLogin(tt.cookie)
			assert.Equal(t, tt.flag, err != nil)

			assert.Equal(t, helpStr, tt.result)

		})
	}
}

func TestDatabase_GetUpdateTime(t *testing.T) {
	tests := []struct {
		name  string
		login string
		flag  bool
	}{
		{
			name:  "Good Get Update Time",
			login: "Dima",
			flag:  false,
		},
		{
			name:  "Bad Get Update Time",
			login: "Dimas",
			flag:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := storage.GetUpdateTime(tt.login)
			assert.Equal(t, tt.flag, err != nil)
		})
	}
}

func TestDatabase_SetCookie(t *testing.T) {
	var value string

	tests := []struct {
		name   string
		login  string
		cookie string
		flag   bool
	}{
		{
			name:   "Good Set Cookie",
			login:  "Dima",
			cookie: "abc",
			flag:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.SetCookie(tt.cookie, tt.login)
			assert.Equal(t, tt.flag, err != nil)

			ctx, cancel := context.WithCancel(context.Background())

			defer cancel()

			err = storage.DB.QueryRowContext(ctx, `select user_id from users where login = $1`, tt.login).Scan(&value)
			assert.Equal(t, tt.flag, err != nil)

			assert.Equal(t, value, tt.cookie)
		})
	}
}

func TestDatabase_SetUpdateTime(t *testing.T) {
	var value string

	tests := []struct {
		name    string
		login   string
		updTime string
		flag    bool
	}{
		{
			name:    "Good Set Update Time",
			login:   "Dima",
			updTime: "abc",
			flag:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.SetUpdateTime(tt.login, tt.updTime)
			assert.Equal(t, tt.flag, err != nil)

			ctx, cancel := context.WithCancel(context.Background())

			defer cancel()

			err = storage.DB.QueryRowContext(ctx, `select update_time from updates where user_id = $1`, tt.login).Scan(&value)
			assert.Equal(t, tt.flag, err != nil)

			assert.Equal(t, value, tt.updTime)
		})
	}
}
