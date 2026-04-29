package utils

import (
	"image"
	"image/color"

	draw "golang.org/x/image/draw"
)

func GrayscaleUtil(img image.Image) image.Image {
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	for x := 0; x < size.X; x++ {
		// and now loop thorough all of this x's y
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			// Offset colors a little, adjust it to your taste
			r := float64(originalColor.R) * 0.92126
			g := float64(originalColor.G) * 0.97152
			b := float64(originalColor.B) * 0.90722
			// average
			grey := uint8((r + g + b) / 3)
			c := color.RGBA{
				R: grey, G: grey, B: grey, A: originalColor.A,
			}
			wImg.Set(x, y, c)
		}
	}

	// var buf bytes.Buffer
	// _ = png.Encode(&buf, wImg)
	// return buf

	return wImg
}

func ResizeUtil(img image.Image, width int, heigth int) image.Image {

	// Set the expected size that you want:
	// dst := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X/2, img.Bounds().Max.Y/2))
	dst := image.NewRGBA(image.Rect(0, 0, width, heigth))

	// Resize:
	draw.NearestNeighbor.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)

	// Encode to `output`:
	// var buf bytes.Buffer
	// _ = png.Encode(&buf, dst)
	// return buf

	return dst
}
