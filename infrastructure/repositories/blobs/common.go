package blobs

import (
	"io/ioutil"
	"log"
)

func LoadPlaceholderPicture() ([]byte, error) {
	picture, err := ioutil.ReadFile("interfaces/web/files/static/images/profile_placeholder.jpg")
	if err != nil {
		log.Println("error reading file")
		return []byte{}, err
	}
	return picture, nil
}
