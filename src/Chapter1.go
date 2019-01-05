package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res    = 0.001
		size   = 100
		nfares = 64
		delay  = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nfares}
	phase := 0.0
	for i := 0; i < nfares; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}

func slap() {
	var a, b = 1, 3
	fmt.Println(a)
	fmt.Println(b)
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}

func PrintInput() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func CountDublicates() {
	counts := make(map[string]int)
	for _, arg := range os.Args[1:] {
		counts[arg]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}
