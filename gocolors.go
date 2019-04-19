package gocolors

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"strings"
)

//Color ... a slice representing R, G, B, A.
type Color [4]uint32

//Newcolor ... a slice representing R, G, B.
type Newcolor [3]int

//ColorHSL ... a slice representing H, S, L.
type ColorHSL [3]int

//ColorSlice ... a slice or RGB colors.
type ColorSlice []Newcolor

//SaveImage ... a simple writer function.
func SaveImage(m image.Image, savename string) {
	out, err := os.Create(savename + ".jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)
}

//CreatePalette ... creates a palette image.
func CreatePalette(colorslice ColorSlice) image.Image {
	k := len(colorslice)
	img := image.NewRGBA(image.Rect(0, 0, k*200, 200))
	for i, c := range colorslice {
		for x := i * 200; x < (i+1)*200; x++ {
			for y := 0; y < 200; y++ {
				img.Set(x, y, color.RGBA{uint8(c[0]), uint8(c[1]), uint8(c[2]), 255})
			}
		}
	}
	return img
}

//Parser ... parses the json result from Colormind API.
func Parser(result string) ColorSlice {
	var parsed ColorSlice
	result = strings.ReplaceAll(result, "[", "")
	result = strings.ReplaceAll(result, "]", "")
	resultSplit := strings.Split(result, ",")
	for i := 0; i < 5; i++ {
		r, errR := strconv.Atoi(resultSplit[i*3])
		g, errG := strconv.Atoi(resultSplit[i*3+1])
		b, errB := strconv.Atoi(resultSplit[i*3+2])
		if errR != nil || errG != nil || errB != nil {
			return nil
		}
		parsed = append(parsed, Newcolor{r, g, b})
	}
	return parsed
}

//ExtractByKmeans ... returns a slice of colors in RGB.
//The basic approach: blur, resize, kmeans cluster.
func ExtractByKmeans(filename string, k int, alpha float64) (ColorSlice, error) {
	img := blurAndResize(filename)
	imgSlice := imageToSlice(img)
	clustered, err := extractByKmeans(imgSlice, k, alpha)
	return clustered, err
}

//Refine ... takes advantage of the Colormind API.
func Refine(clustered ColorSlice) (ColorSlice, error) {
	jsFromPic := colorsliceToJSON(clustered)
	mapOfResults, err := colormind(jsFromPic)
	if err != nil {
		return nil, err
	}
	refined := Parser(string(*mapOfResults["result"]))
	return refined, nil
}

//More will be available.
