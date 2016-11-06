//Server3 обработчик HTTP-запросов
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0, 128, 0, 255}, color.RGBA{255, 0, 0, 255}}

func main() {
	http.HandleFunc("/", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

//Обработчик
func lissajous(out http.ResponseWriter, r *http.Request) {
	const (
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    //Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)

	var cycles = 5.0
	if res := r.URL.Query().Get("cycles"); res != "" {
		cycles, _ = strconv.ParseFloat(res, 64)
		fmt.Printf("%v", cycles)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 //Относительная частота колебаний
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		colorIndex := uint8(rand.Intn(2) + 1)
		fmt.Printf("%v", cycles)
		for t := 0.0; t < cycles*2.0*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Примечание: игнорируем ошибки
}
