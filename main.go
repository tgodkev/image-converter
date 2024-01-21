package main

import (
	"flag"
	"fmt"
	"github.com/containerd/containerd/v2/core/images/converter"
	"github.com/disintegration/imaging"
	"os"
	"path/filepath"
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
    flag.IntVar(&cropSize, "crop", "", "Size to crop images to ex: '300x300' (height x width) (default: '') ")

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
func ProcessImage(path string, cropSize string , quality int, conversionFormat string) {
	fmt.Printf("Processing %s: CropSize=%d, Quality=%d, ConvertTo=%s\n", path, cropSize, quality, conversionFormat)
	 if cropSize != "" {
        var height string
        var width string
        dimensions := strings.Split(cropSize, "x")

            if len(dimensions) == 2 {
                height = dimensions[0]
                width = dimensions[1]
            } else {
                fmt.Printf("Invalid crop size: %s\n", cropSize)
                return
            }


     }



	// Example: if quality < 100 { /* Compression logic */ }
	// Example: if conversionFormat != "" { /* Conversion logic */ }
}
