package imageutils



import(
  "github.com/disintegration/imaging"
  "fmt"
)



func ConvertImage(fileName string, format string) ( error) {
      fmt.Println("Converting image starting now!")
    srcImg, err := imaging.Open(fileName)
    fmt.Println("Converting image to "+format)
    if err != nil{
    return fmt.Errorf("Error: %v" ,err)  
    }
    err = imaging.Save(srcImg, fileName+"."+format)
   if err != nil{
    return fmt.Errorf("Error: %v" ,err)
    }

    



    return nil
}
