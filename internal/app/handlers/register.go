package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"github.com/levshindenis/GophKeeper/internal/app/tools"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func (mh *MyHandler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		dec models.Register
		buf bytes.Buffer
	)

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Wrong data type", http.StatusBadRequest)
		return
	}

	if _, err := buf.ReadFrom(r.Body); err != nil {
		http.Error(w, "Something bad with read body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(buf.Bytes(), &dec); err != nil {
		http.Error(w, "Something bad with decoding JSON", http.StatusInternalServerError)
		return
	}

	if dec.Login == "" || dec.Password == "" || dec.Word == "" {
		http.Error(w, "Wrong data", http.StatusBadRequest)
		return
	}

	cookie, flag, err := mh.GetDB().AddUser(dec)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if flag {
		http.Error(w, "Repeated login", http.StatusBadRequest)
		return
	}

	cryptoBytes, err := tools.GenerateCrypto(12)
	if err != nil {
		log.Fatalf(err.Error())
	}
	secretKey := fmt.Sprintf("%x", cryptoBytes)

	f, err := os.OpenFile("../../.env", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if _, err = f.WriteString(strings.ToUpper(dec.Login) + "_SERVER=" + secretKey + "\n"); err != nil {
		log.Fatalf(err.Error())
	}
	f.Close()

	if err = mh.GetDB().AddUpdateTime(dec.Login, tools.Encrypt(time.Now().Format(time.RFC3339), secretKey)); err != nil {
		http.Error(w, "Something bad with SetUpdateTime", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "Cookie", Value: cookie})

	w.WriteHeader(http.StatusCreated)
}
