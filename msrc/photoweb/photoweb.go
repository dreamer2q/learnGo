
package main

import (
	"io"
	"log"
	"path"
	"net/http"
	"os"
	"io/ioutil"
	"runtime/debug"
	"html/template"
)


const (
	UPLOAD_DIR = "./uploads"
	TEMPLATE_DIR = "./public"
	ListDir = 0x0001
)

var templates = make(map[string]*template.Template)


func init(){
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName);ext != ".html"{
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:",templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		tmpl := path.Base(templateName)
		templates[tmpl] = t	
	}
}

func check(err error){
	if err != nil{
		panic(err)
	}	
}

func renderHtml(w http.ResponseWriter,tmpl string,locals map[string]interface{}){
	err := templates[tmpl].Execute(w, locals)
	check(err)
}

func uploadHandler(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET" {
		renderHtml(w, "upload.html", nil)
	}
	if r.Method == "POST" {
		f,h,err := r.FormFile("image")	
		check(err)
		filename := h.Filename
		defer f.Close()
		t,err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)
		defer t.Close()
		_, err = io.Copy(t, f)
		check(err)
		http.Redirect(w,r,"/view?id="+filename,http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter,r *http.Request){
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath);!exists {
		http.NotFound(w, r)
		return 
	}
	w.Header().Set("Content-Type", "image")
	log.Println("View id=",imagePath)
	http.ServeFile(w, r, imagePath)
}

func isExists(filepath string)bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter,r *http.Request){
	fileInfoArr,err := ioutil.ReadDir(UPLOAD_DIR)
	check(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images,fileInfo.Name())
	}
	locals["images"] = images
	renderHtml(w, "list.html", locals)
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		defer func(){
			if e,ok := recover().(error);ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("WARN: panic in %v. - %v",fn,e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w,r)
	}
}

func staticDirHandler(mux *http.ServeMux,prefix string, staticDir string,flags int){

	mux.HandleFunc(prefix, func(w http.ResponseWriter,r *http.Request){
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w,r)
				return 
			}
		}
		http.ServeFile(w,r,file)
	})

}

func main(){

	mux := http.NewServeMux() 
	staticDirHandler(mux, "/assets/", "./public", 0)
	mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(":1234", mux)

	if err != nil {
		log.Fatal("ListenAndServe:",err.Error())
	}

}
