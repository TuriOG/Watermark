package utils

var (
	AssetsPath      = "assets"
	ImagesPath      = AssetsPath + "/images"
	InputPath       = ImagesPath + "/image.jpg"
	OutputPath      = ImagesPath + "/output.png"
	FontPath        = AssetsPath + "/font.ttf"
	PositionChoices = []string{
		"Center",
		"Right",
		"Left",
		"Top Left",
		"Top Right",
		"Bottom Left",
		"Bottom Right",
	}
	DefaultFontSize   = 20.0
	FontSizeIncrement = 10.5
)
