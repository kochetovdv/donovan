// Server5 - сервер + Лисажу с параметрами
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
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

type Lissajous struct {
	cycles  int     // number of complete x oscillator revolutions
	res     float64 // angular resolution
	size    int     // image canvas covers [-size..+size]
	nframes int     // number of animation frames
	delay   int     // delay between frames in 10ms units
}

func main() {
	url := "localhost:8000"
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(url, nil))
}

// Parse query params
func handler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	l := Lissajous{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}
	for k, v := range values {
		switch k {
		case "cycles":
			l.cycles, _ = strconv.Atoi(v[0])
		case "res":
			l.res, _ = strconv.ParseFloat(v[0], 64)
		case "size":
			l.size, _ = strconv.Atoi(v[0])
		case "nframes":
			l.nframes, _ = strconv.Atoi(v[0])
		case "delay":
			l.delay, _ = strconv.Atoi(v[0])
		}
	}
	lissajous(w, l)
}

func lissajous(out io.Writer, l Lissajous) {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: l.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < l.nframes; i++ {
		rect := image.Rect(0, 0, 2*l.size+1, 2*l.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(l.cycles)*2*math.Pi; t += l.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(l.size+int(x*float64(l.size)+0.5), l.size+int(y*float64(l.size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, l.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
