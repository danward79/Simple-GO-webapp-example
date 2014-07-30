package route

import (
  "net/http"
  "regexp"
  "fmt"
)

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
  fmt.Println("MakeHandler")
  return func(w http.ResponseWriter, r *http.Request) {
    m := GetPath(r.URL.Path)
    if m == nil {
      fmt.Println("MakeHandler: m = ", m)
      http.Redirect(w, r, "/", http.StatusFound)
      return
    }
    fn(w, r, m[2])
  }
}

//"^/(about|contact|index|howto)/([a-zA-Z0-9]+)$"
var validPath = regexp.MustCompile("^/(about|contact|index|howto)/([a-zA-Z0-9]+)$")

func GetPath (path string) []string{
  return validPath.FindStringSubmatch(path)
}

