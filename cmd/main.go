package main

import (
	"flag"
	"fmt"
	"mymodule/imageutils"
	"os"
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

		}(file.Name())
	}
	wg.Wait()
}

// ProcessImage - function for image processing
func ProcessImage(fileName string, cropSize string, quality int, conversionFormat string) error {
	fmt.Printf("Processing %s: CropSize=%d, Quality=%d, ConvertTo=%s\n", fileName, cropSize, quality, conversionFormat)

	if cropSize != "0x0" {

		err := imageutils.CropImage(fileName, cropSize)
		if err != nil {

			return fmt.Errorf("Failed to crop image: %s\n", err)
		}
	}

	if quality < 100 {
		err := imageutils.CompressImage(fileName, quality)
		if err != nil {

			return fmt.Errorf("Failed to compress image: %s\n", err)
		}

	}
	 if conversionFormat != "" {
    //create a new function in imageutils.go to convert image

        err := imageutils.ConvertImage(fileName, conversionFormat)
        if err != nil {
            return fmt.Errorf("Failed to convert image: %s\n", err)
        }


     }

	return nil
}
