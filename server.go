package main

import (
    "log"
    "net/http"
)

func index (w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, "index.html")
}

func authenticate (w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, "authenticate.html")
}

func provision (w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, "provision.html")
}

func reserve (w http.ResponseWriter, req *http.Request) {

}

func generate (w http.ResponseWriter, req *http.Request) {

}

func browserid (w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    http.ServeFile(w, req, "config.json")
}

func main () {
    
    http.HandleFunc("/", index)
    http.HandleFunc("/authenticate", authenticate)
    http.HandleFunc("/provision", provision)
    http.HandleFunc("/reserve", reserve)
    http.HandleFunc("/generate", generate)
    http.HandleFunc("/.well-known/browserid", browserid)
    err := http.ListenAndServeTLS(":443","concat.pem","ssl.key",nil)
    if err != nil {
        log.Fatal(err)
    }

}
