// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	epsX                   = (xmax - xmin) / float64(width)
	epsY                   = (ymax - ymin) / float64(height)
)

var (
	offX = []float64{-epsX, +epsX}
	offY = []float64{-epsY, +epsY}
)

func main() {
	file, err := os.Create("wow.png")
	if err != nil {
		log.Fatalln(err)
		return
	}

	draw(file)
}

func draw(out io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)*epsY + ymin
		for px := 0; px < width; px++ {
			x := float64(px)*epsX + xmin

			superPixels := make([]color.Color, 0)
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					z := complex(x+offX[i], y+offY[j])
					superPixels = append(superPixels, mandelbrot(z))
				}
			}

			// Image point (px, py) represents complex value z.
			img.Set(px, py, avg(superPixels))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func avg(colors []color.Color) color.Color {
	result := color.RGBA64{}
	for _, color := range colors {
		r, g, b, a := color.RGBA()
		result.R += uint16(r)
		result.G += uint16(g)
		result.B += uint16(b)
		result.A += uint16(a)
	}

	count := uint16(len(colors))
	result.R /= count
	result.G /= count
	result.B /= count
	result.A /= count

	return result
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50: // dark red
				return color.RGBA{100, 0, 0, 255}
			default:
				// logarithmic blue gradient to show small differences on the
				// periphery of the fractal.
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}
