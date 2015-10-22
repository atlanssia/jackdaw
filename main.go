package main

import (
    "fmt"
    "time"

    "github.com/samuel/go-zookeeper/zk"
    kfk "github.com/Shopify/sarama"
    "net/http"
    "os"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("static")))
    http.HandleFunc("/status", status)

    bind := ":8600"
    fmt.Printf("listening on %s...", bind)
    err := http.ListenAndServe(bind, nil)
    if err != nil {
        panic(err)
    }
}

func watchData() {
    c, _, err := zk.Connect([]string{"10.180.130.177"}, time.Second) //*10)
    if err != nil {
        panic(err)
    }
    defer c.Close()

//    path := "/consumers/console-consumer-45553/offsets/ttt/0"

//    children, childrenStat, err := zc.Children(path)
//    if err != nil {
//        panic(err)
//    }
//    fmt.Printf("%+v\n %+v\n", children, childrenStat)

    // from kafka
    client, err := kfk.NewClient([]string{"10.180.130.177:9092"}, kfk.NewConfig())
    if err != nil {
        panic(err)
    }
    client.Config().ClientID = "jackdaw"
    latestOffset, err := client.GetOffset("ttt", 0, kfk.OffsetNewest)
    if err != nil {
        panic(err)
    }

    fmt.Printf("$$$$$$: %d\n", latestOffset)

//    bts, stat, ch, err := zc.GetW(path)
//    if err != nil {
//        panic(err)
//    }
//    fmt.Printf("%s *** %+v\n", string(bts), stat)

//    e := <-ch
//    fmt.Printf("--- %+v\n", e)
//    if e.Type == zk.EventNodeDataChanged {
//        watchData(zc)
//    }
}
