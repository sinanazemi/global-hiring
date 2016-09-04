package util

import (
  //"os"
  //"bufio"
  "fmt"
  "strconv"
  "net/http"
  "database/sql"
)

const ImageService = "/image"

const ImageParam = "ImageID"

type Image struct {
  Id int
  Category string
  Type string
  ParentId int
  Buffer []byte
}

func GetImageURL(imageID int) string {
  return ImageService + "?" + ImageParam + "=" + strconv.Itoa(imageID)
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {

  imageIdStr := r.URL.Query()[ImageParam][0]
  if len(imageIdStr) == 0 {
    fmt.Fprintf(w, "No Image ID Found")
    return
  }

  imageId, err := strconv.Atoi(imageIdStr)
  if err != nil {
    fmt.Fprintln(w, err)
    return
  }

  query := "Select id, category, type, parentID, image From image where id = $1";

  imageArr, err := Select(readImage, query, imageId)
  if err != nil {
    fmt.Fprintln(w, err)
    return
  }

  image := imageArr[0].(Image)

  w.Header().Set("Content-Typee", image.Type)
  w.Write(image.Buffer)
}


func readImage(rows *sql.Rows) (interface{}, error) {
  var image Image = Image{}

  err := rows.Scan(&image.Id, &image.Category, &image.Type, &image.ParentId, &image.Buffer)
  return image, err
}


/*
func InsertAll() {

  names := []string{"customer_service", "admin_support", "sales", "accounting", "legal", "translation", "writing", "design", "engineering", "data_science", "it_network", "software_dev"}
  postfixes :=  []string{"", "_selected"}
  sizes := []string{"1", "2", "3", "4", "5"}

  parents := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
  postfixesShort := []string{"u", "s"}

  query :=
    "INSERT INTO Image " +
    "(Category, Type, ParentID, image) " +
    "VALUES($1, $2, $3, $4) " +
    "returning ID"

  for i, _ := range names {
    for j, _ := range postfixes {
      for k, _ := range sizes {
        fileName := "images\\" + names[i] + postfixes[j] + "@" + sizes[k] + "x.png"
        print(fileName)

        imgFile, err := os.Open(fileName) // a QR code image

        if err != nil {
          fmt.Println(err)
          continue
        }

        // create a new buffer base on file size
        fInfo, _ := imgFile.Stat()
        var size int64 = fInfo.Size()
        print(size)
        buf := make([]byte, size)

        // read file content into buffer
        fReader := bufio.NewReader(imgFile)
        fReader.Read(buf)

        imgFile.Close()

        _, err = Insert(query, "MS" + sizes[k] + postfixesShort[j] , "image/png", parents[i], buf)
      }
    }
  }

}
*/
