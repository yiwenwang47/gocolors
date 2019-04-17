package main

import (
	"gocolors"
)

func main() {
	_, err := gocolors.ExtractAndSaveRefined("c2649bb8.jpeg", "palette_1", "palette_1_Refined", 5, 0.02)
	if err != nil {
		panic(err)
	}
}
