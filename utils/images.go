package utils

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

func LoadImage(imagePath string) image.Image {
	imageFile, err := os.Open(imagePath)

	defer func(imageFile *os.File) {
		err := imageFile.Close()
		if err != nil {
			panic(err)
		}
	}(imageFile)

	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(imageFile)

	if err != nil {
		panic(err)
	}

	return img
}

func SaveImage(imagePath string, image *image.RGBA) {
	imageFile, err := os.Create(imagePath)

	defer func(imageFile *os.File) {
		err := imageFile.Close()
		if err != nil {
			panic(err)
		}
	}(imageFile)

	if err != nil {
		panic(err)
	}

	err = png.Encode(imageFile, image)

	if err != nil {
		panic(err)
	}
}

func loadFont() *truetype.Font {
	fontBytes, err := os.ReadFile(FontPath)

	if err != nil {
		panic(err)
	}

	parsedFont, err := truetype.Parse(fontBytes)

	if err != nil {
		panic(err)
	}

	return parsedFont
}

func GetFace(fontSize float64) font.Face {
	return truetype.NewFace(loadFont(), &truetype.Options{
		Size: fontSize,
	})
}

func EditFontSize(choice string, currentFontSize float64) float64 {
	if choice == "higher" {
		currentFontSize += FontSizeIncrement
	} else {
		currentFontSize -= FontSizeIncrement
	}

	return currentFontSize
}
