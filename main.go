package main

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	// load and open the image
	src, err := imaging.Open("assets/images/ascii-pineapple.jpg")
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}

	imageWidth := src.Bounds().Dx()
	imageHeight := src.Bounds().Dy()

	fmt.Printf("Image width is: %v\nImage height is %v\n", imageWidth, imageHeight)
}
