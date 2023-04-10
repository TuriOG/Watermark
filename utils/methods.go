package utils

import (
	"Watermark/types/watermark"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func CalculatePosition(position string, watermarkParams watermark.Params) (int, int) {
	imageX, imageY := watermarkParams.Position.X, watermarkParams.Position.Y
	waterMarkWidth := watermarkParams.Width
	fontSize := watermarkParams.FontSize
	gap := watermarkParams.Gap

	coords := strings.Split(position, "-")

	if len(coords) == 1 {
		coords = []string{"center", coords[0]}
	}

	switch coords[0] {
	case "center":
		imageY /= 2
	case "top":
		imageY = int(fontSize)
	case "bottom":
		imageY -= gap
	}

	switch coords[1] {
	case "center":
		imageX = (imageX - waterMarkWidth) / 2
	case "right":
		imageX -= waterMarkWidth + gap
	case "left":
		imageX = gap
	}

	return imageX, imageY
}

func GetValueFromChoice(choice string) string {
	splitString := strings.Split(choice, " ")

	for x := range splitString {
		splitString[x] = strings.ToLower(splitString[x])
	}

	return strings.Join(splitString, "-")
}

func TextTooLong(textLength, imageX, gap int, position string) bool {
	if position == "center" {
		return textLength > imageX
	}

	return textLength+gap > imageX
}

func ClearScreen() {
	var command *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		command = exec.Command("clear")
	case "windows":
		command = exec.Command("cmd", "/c", "cls")
	}

	command.Stdout = os.Stdout
	err := command.Run()

	if err != nil {
		return
	}
}
