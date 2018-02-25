package gcloud

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/sj14/web-demo/infrastructure/repositories/blobs"
)

func (interactor *GCloudStore) StoreUserProfilePicture(userId int64, file []byte) (int64, error) {
	ctx := context.Background()

	var fileReader = bytes.NewReader(file)
	userIdStr := strconv.FormatInt(userId, 10)

	wc := interactor.client.Bucket("TODO").Object("user/" + userIdStr + "/images/profile.jpg").NewWriter(ctx)
	if _, err := io.Copy(wc, fileReader); err != nil {
		return -1, err
	}
	if err := wc.Close(); err != nil {
		return -1, err
	}
	return -1, nil
}

func (interactor *GCloudStore) RetrieveUserProfilePicture(userId int64) ([]byte, error) {
	ctx := context.Background()

	userIdStr := strconv.FormatInt(userId, 10)

	rc, err := interactor.client.Bucket("TODO").Object("user/" + userIdStr + "/images/profile.jpg").NewReader(ctx)
	if err != nil {
		return blobs.LoadPlaceholderPicture()
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return blobs.LoadPlaceholderPicture()
	}
	return data, nil
}
