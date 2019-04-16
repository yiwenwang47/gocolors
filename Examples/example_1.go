package main

import (
	"github.com/yiwenwang9702/gocolors"
)

func main() {
	_, err := gocolors.ExtractAndSaveRefined("c2649bb8.jpeg", "palette_0", "palette_Refined", 5, 0.02)
	if err != nil {
		panic(err)
	}
}
