package main

import (
  "io/ioutil"
  "net/http"
  "flag"
  "net"
  "log"
  
  "webapp/lib/page"
  "webapp/lib/route"
  "fmt"
  "strings"
)

const (
  Address string = ":8080"
  TemplatePath string = "./app/template/"
  StaticPath string = "./app/static/"
  RedirectPath string = "/index"
)

func init () {
  page.CreateTemplates (TemplatePath)
}

func generalHandler(w http.ResponseWriter, r *http.Request, title string) {
  p := &page.Page{Title: strings.Title(title)}
  page.RenderTemplate(w, title + ".tmpl", p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  p := &page.Page{Title: "Home"}
  page.RenderTemplate(w, "index.tmpl", p)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
  p := &page.Page{Title: "About"}
  page.RenderTemplate(w, "about.tmpl", p)
}

func contactHandler(w http.ResponseWriter, r *http.Request, title string) {
  p := &page.Page{Title: title}
  page.RenderTemplate(w, title + ".tmpl", p)
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
  p := &page.Page{Title: "Articles"}
  page.RenderTemplate(w, "articles.tmpl", p)
}

func howtoHandler(w http.ResponseWriter, r *http.Request, title string) {
  p, err := page.LoadPage(title)
  if err != nil {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }
  
  page.RenderTemplate(w, "howto.tmpl", p)
}

func pageHandler(w http.ResponseWriter, r *http.Request, title string) {
  fmt.Println("PAGEHANDLER")
  p, err := page.LoadPage(title)
  if err != nil {
    http.Redirect(w, r, "/" + title, http.StatusFound)
    return
  }
  
  fmt.Println(title)
  p = &page.Page{Title: strings.Title(title)}
  fmt.Println(p)
  
  page.RenderTemplate(w, title + ".tmpl", p)
}


var addr = flag.Bool("addr", false, "find open address and print to final-port.txt")

func main() {
  flag.Parse()
  
  //File route handlers
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("app/css"))))
  http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("app/js"))))
 
  //Static route handlers - Need to use closure
  http.HandleFunc("/", route.MakeHandler(indexHandler))
  http.HandleFunc("/index/", route.MakeHandler(generalHandler))
  http.HandleFunc("/about/", route.MakeHandler(generalHandler))
  http.HandleFunc("/contact/", route.MakeHandler(generalHandler))
  http.HandleFunc("/howto/", route.MakeHandler(howtoHandler))
  http.HandleFunc("/articles/", articlesHandler)
  
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