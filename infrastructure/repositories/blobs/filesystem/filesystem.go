package filesystem

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type FilesystemStore struct {
	basePath string
}

func NewFilesystemStore(basePath string) *FilesystemStore {
	fsStore := &FilesystemStore{basePath}
	return fsStore
}

func (interactor *FilesystemStore) StoreFile(filePath string, dat []byte) error {
	fullPath := interactor.basePath + filePath
	dirPath := filepath.Dir(fullPath)

	err := os.MkdirAll(dirPath, 0700)
	if err != nil {
		log.Println("failed to create dir:", err)
		return err
	}

	err = ioutil.WriteFile(fullPath, dat, 0644)
	if err != nil {
		log.Println("failed to store file:", err)
		return err
	}
	return nil
}

func (interactor *FilesystemStore) RetrieveFile(filePath string) ([]byte, error) {
	dat, err := ioutil.ReadFile(interactor.basePath + filePath)
	if err != nil {
		return []byte{}, err
	}
	return dat, nil
}
