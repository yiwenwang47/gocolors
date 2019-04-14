package gocolors

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

// SaveImage is a simple writer function.
func SaveImage(m image.Image, savename string) {
	out, err := os.Create(savename + ".jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)
}

//ExtractRGB returns a slice of colors in RGB.
//The basic approach: blur, resize, kmeans cluster.
func ExtractRGB(filename string, k int, alpha float64) []Newcolor {
	img := blurAndResize(filename)
	imgSlice := imageToSlice(img)
	clustered := extractByKmeans(imgSlice, k, alpha)
	return clustered
}

//ExtractPalette extracts a palette and saves it as a jpg file.
func ExtractPalette(filename string, savename string, k int, alpha float64) {
	clustered := ExtractRGB(filename, k, alpha)
	img := image.NewRGBA(image.Rect(0, 0, k*200, 200))
	for i, c := range clustered {
		for x := i * 200; x < (i+1)*200; x++ {
			for y := 0; y < 200; y++ {
				img.Set(x, y, color.RGBA{uint8(c[0]), uint8(c[1]), uint8(c[2]), 255})
			}
		}
	}
}

//More will be available.
