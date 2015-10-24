package main

import (
    "fmt"
    "time"
    "github.com/samuel/go-zookeeper/zk"
    kfk "github.com/Shopify/sarama"
    "net/http"
    "github.com/araframework/ara"
    "encoding/json"
)

func init() {
    initConf()
}

func main() {
    router := ara.NewRouter()

    router.Handle("/", http.FileServer(http.Dir("static")))
    router.HandleFunc("/topics", listTopics)

    ara.Start(router)
}

func listTopics(w http.ResponseWriter, r *http.Request) {

    path := "/brokers/topics"
    topics, err := lsChildren(path)
    if err != nil {
        writeError(w, err)
        return
    }

    resp := make(map[string]string, len(topics))
    for _, topic := range topics {
        resp[topic] = "/uri/here?abc=123"
    }

    encoder := json.NewEncoder(w)
    err = encoder.Encode(resp)
//    b, err := json.Marshal(topics)
//    if err != nil {
//        fmt.Println("error:", err)
//    }
//    w.Write(b)

    // from kafka
    client, err := kfk.NewClient([]string{"127.0.0.1:9092"}, kfk.NewConfig())
    if err != nil {
        panic(err)
    }
    client.Config().ClientID = "jackdaw"
    latestOffset, err := client.GetOffset("tpk001", 0, kfk.OffsetNewest)
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

func lsChildren(path string) (children []string, err error) {
    zc, _, err := zk.Connect(appConf.Zookeepers, time.Second)
    if err != nil {
        return
    }
    defer zc.Close()

    children, _, err = zc.Children(path)
    if err != nil {
        return
    }
    return
}

//ls /brokers/ids
//[0]

//get /brokers/ids/0
//{"jmx_port":-1,"timestamp":"1445518677618","host":"1.2.3.4","version":1,"port":9092}

//ls /brokers/topics
//[ttt]
//get /brokers/topics/ttt
//{"version":1,"partitions":{"0":[0]}}
//get /brokers/topics/ttt/partitions/0/state
//{"controller_epoch":2,"leader":0,"version":1,"leader_epoch":1,"isr":[0]}
