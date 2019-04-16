package gocolors

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

// SaveImage ... a simple writer function.
func SaveImage(m image.Image, savename string) {
	out, err := os.Create(savename + ".jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)
}

//ExtractRGB ... returns a slice of colors in RGB.
//The basic approach: blur, resize, kmeans cluster.
func ExtractRGB(filename string, k int, alpha float64) []Newcolor {
	img := blurAndResize(filename)
	imgSlice := imageToSlice(img)
	clustered := extractByKmeans(imgSlice, k, alpha)
	return clustered
}

//ExtractAndSave ... extracts a palette and saves it as a jpg file.
func ExtractAndSave(filename string, savename string, k int, alpha float64) {
	clustered := ExtractRGB(filename, k, alpha)
	img := image.NewRGBA(image.Rect(0, 0, k*200, 200))
	for i, c := range clustered {
		for x := i * 200; x < (i+1)*200; x++ {
			for y := 0; y < 200; y++ {
				img.Set(x, y, color.RGBA{uint8(c[0]), uint8(c[1]), uint8(c[2]), 255})
			}
		}
	}
	SaveImage(img, savename)
}

//ExtractAndSaveRefined ... extracts a palette, refines it and saves both.
func ExtractAndSaveRefined(filename string, savenameOriginal string, savenameRefined string, k int, alpha float64) ([]Newcolor, error) {
	clustered := ExtractRGB(filename, k, alpha)
	jsFromPic := colorsliceToJSON(clustered)
	mapOfResults, err := colormind(jsFromPic)
	if err != nil {
		return nil, err
	}
	refined := Parser(string(*mapOfResults["result"]))

	img := image.NewRGBA(image.Rect(0, 0, k*200, 200))
	for i, c := range clustered {
		for x := i * 200; x < (i+1)*200; x++ {
			for y := 0; y < 200; y++ {
				img.Set(x, y, color.RGBA{uint8(c[0]), uint8(c[1]), uint8(c[2]), 255})
			}
		}
	}
	SaveImage(img, savenameOriginal)
	imgRefined := image.NewRGBA(image.Rect(0, 0, 5*200, 200))
	for i, c := range refined {
		for x := i * 200; x < (i+1)*200; x++ {
			for y := 0; y < 200; y++ {
				imgRefined.Set(x, y, color.RGBA{uint8(c[0]), uint8(c[1]), uint8(c[2]), 255})
			}
		}
	}
	SaveImage(imgRefined, savenameRefined)
	return refined, nil
}

//More will be available.
