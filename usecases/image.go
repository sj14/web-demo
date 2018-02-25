package usecases

func NewImageUsecases(imageRepository imageRepositoryInterface) ImageUsecases {
	return ImageUsecases{repository: imageRepository}
}

type ImageUsecases struct {
	repository imageRepositoryInterface
}

type imageRepositoryInterface interface {
	StoreUserProfilePicture(userId int64, file []byte) (int64, error)
	RetrieveUserProfilePicture(userId int64) ([]byte, error)
}

func (interactor *ImageUsecases) StoreUserProfilePicture(userId int64, dat []byte) error {
	// Check Filetype, Convert and remove EXIF/meta info
	dat, err := interactor.checkFiletypeAndConvert(dat)
	if err != nil {
		return err
	}

	// Resize Image
	dat, err = interactor.resizeJpegImage(dat)
	if err != nil {
		return err
	}

	_, err = interactor.repository.StoreUserProfilePicture(userId, dat)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *ImageUsecases) RetrieveUserProfilePicture(userId int64) ([]byte, error) {
	image, err := interactor.repository.RetrieveUserProfilePicture(userId)
	if err != nil {
		return []byte{}, err
	}
	return image, nil
}
