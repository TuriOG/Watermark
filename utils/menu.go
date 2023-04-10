package utils

import (
	"Watermark/types"
	"github.com/manifoldco/promptui"
	"github.com/manifoldco/promptui/list"
	"golang.org/x/image/colornames"
	"image/color"
	"strings"
)

func displayPrompt[T types.Prompt](label string, items []string, searcher list.Searcher, getSelectedValue func(string) T) T {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "=> {{ . | cyan }}",
		Inactive: "{{ . | red }}",
		Selected: "You chose {{ . | cyan}}!",
	}

	prompt := promptui.Select{
		Label:     label,
		Items:     items,
		Templates: templates,
		Searcher:  searcher,
	}

	_, result, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	return getSelectedValue(result)
}

func DisplayBaseMenu() string {
	watermarkPosition := displayPrompt(
		"Welcome! Please, select the position for the watermark:",
		PositionChoices,
		nil,
		func(result string) string {
			return GetValueFromChoice(result)
		},
	)

	return watermarkPosition
}

func DisplayColorMenu() color.RGBA {
	watermarkColor := displayPrompt(
		"Great! Now select the color. You can search a specific one by hitting /",
		colornames.Names,
		func(input string, index int) bool {
			colorFound := colornames.Names[index]
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(colorFound, input)
		},
		func(result string) color.RGBA {
			return colornames.Map[result]
		},
	)

	return watermarkColor
}

func DisplayEditingMenu() string {
	editingChoice := displayPrompt(
		"Image saved! Do you want to change something?",
		[]string{"Edit Font Size", "Change Color", "Exit"},
		nil,
		func(result string) string {
			return GetValueFromChoice(result)
		},
	)

	return editingChoice
}

func DisplayFontEditMenu() string {
	editingChoice := displayPrompt(
		"How should it be?",
		[]string{"Higher", "Lower"},
		nil,
		func(result string) string {
			return GetValueFromChoice(result)
		},
	)

	return editingChoice
}
