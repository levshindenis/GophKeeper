package models

import "mime/multipart"

type Register struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Word     string `json:"word"`
}

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Text struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Comment     string `json:"comment"`
	Favourite   bool   `json:"favourite"`
}

type File struct {
	Name      string `json:"name"`
	Comment   string `json:"comment"`
	Favourite bool   `json:"favourite"`
}

type Card struct {
	Bank      string `json:"bank"`
	Number    string `json:"number"`
	Month     int    `json:"month"`
	Year      int    `json:"year"`
	CVV       int    `json:"cvv"`
	Owner     string `json:"owner"`
	Comment   string `json:"comment"`
	Favourite bool   `json:"favourite"`
}

type ChText struct {
	OldName        string `json:"old_name"`
	NewName        string `json:"new_name"`
	NewDescription string `json:"new_description"`
	NewComment     string `json:"new_comment"`
	NewFavourite   bool   `json:"new_favourite"`
}

type ChFile struct {
	OldName      string `json:"old_name"`
	NewName      string `json:"new_name"`
	NewComment   string `json:"new_comment"`
	NewFavourite bool   `json:"new_favourite"`
}

type ChCard struct {
	OldBank      string `json:"old_bank"`
	OldNumber    string `json:"old_number"`
	NewBank      string `json:"new_bank"`
	NewNumber    string `json:"new_number"`
	NewMonth     int    `json:"new_month"`
	NewYear      int    `json:"new_year"`
	NewCVV       int    `json:"new_cvv"`
	NewOwner     string `json:"new_owner"`
	NewComment   string `json:"new_comment"`
	NewFavourite bool   `json:"new_favourite"`
}

type CloudFile struct {
	Filename string
	Data     multipart.File
	Size     int64
}

type ChCloudFile struct {
	OldFilename string
	OldData     multipart.File
	OldSize     int64
	NewFilename string
	NewData     multipart.File
	NewSize     int64
}

type TeaErr struct {
	ToState string
	Err     string
}
