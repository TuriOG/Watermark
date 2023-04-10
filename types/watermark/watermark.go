package watermark

import (
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
)

type Watermark struct {
	Text  string
	Color color.RGBA
	Font  *font.Face
	Point fixed.Point26_6
}

func (w *Watermark) ApplyWatermark(imageOut *image.RGBA) {
	d := font.Drawer{
		Dst:  imageOut,
		Src:  image.NewUniform(w.Color),
		Face: *w.Font,
		Dot:  w.Point,
	}

	d.DrawString(w.Text)
}