package app

import (
  "ttd/routes"
  "ttd/config"
  "net/http"
  _"fmt"
  "log"
  )
type Logging struct {
  Handler http.Handler
}

func (middleware *Logging) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 // fmt.Printf("Method : %v \n Path : %v \n Proto : %v \n",r.Method,r.URL.Path,r.Proto)
 log.Printf("%v %v \n",r.Method,r.URL.Path)
  middleware.Handler.ServeHTTP(w,r)
} 
func App() {
  logging := &Logging{
    Handler : routes.Routes(),
  }
  server := http.Server{
    Addr : config.Address ,
    Handler : logging,
  }
  
  server.ListenAndServe()
}