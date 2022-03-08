package util

import (
	"image"
	"log"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

type Pixel struct {
	R, G, B int8
}

func openImage(filePath string) image.Image {
	file, err := os.Open(filePath)
	if err != nil {
		file.Close()
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func writeImage(filePath string) {

}
