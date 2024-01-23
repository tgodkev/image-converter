package imageutils


import(
    "github.com/disintegration/imaging"
    "fmt"
    "os"
)

func CompressImage(fileName string, quality int) {

    src, err := imaging.Open(fileName)
    if err != nil {
        fmt.Printf("Failed to open image: %s\n", err)
        return
    }

    src = imaging.Resize(src, 800, 0, imaging.Lanczos)

    err = imaging.Save(src, fileName, imaging.JPEGQuality(quality))

    if err != nil {
        fmt.Printf("Failed to save image: %s\n", err)
        return
    }

    return



}


