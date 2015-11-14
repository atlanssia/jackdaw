package utils

import (
    "os"
    "encoding/json"
    "github.com/araframework/ara"
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
        ara.Logger().Debug("error: %v", err)
    }
    ara.Logger().Debug("%v", AppConf)
}
