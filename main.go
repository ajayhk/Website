package main
 
import (
    "log"
    "net/http"
)
 
func main() {
    http.Handle("/", http.FileServer(http.Dir("~/workspace/goWS/src/bitbucket.org/myhomepage/static/")))
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
