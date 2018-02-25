package usecases

import (
	"bufio"
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"

	"github.com/nfnt/resize"
)

func (interactor *ImageUsecases) checkFiletypeAndConvert(dat []byte) ([]byte, error) {
	// Sources:
	// https://www.socketloop.com/tutorials/golang-how-to-verify-uploaded-file-is-image-or-allowed-file-types
	// https://stackoverflow.com/a/38175140/7125878

	buff := make([]byte, 512) // docs tell that it take only first 512 bytes into consideration

	file := bytes.NewReader(dat)
	if _, err := file.Read(buff); err != nil {
		return []byte{}, err
	}

	filetype := http.DetectContentType(buff)

	// Not sure how to correcly reset the Reader,
	// just creating the Reader again :-/
	file = bytes.NewReader(dat)

	switch filetype {
	case "image/jpeg", "image/jpg":
		imageFile, err := jpeg.Decode(file)
		if err != nil {
			return []byte{}, err
		}
		// Should also remove EXIF/meta info
		jpegBytes, err := interactor.imageToJpeg(imageFile)
		if err != nil {
			return []byte{}, err
		}
		return jpegBytes, nil

	case "image/gif":
		imageFile, err := gif.Decode(file)
		if err != nil {
			return []byte{}, err
		}
		jpegBytes, err := interactor.imageToJpeg(imageFile)
		if err != nil {
			return []byte{}, err
		}
		return jpegBytes, nil

	case "image/png":
		imageFile, err := png.Decode(file)
		if err != nil {
			return []byte{}, err
		}
		jpegBytes, err := interactor.imageToJpeg(imageFile)
		if err != nil {
			return []byte{}, err
		}
		return jpegBytes, nil

	default:
		return []byte{}, errors.New("unsupported file type uploaded")
	}
}

func (interactor *ImageUsecases) imageToJpeg(imageFile image.Image) ([]byte, error) {
	jpegBuffer := bytes.NewBuffer(nil) // same as &bytes.Buffer{}, so it is ok as well.
	bufferWriter := bufio.NewWriter(jpegBuffer)

	err := jpeg.Encode(bufferWriter, imageFile, nil)
	if err != nil {
		return []byte{}, err
	}
	bufferWriter.Flush() // Don't forget to flush!
	return jpegBuffer.Bytes(), nil
}

func (interactor *ImageUsecases) resizeJpegImage(dat []byte) ([]byte, error) {
	reader, err := jpeg.Decode(bytes.NewReader(dat))
	if err != nil {
		return []byte{}, err
	}
	resizedImage := resize.Thumbnail(500, 500, reader, resize.Bilinear)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, resizedImage, nil)
	if err != nil {
		return []byte{}, err
	}
	resizedBytes := buf.Bytes()

	return resizedBytes, nil
}
