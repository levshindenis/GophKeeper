package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/levshindenis/GophKeeper/internal/app/models"
	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (m *model) ChangeText() {
	var (
		at models.ChText
	)
	newName := m.textItem.Name
	newDescription := m.textItem.Description
	newComment := m.textItem.Comment
	arr := strings.Split(m.helpStr, "///")
	if arr[1] != "" {
		newName = tools.Encrypt(arr[1], m.secretKey)
	}
	if arr[2] != "" {
		newDescription = tools.Encrypt(arr[2], m.secretKey)
	}
	if arr[3] != "" {
		newComment = tools.Encrypt(arr[3], m.secretKey)
	}

	at = models.ChText{OldName: m.textItem.Name, NewName: newName,
		NewDescription: newDescription, NewComment: newComment, NewFavourite: m.textItem.Favourite}

	if err := m.db.ChangeText(m.userId, at); err != nil {
		m.ErrorState(err.Error(), "change_text_name")
		return
	}

	m.state = "text_view"
	m.helpStr = newName
	m.cursor = 5
	m.TextInfo()
}

func (m *model) ChangeCard() {
	var (
		at models.ChCard
	)
	newBank := m.cardItem.Bank
	newNumber := m.cardItem.Number
	newMonth := m.cardItem.Month
	newYear := m.cardItem.Year
	newCVV := m.cardItem.CVV
	newOwner := m.cardItem.Owner
	newComment := m.cardItem.Comment
	arr := strings.Split(m.helpStr, "///")
	if arr[1] != "" {
		newBank = tools.Encrypt(arr[1], m.secretKey)
	}
	if arr[2] != "" {
		newNumber = tools.Encrypt(arr[2], m.secretKey)
	}
	if arr[3] != "" {
		newMonth = tools.Encrypt(arr[3], m.secretKey)
	}
	if arr[4] != "" {
		newYear = tools.Encrypt(arr[4], m.secretKey)
	}
	if arr[5] != "" {
		newCVV = tools.Encrypt(arr[5], m.secretKey)
	}
	if arr[6] != "" {
		newOwner = tools.Encrypt(arr[6], m.secretKey)
	}
	if arr[7] != "" {
		newComment = tools.Encrypt(arr[7], m.secretKey)
	}

	at = models.ChCard{OldBank: m.cardItem.Bank, OldNumber: m.cardItem.Number, NewBank: newBank,
		NewNumber: newNumber, NewMonth: newMonth, NewYear: newYear, NewCVV: newCVV, NewOwner: newOwner,
		NewComment: newComment, NewFavourite: m.cardItem.Favourite}

	if err := m.db.ChangeCard(m.userId, at); err != nil {
		m.ErrorState(err.Error(), "change_card_name")
		return
	}

	m.state = "card_view"
	m.helpStr = newBank + "///" + newNumber
	m.cursor = 8
	m.CardInfo()
}

func (m *model) ChangeFile() {
	var (
		at models.ChFile
	)
	newName := m.fileItem.Name
	newComment := m.fileItem.Comment
	arr := strings.Split(m.helpStr, "///")
	if arr[1] != "" {
		newName = tools.Encrypt(arr[1], m.secretKey)
	}
	if arr[2] != "" {
		newComment = tools.Encrypt(arr[2], m.secretKey)
	}

	at = models.ChFile{OldName: m.fileItem.Name, NewName: tools.Encrypt(filepath.Base(newName), m.secretKey),
		NewComment: newComment, NewFavourite: m.fileItem.Favourite}

	if err := m.db.ChangeFile(m.userId, at); err != nil {
		m.ErrorState(err.Error(), "change_file_name")
		return
	}

	if newName != m.fileItem.Name {
		if err := m.cloud.DeleteFile(m.userId, arr[0]); err != nil {
			m.ErrorState(err.Error(), "change_file_name")
			return
		}
		if err := m.cloud.AddFile(m.userId, newName); err != nil {
			m.ErrorState(err.Error(), "change_file_name")
			return
		}

		if err := os.Remove("/tmp/keeper/files/" + m.userId + "/" + arr[0]); err != nil {
			m.ErrorState(err.Error(), "change_file_name")
			return
		}

		source, err := os.Open(newName)
		if err != nil {
			m.ErrorState(err.Error(), "change_file_name")
			return
		}
		defer source.Close()

		destination, err := os.Create("/tmp/keeper/files/" + m.userId + "/" + filepath.Base(newName))
		if err != nil {
			m.ErrorState(err.Error(), "change_file_name")
			return
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			m.ErrorState(err.Error(), "change_file_name")
			return
		}
	}

	m.state = "file_view"
	m.helpStr = filepath.Base(newName)
	m.cursor = 5
	m.FileInfo()
}
