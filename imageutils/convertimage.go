package imageutils



import(
  "github.com/disintegration/imaging"
  "fmt"
)



func ConvertImage(fileName string, format string) ( error) {

    srcImg, err := imaging.Open(fileName)
    if err != nil{
    return fmt.Errorf("Error: %v" ,err)  
    }
    err = imaging.Save(srcImg, fileName+"."+format)
   if err != nil{
    return fmt.Errorf("Error: %v" ,err)
    }

    



    return nil
}
