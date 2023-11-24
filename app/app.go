package app

import (
  "ttd/routes"
  "ttd/config"
  "net/http"
  )

func App() {
  server := http.Server{
    Addr : config.Address ,
    Handler : routes.Routes(),
  }
  
  server.ListenAndServe()
}