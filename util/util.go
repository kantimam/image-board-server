package util

import (
	"fmt"

	"github.com/disintegration/imaging"
)

func StoreThumbnails(imgSrc string, imgName string) (string, error) {

	img, err := imaging.Open(imgSrc)
	if err != nil {
		return "", err
	}
	resizedImg := imaging.Resize(img, 200, 0, imaging.Lanczos)
	err = imaging.Save(resizedImg, fmt.Sprintf("./static/thumbnails/thumbnail_%s", imgName))
	if err != nil {
		return "", err
	}
	return "", nil
}
