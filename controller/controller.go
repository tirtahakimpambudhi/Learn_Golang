package controller

import (
  "net/http"
  "ttd/views"
  "fmt"
  "time"
  "path/filepath"
  "os"
  "ttd/config"
  "io"
  "log"
  "strings"
  "github.com/gorilla/mux"
)
type DATA struct {
 Title    string
 Message  string
 Data    interface{}
}
// Implement your controllers here


func Home(w http.ResponseWriter, r *http.Request) {
  views.Render(w,"index",DATA{
    Title : "Testing",
    Message : "Template Caching",
    Data : "Home Page",
  })
}

func About(w http.ResponseWriter, r *http.Request) {
  tmpl := views.HTML{
    Main : "templates/about.gohtml",
    Header : "templates/layouts/header.gohtml",
    Footer : "templates/layouts/footer.gohtml",
  }
  
  views.RenderWithLayouts(w,tmpl,DATA{
    Title : "Siswa Kelas XI",
    Message : "Siswa",
    Data : "Test",
  })
  
  // http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func Form(w http.ResponseWriter, r *http.Request) {
  views.Render(w,"form-upload",DATA{
    Title : "File Upload",
    Message : "File Upload Testing",
    Data : "File Upload",
  })
}

func Upload(w http.ResponseWriter, r *http.Request) {
  fname := r.FormValue("fname")
  lname := r.FormValue("lname")
  
  file , fileHeader , err := r.FormFile("avatar")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  
  if err := r.ParseMultipartForm(1024); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
  
  filename := fmt.Sprintf("%v_%v_%v%v",time.Now().UnixNano(),fname,lname,filepath.Ext(fileHeader.Filename))
  
  fileLocation := filepath.Join(config.WorkDir,"static","img",filename)
  
 targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
defer targetFile.Close()

if _, err := io.Copy(targetFile, file); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

pathFile := strings.Split(fileLocation,"/")
pathFileName := pathFile[len(pathFile) - 1]
log.Printf("file '%v' success to save in direcory %v",filename,fileLocation)
  views.Render(w,"file-img",DATA{
    Title : "File Img",
    Message : fmt.Sprintf("%v %v",fname,lname),
    Data : pathFileName,
  })
}


func Download(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  image := vars["image"]
  
  pathfile := filepath.Join("static","img",image)
  w.Header().Add("Content-Disposition",fmt.Sprintf("attachment;filename=%v",image))
  http.ServeFile(w,r,pathfile)
}