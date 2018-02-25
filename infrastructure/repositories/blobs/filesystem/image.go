package filesystem

import (
	"strconv"

	"gitlab.com/sj14/web-clean/src/infrastructure/repositories/blobs"
)

func (interactor *FilesystemStore) StoreUserProfilePicture(userId int64, dat []byte) (int64, error) {
	userIdStr := strconv.FormatInt(userId, 10)
	filePath := "user/" + userIdStr + "/images/profile.jpg"
	err := interactor.StoreFile(filePath, dat)
	if err != nil {
		return -1, err
	}
	return -1, nil
}

func (interactor *FilesystemStore) StoreGastronomeProfilePicture(applicantId int64, dat []byte) (int64, error) {
	applicantIdStr := strconv.FormatInt(applicantId, 10)

	filePath := "gastronome/" + applicantIdStr + "/images/profile.jpg"
	err := interactor.StoreFile(filePath, dat)
	if err != nil {
		return -1, err
	}
	return -1, nil
}

func (interactor *FilesystemStore) StoreApplicantProfilePicture(applicantId int64, dat []byte) (int64, error) {
	applicantIdStr := strconv.FormatInt(applicantId, 10)

	filePath := "applicant/" + applicantIdStr + "/images/profile.jpg"
	err := interactor.StoreFile(filePath, dat)
	if err != nil {
		return -1, err
	}
	return -1, nil
}

func (interactor *FilesystemStore) RetrieveUserProfilePicture(userId int64) ([]byte, error) {
	userIdStr := strconv.FormatInt(userId, 10)
	filePath := "user/" + userIdStr + "/images/profile.jpg"

	picture, err := interactor.RetrieveFile(filePath)
	if err == nil {
		// NO error, return the profile picture
		return picture, nil
	}

	// On error, load placeholder image
	placeholder, err := blobs.LoadPlaceholderPicture()
	if err != nil {
		return []byte{}, err
	}
	return placeholder, nil
}

func (interactor *FilesystemStore) RetrieveGastronomeProfilePicture(userId int64) ([]byte, error) {
	userIdStr := strconv.FormatInt(userId, 10)
	filePath := "gastronome/" + userIdStr + "/images/profile.jpg"

	picture, err := interactor.RetrieveFile(filePath)
	if err == nil {
		// NO error, return the profile picture
		return picture, nil
	}

	// On error, load placeholder image
	placeholder, err := blobs.LoadPlaceholderPicture()
	if err != nil {
		return []byte{}, err
	}
	return placeholder, nil
}

func (interactor *FilesystemStore) RetrieveApplicantProfilePicture(applicantId int64) ([]byte, error) {
	applicantIdStr := strconv.FormatInt(applicantId, 10)
	filePath := "applicant/" + applicantIdStr + "/images/profile.jpg"

	picture, err := interactor.RetrieveFile(filePath)
	if err == nil {
		// NO error, return the profile picture
		return picture, nil
	}

	// On error, load placeholder image
	placeholder, err := blobs.LoadPlaceholderPicture()
	if err != nil {
		return []byte{}, err
	}
	return placeholder, nil
}
