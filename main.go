package main

import (
	"Watermark/types"
	"Watermark/types/watermark"
	"Watermark/utils"
	"bufio"
	"fmt"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"time"
)

func main() {
	waterMarkPosition := utils.DisplayBaseMenu()
	waterMarkColor := utils.DisplayColorMenu()
	fontSize := utils.DefaultFontSize
	face := utils.GetFace(fontSize)

	srcImage := utils.LoadImage(utils.InputPath)
	imageBounds := srcImage.Bounds()
	imageCoords := types.Coords{X: imageBounds.Dx(), Y: imageBounds.Dy()}

	scanner := bufio.NewScanner(os.Stdin)
	var inputText string
	var waterMarkWidth int

	for {
		fmt.Println("Ok! Now enter the text that will be used as a watermark:")
		scanner.Scan()
		inputText = scanner.Text()
		waterMarkWidth = font.MeasureString(face, inputText).Ceil()

		if len(inputText) == 0 {
			fmt.Println("You didn't enter any text. Please try again.")
			continue
		}

		if utils.TextTooLong(waterMarkWidth, imageCoords.X, waterMarkPosition) {
			fmt.Println("The text is too long! Please try again.")
			continue
		}

		break
	}

	var editingChoice string
	for {
		if editingChoice == "exit" {
			break
		}

		newImage := image.NewRGBA(imageBounds)
		draw.Draw(newImage, imageBounds, srcImage, imageBounds.Min, draw.Src)

		waterMarkX, waterMarkY := utils.CalculateWaterMarkPosition(waterMarkPosition, watermark.Params{
			ImgCoordinates: imageCoords,
			Width:          waterMarkWidth,
			FontSize:       fontSize,
		})

		imgWatermark := watermark.Watermark{
			Text:  inputText,
			Color: waterMarkColor,
			Font:  &face,
			Point: fixed.Point26_6{X: fixed.Int26_6(waterMarkX * 64), Y: fixed.Int26_6(waterMarkY * 64)},
		}

		imgWatermark.ApplyWatermark(newImage)
		utils.SaveImage(utils.OutputPath, newImage)
		utils.ClearScreen()

		for {
			editingChoice = utils.DisplayEditingMenu()

			switch editingChoice {
			case "edit-font-size":
				fontChoice := utils.DisplayFontEditMenu()

				newFontSize := utils.EditFontSize(fontChoice, fontSize)
				newFace := utils.GetFace(newFontSize)
				newWaterMarkWidth := font.MeasureString(newFace, inputText).Ceil()

				if utils.TextTooLong(newWaterMarkWidth, imageCoords.X, waterMarkPosition) {
					fmt.Println("The text is too long! You can't increase it any further.")
					continue
				}

				fontSize = newFontSize
				waterMarkWidth = newWaterMarkWidth
				face = newFace

			case "change-color":
				waterMarkColor = utils.DisplayColorMenu()
			default:
				break
			}

			break
		}

		time.Sleep(time.Millisecond)
		utils.ClearScreen()
	}

	fmt.Print("\nGoodbye!")
}
