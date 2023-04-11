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

	srcImage := utils.LoadImage(utils.ImagePath)
	bounds := srcImage.Bounds()
	coords := types.Coords{X: bounds.Dx(), Y: bounds.Dy()}
	gap := utils.Gap

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	var waterMarkWidth int

	for {
		fmt.Println("Ok! Now enter the text that will be used as a watermark:")
		scanner.Scan()
		text = scanner.Text()
		waterMarkWidth = font.MeasureString(face, text).Ceil()

		if len(text) == 0 {
			fmt.Println("You didn't enter any text. Please try again.")
			continue
		}

		if utils.TextTooLong(waterMarkWidth, coords.X, gap, waterMarkPosition) {
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

		newImage := image.NewRGBA(bounds)
		draw.Draw(newImage, bounds, srcImage, bounds.Min, draw.Src)

		x, y := utils.CalculatePosition(waterMarkPosition, watermark.Params{
			Position: coords,
			Width:    waterMarkWidth,
			FontSize: fontSize,
			Gap:      gap,
		})

		imgWatermark := watermark.Watermark{
			Text:  text,
			Color: waterMarkColor,
			Font:  &face,
			Point: fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)},
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
				newWaterMarkWidth := font.MeasureString(newFace, text).Ceil()

				if utils.TextTooLong(newWaterMarkWidth, coords.X, gap, waterMarkPosition) {
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

		time.Sleep(time.Second / 2)
		utils.ClearScreen()
	}

	fmt.Print("\nGoodbye!")
}
