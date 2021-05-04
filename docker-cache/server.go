package main 

import (
    "fmt"
    "github.com/Pallinder/go-randomdata"
    "net/http"
    "log"
    "encoding/json"
)

func main(){
    http.HandleFunc("/health",HealthHandler)
    http.HandleFunc("/data",ExternalHandler)
    log.Fatal(http.ListenAndServe(":8080",nil))
}


func HealthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ok")
}


func ExternalHandler(w http.ResponseWriter, r *http.Request) { 
    results := make(map[string]interface{})
    results["name"] = randomdata.SillyName()
    results["gender"] = randomdata.Male 
    results["country"] = randomdata.FullCountry

    b , err := json.Marshal(results)

    if err != nil {
        log.Println(err)
        fmt.Fprintf(w,err.Error())
        return
    }

    fmt.Fprintf(w,string(b))

}
