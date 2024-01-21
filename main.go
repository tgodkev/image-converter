package main

import (
	"flag"
	"fmt"
	"github.com/containerd/containerd/v2/core/images/converter"
	"github.com/disintegration/imaging"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var (
		format             string
		compressionQuality int
		targetFormat       string
		cropSize           string
	)

	// Define and parse command-line flagear

	flag.StringVar(&format, "format", "jpg", "Image format to process (default: jpg)")
	flag.IntVar(&compressionQuality, "quality", 75, "Compression quality in percentage (default: 75)")
	flag.StringVar(&targetFormat, "convert", "webp", "Format to convert to (default: webp)")
	flag.StringVar(&cropSize, "crop", "", "Size to crop images to ex: '300x300' (height x width) (default: '') ")

	flag.Parse()

	// Image processing logic
	dir, _ := os.Getwd()
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+format {
			ProcessImage(path, cropSize, compressionQuality, targetFormat)
		}
		return nil
	})
}

// ProcessImage - function for image processing
func ProcessImage(path string, cropSize string, quality int, conversionFormat string) {
	fmt.Printf("Processing %s: CropSize=%d, Quality=%d, ConvertTo=%s\n", path, cropSize, quality, conversionFormat)

	if cropSize != "" {
		dimensions := strings.Split(cropSize, "x")

		if len(dimensions) == 2 {
			heightStr, widthStr := dimensions[0], dimensions[1]

			// Convert heightStr and widthStr to integers
			height, err := strconv.Atoi(heightStr)
			if err != nil {
				fmt.Printf("Invalid height: %s\n", heightStr)
				return
			}

			width, err := strconv.Atoi(widthStr)
			if err != nil {
				fmt.Printf("Invalid width: %s\n", widthStr)
				return
			}

			img, err := imaging.Open(path)
			if err != nil {
				fmt.Printf("Failed to open image: %s\n", err)
				return
			}

			resizedImg := imaging.Resize(img, height, width, imaging.Lanczos)

			outputFilePath := "resized_" + path
			err = imaging.Save(resizedImg, outputFilePath)

			fmt.Printf("Image processed and saved to %s\n", outputFilePath)

			if err != nil {
				fmt.Printf("Failed to save resized image: %s\n", err)
				return
			}

		} else {
			fmt.Printf("Invalid crop size: %s\n", cropSize)
			return
		}
	}

	// Example: if quality < 100 { /* Compression logic */ }
	// Example: if conversionFormat != "" { /* Conversion logic */ }
}
