package utils

import (
    "os"
    "encoding/json"
    "fmt"
)

type App struct {
    Zookeepers []string `json:"zookeepers"`
}

var AppConf App

func init() {
    file, _ := os.Open("conf/app.json")
    defer file.Close()
    
    decoder := json.NewDecoder(file)
    AppConf = App{}
    err := decoder.Decode(&AppConf)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(AppConf)
}
