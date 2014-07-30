package main

import (
  "io/ioutil"
  "net/http"
  "flag"
  "net"
  "log"
  "github.com/danward79/simpleweb"
  "github.com/drone/routes"
  "strings"
  "os"
  "html/template"
  "path"
)

const (
  Template string = "./app/template/"
  Static string = "./app/static/"
  Redirect string = "/index"
  Address string = ":8080"
)

func init () {
  smplweb.SetPathConfig(&smplweb.PathConfig{Template: Template, Static: Static, Redirect: Redirect})
  
  smplweb.CreateTemplates()
  
  mux := routes.New()
  mux.Get("/", smplweb.GeneralHandler)
  mux.Get("/articles", articlesHandler)
  mux.Get("/:page", smplweb.GeneralHandler)
  mux.Get("/:page/:contents", smplweb.GeneralContentHandler)
  mux.Get("/:page/:folder/:contents", smplweb.GeneralContentHandler)

  http.Handle("/", mux)
}

//Get a folder & file list
func getFilesInfo(path string) ([]os.FileInfo, error){
  file, err := os.Open(path)
  defer file.Close()
  if err != nil {
    return nil, err
  }
  
  contents, err := file.Readdir(0)
  if err != nil {
    return nil, err
  }
  return contents, nil
}

// Add an HTML line for a file
func addFileItem(path string, name string) (string, error){
  return `<a href="/howto/` + path + `" class="list-group-item">` + name + `</a><br>`, nil
}

// Add an HTML line for a folder
func addFolderHeaderItem(name string) (string, error){
  return `<p class="lead">` + name + `</p>`, nil
}

// Build page content from the folder info passed
func contentLoop(contents []os.FileInfo, p string) (string, error){
  
  p = path.Base(p)

  var body string = ""
  
  for item := range contents {
    if contents[item].Mode().IsDir() {
      if (string(contents[item].Name()[0])) != "." {
        b, _ := addFolderHeaderItem(contents[item].Name())
        body = body + b
      
        c, err := getFilesInfo(Static + contents[item].Name() + "/")
        if err != nil {
          return "", err
        }
        b, _ = contentLoop(c, contents[item].Name() + "/")
        body = body + b
      }
      
    } else {
      name := strings.TrimSuffix(contents[item].Name(), path.Ext(contents[item].Name()))
      b, _ := addFileItem(p + "/" + name, name)
      body = body + b
    }
    
  }
  
  return body, nil
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
  page := &smplweb.Page{}

  contents, err := getFilesInfo(Static)
  if err != nil {
    return 
  }
  
  var body string = `<div class="list-group">`
  
  b, _ := contentLoop(contents, Static)
  body = body + b + `</div>`  
  
  page.Body = template.HTML(body)
  page.Title = strings.Title("articles")
  smplweb.RenderTemplate(w, "articles.tmpl", page)
}

var addr = flag.Bool("addr", false, "find open address and print to final-port.txt")

func main() {
  flag.Parse()
  
  //File route handlers
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("app/css"))))
  http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("app/js"))))
  
  if *addr {
    l, err := net.Listen("tcp", ":0")
    if err != nil {
      log.Fatal(err)
    }
    err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
    if err != nil {
      log.Fatal(err)
    }
    s := &http.Server{}
    s.Serve(l)
    return
  }
  
  http.ListenAndServe(Address, nil)
}