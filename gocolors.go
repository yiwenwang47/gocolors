package main

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"os"

	// "github.com/parnurzeal/gorequest"
	_ "image/jpeg"
)

type color [4]uint32
type newcolor [3]int

func rbgaToRgb(c color) newcolor {
	var converted newcolor
	a := float64(c[3]) / 32768
	r := (float64(c[0]) * a) / 256
	g := (float64(c[1]) * a) / 256
	b := (float64(c[2]) * a) / 256
	converted[0] = int(r)
	converted[1] = int(g)
	converted[2] = int(b)
	return converted
}
func main() {
	var pixel color
	// request := gorequest.New()
	// resp, body, errs := request.Get("https://www.youtube.com/").End()
	// if errs != nil {
	// 	panic(errs)
	// }
	// if resp.StatusCode == 200 {
	// 	fmt.Println(body[120:140])
	// }
	reader, err := os.Open("test_pics/IMG_1036.JPG")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]interface{})
	data["model"] = "default"
	input := make([]newcolor, 0, 3)
	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			pixel = [4]uint32{r, g, b, a}
			convertedpixel := rbgaToRgb(pixel)
			input = append(input, convertedpixel)
		}
	}
	data["input"] = input

	js, _ := json.Marshal(data)
	fmt.Println(string(js)[len(js)-50:])
}
