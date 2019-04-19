package main

import "github.com/yiwenwang9702/gocolors"

func wrong(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	clustered, err := gocolors.ExtractByKmeans("c2649bb8.jpeg", 5, 0.01)
	wrong(err)
	refined, err := gocolors.Refine(clustered)
	wrong(err)
	img := gocolors.CreatePalette(clustered)
	imgR := gocolors.CreatePalette(refined)
	gocolors.SaveImage(img, "palette_0")
	gocolors.SaveImage(imgR, "palette_0_Refined")
}
