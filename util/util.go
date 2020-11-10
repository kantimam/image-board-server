package util

import (
	"fmt"
	"time"

	"github.com/disintegration/imaging"
)

func StoreThumbnails(imgSrc string, imgName string) (string, error) {
	time.Sleep(time.Second * 5)
	fmt.Printf("started %s \n", imgSrc)
	img, err := imaging.Open(imgSrc)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	resizedImg := imaging.Resize(img, 200, 0, imaging.Lanczos)
	err = imaging.Save(resizedImg, fmt.Sprintf("./static/thumbnails/thumbnail_%s", imgName))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("done")
	return "", nil
}
