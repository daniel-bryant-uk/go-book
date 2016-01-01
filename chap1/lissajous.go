//Lissajous generates GIF animations of random Lissajous figures
package main

import (
	"image/color"
	"os"
	"io"
	"math/rand"
	"image/gif"
	"image"
	"math"
	"fmt"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}}

const (
	whiteIndex = 0 //first color in palette
	blackIndex = 1
	thirdIndex = 2
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0;

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			var lineColorIndex uint8
			if i % 2 == 0 {
				lineColorIndex = blackIndex
				fmt.Fprintf(os.Stderr, "-")
			} else {
				lineColorIndex = thirdIndex
				fmt.Fprintf(os.Stderr, ".")
			}
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), lineColorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //NOTE: ignoring encoding errors
}


