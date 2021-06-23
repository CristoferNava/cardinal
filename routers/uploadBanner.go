package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/models"
)

// UploadBanner handles the client request, receives an image in a form a create the avatar in the database
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	// get the file from the form
	fileForm, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	extension := strings.Split(handler.Filename, ".")[1]
	var fileRoute string = "uploads/banners/" + IDUser + "." + extension

	// create a new file so we can work with all the permisions over it
	var file *os.File
	file, err = os.OpenFile(fileRoute, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Eror while uploading the image "+err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(file, fileForm) // copy fileForm to file
	if err != nil {
		http.Error(w, "Error while copying the image "+err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.User
	var status bool

	user.Banner = IDUser + "." + extension
	status, err = db.ChangeProfile(user, IDUser)
	if err != nil || !status {
		http.Error(w, "Error while storing the image "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
