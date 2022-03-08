package util

import (
	"fmt"
	"image"
	"os"
)

type Pixel struct {
	R, G, B uint32
}

type IMG struct {
	Data                   [][]Pixel
	Xmax, Xmin, Ymax, Ymin uint32
}

// Convert image to 2d slice
func imageToIMG(img image.Image) IMG {
	Xmax := img.Bounds().Max.X
	Xmin := img.Bounds().Min.X
	Ymax := img.Bounds().Max.Y
	Ymin := img.Bounds().Min.Y
	newIMG := IMG{}
	newIMG.Xmax = uint32(Xmax)
	newIMG.Ymax = uint32(Ymax)
	newIMG.Xmin = uint32(Xmin)
	newIMG.Ymin = uint32(Ymin)
	newIMG.Data = make([][]Pixel, Xmax-Xmin)
	for i, _ := range newIMG.Data {
		newIMG.Data[i] = make([]Pixel, Ymax-Ymin)
	}
	for x := Xmin; x < Xmax; x++ {
		for y := Ymin; y < Ymin; y++ {
			newIMG.Data[x][y].R, newIMG.Data[x][y].G, newIMG.Data[x][y].B, _ = img.At(x, y).RGBA()
		}
	}
	return newIMG
}

// Retrieve image data of a certain channel
func (img IMG) imgSlice(channel string) [][]uint32 {
	res := make([][]uint32, img.Xmax-img.Xmin)
	for i, _ := range res {
		res[i] = make([]uint32, img.Ymax-img.Ymin)
	}
	switch channel {
	case "R":
		for row, rowData := range img.Data {
			for column, columnData := range rowData {
				res[row][column] = columnData.R
			}
		}

	case "G":
		for row, rowData := range img.Data {
			for column, columnData := range rowData {
				res[row][column] = columnData.G
			}
		}

	case "B":
		for row, rowData := range img.Data {
			for column, columnData := range rowData {
				res[row][column] = columnData.B
			}
		}

	default:
		fmt.Println("imgSlice channel must be R or G or B.")
		os.Exit(0)
	}
	return res
}
