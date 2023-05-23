package exports

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
)

func DecBufToImg(data []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return img, nil
}

func EncPngImgToBuf(img image.Image, buf *bytes.Buffer) error {
	err := png.Encode(buf, img)
	if err != nil {
		return err
	}
	return nil
}

func EncJpgImgToBuf(img image.Image, buf *bytes.Buffer) error {
	err := jpeg.Encode(buf, img, nil)
	if err != nil {
		return err
	}
	return nil
}
