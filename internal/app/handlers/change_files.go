package handlers

import (
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"net/http"
	"strconv"
)

func (mh *MyHandler) ChangeFiles(w http.ResponseWriter, r *http.Request) {
	var (
		dbFiles    []models.ChFile
		cloudFiles []models.ChCloudFile
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
		oldComment := r.Form.Get("old_comment" + strconv.Itoa(i))
		newComment := r.Form.Get("new_comment" + strconv.Itoa(i))
		oldFavourite, _ := strconv.ParseBool(r.Form.Get("old_favourite" + strconv.Itoa(i)))
		newFavourite, _ := strconv.ParseBool(r.Form.Get("new_favourite" + strconv.Itoa(i)))

		oldF, oldHandler, err := r.FormFile("old_file" + strconv.Itoa(i))
		if err != nil {
			http.Error(w, "Error FormFile", http.StatusInternalServerError)
			return
		}

		newF, newHandler, err := r.FormFile("new_file" + strconv.Itoa(i))
		if err != nil {
			http.Error(w, "Error FormFile", http.StatusInternalServerError)
			return
		}

		dbFiles = append(dbFiles,
			models.ChFile{OldName: oldHandler.Filename, OldComment: oldComment, OldFavourite: oldFavourite,
				NewName: newHandler.Filename, NewComment: newComment, NewFavourite: newFavourite})
		cloudFiles = append(cloudFiles,
			models.ChCloudFile{OldFilename: oldHandler.Filename, OldData: oldF, OldSize: oldHandler.Size,
				NewFilename: newHandler.Filename, NewData: newF, NewSize: newHandler.Size})

		oldF.Close()
		newF.Close()
	}

	if err := mh.GetCloud().ChangeFiles(userId+"ooo", cloudFiles); err != nil {
		http.Error(w, "Error with ChangeFiles(cloud)", http.StatusInternalServerError)
		return
	}

	if err := mh.GetDB().ChangeFiles(userId, dbFiles); err != nil {
		http.Error(w, "Error with ChangeFiles(db)", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
