package types

import "image/color"

type Prompt interface {
	string | int | color.RGBA
}