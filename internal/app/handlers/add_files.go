package handlers

import (
	"net/http"
	"strconv"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (mh *MyHandler) AddFiles(w http.ResponseWriter, r *http.Request) {
	var (
		dbFiles    []models.File
		cloudFiles []models.CloudFile
	)

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	cookie, _ := r.Cookie("Cookie")
	userId := mh.GetCookie().GetUserId(cookie.Value)

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	count, _ := strconv.Atoi(r.Form.Get("count"))

	for i := 1; i <= count; i++ {
		comment := r.Form.Get("comment" + strconv.Itoa(i))
		favourite, _ := strconv.ParseBool(r.Form.Get("favourite" + strconv.Itoa(i)))

		f, handler, err := r.FormFile("file" + strconv.Itoa(i))
		if err != nil {
			http.Error(w, "Error FormFile", http.StatusInternalServerError)
			return
		}

		dbFiles = append(dbFiles, models.File{Name: handler.Filename, Comment: comment, Favourite: favourite})
		cloudFiles = append(cloudFiles, models.CloudFile{Filename: handler.Filename, Data: f, Size: handler.Size})

		f.Close()
	}

	if err := mh.GetCloud().AddFiles(userId+"ooo", cloudFiles); err != nil {
		http.Error(w, "Error with AddFiles(cloud)", http.StatusInternalServerError)
		return
	}

	if err := mh.GetDB().AddFiles(userId, dbFiles); err != nil {
		http.Error(w, "Error with AddFiles(db)", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
