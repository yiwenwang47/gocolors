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
	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
	"github.com/nfnt/resize"
	"github.com/parnurzeal/gorequest"
)

type Color [4]uint32
type Newcolor [3]int

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

func saveImage(m image.Image, savename string) {
	out, err := os.Create(savename + ".jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)
}

func rbgaToRgb(c Color) Newcolor {
	a := float64(c[3]) / 32768
	r := int((float64(c[0]) * a) / 256)
	g := int((float64(c[1]) * a) / 256)
	b := int((float64(c[2]) * a) / 256)
	converted := Newcolor{r, g, b}
	return converted
}

func imageToSlice(m image.Image) []Newcolor {
	var pixel Color
	colorSlice := make([]Newcolor, 0, 3)
	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			pixel = Color{r, g, b, a}
			convertedpixel := rbgaToRgb(pixel)
			colorSlice = append(colorSlice, convertedpixel)
		}
	}
	return colorSlice
}

func extractByKmeans(colorSlice []Newcolor, k int, alpha float64) []Newcolor {
	var data clusters.Observations
	var clustered []Newcolor
	for _, pixel := range colorSlice {
		data = append(data, clusters.Coordinates{
			float64(pixel[0]),
			float64(pixel[1]),
			float64(pixel[2])})
	}
	km, err := kmeans.NewWithOptions(alpha, nil)
	clusters, err := km.Partition(data, k)
	if err != nil {
		panic(err)
	}
	for _, c := range clusters {
		newPixel := [3]int{int(c.Center[0]), int(c.Center[1]), int(c.Center[2])}
		clustered = append(clustered, newPixel)
	}
	return clustered
}

func colorsliceToJSON(colorSlice []Newcolor) string {
	data := make(map[string]interface{})
	data["model"] = "default"
	data["input"] = colorSlice
	jsonFromImage, _ := json.Marshal(data)
	return string(jsonFromImage)
}

func colormindTemp(jsonFromImage string) {
	request := gorequest.New()
	resp, body, errs := request.Get("http://colormind.io/api/").Send(jsonFromImage).End()
	if errs != nil {
		panic(errs)
	}
	if resp.StatusCode == 200 {
		fmt.Println(body)
	}
}