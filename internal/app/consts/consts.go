package consts

func GetStatesMap() map[string][]string {
	return map[string][]string{
		"start":  {"Регистрация", "Вход", "Выйти из программы"},
		"repeat": {"Да", "Нет"},
		"menu": {"Показать все записи", "Показать все файлы", "Показать все карты",
			"Посмотреть избранное", "Сменить пользователя", "Удалить пользователя", "Выйти из программы"},
	}
}

func GetNextStatesMap() map[string]string {
	return map[string]string{
		"reg_input_login":         "reg_input_password",
		"reg_input_password":      "reg_input_word",
		"reg_input_word":          "registration",
		"log_input_login":         "log_input_password",
		"log_input_password":      "login",
		"add_text_name":           "add_text_description",
		"add_text_description":    "add_text_comment",
		"add_text_comment":        "add_text",
		"add_card_name":           "add_card_number",
		"add_card_number":         "add_card_month",
		"add_card_month":          "add_card_year",
		"add_card_year":           "add_card_cvv",
		"add_card_cvv":            "add_card_owner",
		"add_card_owner":          "add_card_comment",
		"add_card_comment":        "add_card",
		"add_file_comment":        "add_file",
		"change_text_name":        "change_text_description",
		"change_text_description": "change_text_comment",
		"change_text_comment":     "change_text",
		"change_card_name":        "change_card_number",
		"change_card_number":      "change_card_month",
		"change_card_month":       "change_card_year",
		"change_card_year":        "change_card_cvv",
		"change_card_cvv":         "change_card_owner",
		"change_card_owner":       "change_card_comment",
		"change_card_comment":     "change_card",
		"change_file_name":        "change_file",
	}
}
