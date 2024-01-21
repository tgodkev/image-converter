package main

import (
	"flag"
	"fmt"
	"github.com/disintegration/imaging"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	dir     = flag.String("dir", ".", "Directory to process images")
	resize  = flag.String("resize", "0x0", "Target size for resizing, formatted as WIDTHxHEIGHT. Use '0x0' to skip resizing.")
	quality = flag.Int("quality", 75, "Compression quality (1-100)")
	format  = flag.String("format", "jpg", "Target format for conversion (jpg, png, gif, webp)")
)

func main() {
	flag.Parse()

	// Validate inputs...

	files, err := os.ReadDir(*dir)
	if err != nil {
		fmt.Printf("Failed to read directory: %v\n", err)
		return
	}

	var wg sync.WaitGroup
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()
			ProcessImage(fileName, *resize, *quality, *format)
			// Add your logic to handle resizing, compression, and conversion
		}(file.Name())
	}
	wg.Wait()
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
