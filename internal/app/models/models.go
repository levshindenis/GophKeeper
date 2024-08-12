package models

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
	Favourite   string `json:"favourite"`
}

type File struct {
	Name      string `json:"name"`
	Comment   string `json:"comment"`
	Favourite string `json:"favourite"`
}

type Card struct {
	Bank      string `json:"bank"`
	Number    string `json:"number"`
	Month     string `json:"month"`
	Year      string `json:"year"`
	CVV       string `json:"cvv"`
	Owner     string `json:"owner"`
	Comment   string `json:"comment"`
	Favourite string `json:"favourite"`
}

type ChText struct {
	OldName        string `json:"old_name"`
	NewName        string `json:"new_name"`
	NewDescription string `json:"new_description"`
	NewComment     string `json:"new_comment"`
	NewFavourite   string `json:"new_favourite"`
}

type ChFile struct {
	OldName      string `json:"old_name"`
	NewName      string `json:"new_name"`
	NewComment   string `json:"new_comment"`
	NewFavourite string `json:"new_favourite"`
}

type ChCard struct {
	OldBank      string `json:"old_bank"`
	OldNumber    string `json:"old_number"`
	NewBank      string `json:"new_bank"`
	NewNumber    string `json:"new_number"`
	NewMonth     string `json:"new_month"`
	NewYear      string `json:"new_year"`
	NewCVV       string `json:"new_cvv"`
	NewOwner     string `json:"new_owner"`
	NewComment   string `json:"new_comment"`
	NewFavourite string `json:"new_favourite"`
}

type TeaErr struct {
	ToState string
	Err     string
}
