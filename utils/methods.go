package utils

import (
	"Watermark/types/watermark"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func CalculateWaterMarkPosition(position string, watermarkParams watermark.Params) (int, int) {
	imageX, imageY := watermarkParams.ImgCoordinates.X, watermarkParams.ImgCoordinates.Y
	waterMarkWidth := watermarkParams.Width
	fontSize := int(watermarkParams.FontSize)
	defaultFontSize := int(DefaultFontSize)

	coords := strings.Split(position, "-")

	if len(coords) == 1 {
		coords = []string{"center", coords[0]}
	}

	switch coords[0] {
	case "center":
		imageY /= 2
	case "top":
		imageY = fontSize
	case "bottom":
		imageY -= defaultFontSize
	}

	switch coords[1] {
	case "center":
		imageX = (imageX - waterMarkWidth) / 2
	case "right":
		imageX -= waterMarkWidth + defaultFontSize
	case "left":
		imageX = defaultFontSize
	}

	return imageX, imageY
}

func TextTooLong(textLength, imageX int, position string) bool {
	if position == "center" {
		return textLength > imageX
	}

	return textLength+int(DefaultFontSize) > imageX
}

func ClearScreen() {
	var command *exec.Cmd

	if runtime.GOOS == "windows" {
		command = exec.Command("cmd", "/c", "cls")
	} else {
		command = exec.Command("clear")
	}

	command.Stdout = os.Stdout
	err := command.Run()

	if err != nil {
		return
	}
}
