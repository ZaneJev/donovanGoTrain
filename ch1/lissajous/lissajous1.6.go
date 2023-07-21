// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	palette1 = color.Palette{
		color.RGBA{R: 255, G: 0, B: 115, A: 255},
		color.RGBA{R: 123, G: 100, B: 31, A: 255},
	}

	palette2 = color.Palette{
		color.RGBA{R: 255, A: 255},         // Красный
		color.RGBA{G: 255, A: 255},         // Зеленый
		color.RGBA{B: 255, A: 255},         // Синий
		color.RGBA{R: 255, G: 255, A: 255}, // Желтый
		color.RGBA{R: 255, B: 255, A: 255}, // Фиолетовый
	}

	palette3 = color.Palette{
		color.RGBA{R: 255, A: 255},                 // Красный
		color.RGBA{G: 255, A: 255},                 // Зеленый
		color.RGBA{B: 255, A: 255},                 // Синий
		color.RGBA{A: 255},                         // Черный
		color.RGBA{R: 255, G: 255, B: 255, A: 255}, // Белый
	}

	palettes []color.Palette
)

// !+main

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajousHard(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajousHard(os.Stdout)
}

func lissajousHard(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	palettes = append(palettes, palette1, palette2, palette3)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		for j := 0; j < len(palettes); j++ {
			img := image.NewPaletted(rect, palettes[j])
			for t := 0.0; t < cycles*2*math.Pi; t += res {
				x := math.Sin(t)
				y := math.Sin(t*freq + phase)
				// TODO картинки одинаковые, надо разобраться почему
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					uint8(int(t/res)%len(palettes)))
			}
			phase += 0.1
			anim.Delay = append(anim.Delay, delay)
			anim.Image = append(anim.Image, img)
			f, _ := os.Create("animation" + strconv.Itoa(j) + ".gif")
			gif.EncodeAll(f, &anim)
			f.Close()
		}
	}
}

//!-main
