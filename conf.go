package main
import (
    "os"
    "encoding/json"
    "fmt"
)

type App struct {
    Zookeepers []string `json:"zookeepers"`
}

var appConf App

func initConf() {
    file, _ := os.Open("conf/app.json")
    defer file.Close()
    
    decoder := json.NewDecoder(file)
    appConf = App{}
    err := decoder.Decode(&appConf)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(appConf)
}
