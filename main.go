/*================================================================
*   FileName    : main.go
*   Author      : mavinyeh
*   Date        : 2019-05-09
*   Description : 
================================================================*/
package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hi there, I Love %s!", r.URL.Path[1:])

}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe (":8080", nil))
}

