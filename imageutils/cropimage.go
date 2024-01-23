package imageutils

import (
	"fmt"
	"github.com/disintegration/imaging"
	"strconv"
	"strings"
)

func CropImage(fileName string, dimensions string) error {
	if dimensions != "0x0" {
		parts := strings.Split(dimensions, "x")

        fmt.Printf("parts: %v\n", parts)
		if len(parts) == 2 {
			heightStr, widthStr := parts[0], parts[1]

			height, err := strconv.Atoi(heightStr)
			if err != nil {
				fmt.Printf("Invalid height: %s\n", heightStr)
				return err
			}

			width, err := strconv.Atoi(widthStr)
			if err != nil {
				fmt.Printf("Invalid width: %s\n", widthStr)
				return err
			}

			img, err := imaging.Open(fileName)
			if err != nil {
				fmt.Printf("Failed to open image: %s\n", err)
				return err
			}

			resizedImg := imaging.Resize(img, width, height, imaging.Lanczos)

			outputFilePath := "resized_" + fileName
			err = imaging.Save(resizedImg, outputFilePath)
			if err != nil {
				fmt.Printf("Failed to save resized image: %s\n", err)
				return err
			}

			fmt.Printf("Image processed and saved to %s\n", outputFilePath)
		} else {
			return fmt.Errorf("Invalid crop size: %s", dimensions)
		}
	}

return nil
}

