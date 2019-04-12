package gocolors

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"math"
	"os"

	"image/jpeg"
	_ "image/png" //For png files.

	"github.com/esimov/stackblur-go"
	"github.com/nfnt/resize"
	"github.com/parnurzeal/gorequest"
)

type color [4]uint32
type newcolor [3]int

func blurRadius(m image.Image) uint32 {
	bounds := m.Bounds()
	height := bounds.Max.Y - bounds.Min.Y
	width := bounds.Max.X - bounds.Min.X
	x := math.Sqrt(float64(width * height))
	x = x / 10
	return uint32(x)
}

func blurAndResize(name string) image.Image {
	reader, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	mPreBlur, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	radius := blurRadius(mPreBlur)
	mPreResize := stackblur.Process(mPreBlur, radius)
	m := resize.Resize(100, 0, mPreResize, resize.Lanczos3)
	return m
}

func saveResized(m image.Image, savename string) {
	out, err := os.Create(savename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)
}

func rbgaToRgb(c color) newcolor {
	a := float64(c[3]) / 32768
	r := int((float64(c[0]) * a) / 256)
	g := int((float64(c[1]) * a) / 256)
	b := int((float64(c[2]) * a) / 256)
	converted := newcolor{r, g, b}
	return converted
}

func imageToSlice(m image.Image) []newcolor {
	var pixel color
	colorSlice := make([]newcolor, 0, 3)
	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			pixel = color{r, g, b, a}
			convertedpixel := rbgaToRgb(pixel)
			colorSlice = append(colorSlice, convertedpixel)
		}
	}
	return colorSlice
}

// Some resampling algorithm needs to be written.
// Maybe kmeans.

func imageToJSON(m image.Image) string {
	input := imageToSlice(m)
	data := make(map[string]interface{})
	data["model"] = "default"
	data["input"] = input
	jsonFromImage, _ := json.Marshal(data)
	return string(jsonFromImage)
}

func colormindGoAPI(jsonFromImage string) {
	request := gorequest.New()
	resp, body, errs := request.Get("http://colormind.io/api/").Send(jsonFromImage).End()
	if errs != nil {
		panic(errs)
	}
	if resp.StatusCode == 200 {
		fmt.Println(body)
	}
}
