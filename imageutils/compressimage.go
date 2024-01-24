package imageutils

import (
	"fmt"
	"github.com/disintegration/imaging"
)

func CompressImage(fileName string, quality int) error {

	src, err := imaging.Open(fileName)

	fmt.Println("Compressing image starting now!")
	fmt.Println("Compressing image" + fileName + "with quality" + string(quality))

	if err != nil {

		return fmt.Errorf("Failed to open image: %s\n", err)
	}

	src = imaging.Resize(src, 800, 0, imaging.Lanczos)

	err = imaging.Save(src, fileName, imaging.JPEGQuality(quality))

	if err != nil {

		return fmt.Errorf("Failed to save image: %s\n", err)
	}

	fmt.Printf("Image compressed and saved to %s\n", fileName)
	return nil

}
