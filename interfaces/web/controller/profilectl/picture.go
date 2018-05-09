package profilectl

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func (interactor *ProfileController) PostPicture(w http.ResponseWriter, r *http.Request) {
	// The max multipartfile size: 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, 1000*1000)

	// NOT a filelimit!
	// 1 MB in memory and then cache on disk
	r.ParseMultipartForm(1000 * 1000)

	_, fileHeader, err := r.FormFile("upload_picture")
	if err != nil {
		// This also happens, when no file was uploaded.
		log.Println("uploading profile picture failed: ", err)
		interactor.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		log.Println("opening file header failed:", err)
		interactor.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Println("storing profile picture failed: ", err)
		interactor.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	userID, ok := interactor.Cookie.SessionGetUserId(r)
	if !ok {
		log.Println("user not logged in")
		interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	err = interactor.ImageUsecases.StoreUserProfilePicture(userID, buf.Bytes())
	if err != nil {
		log.Println("storing profile picture failed: ", err)
		interactor.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/profile/edit", http.StatusFound)
}
