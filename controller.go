package main
import (
    "github.com/araframework/ara"
    kfk "github.com/Shopify/sarama"
    "net/http"
    "encoding/json"
    "fmt"
    "time"
    "github.com/samuel/go-zookeeper/zk"
)

type Controller struct {
    ara.Controller
}

func (c *Controller) ListTopics(w http.ResponseWriter, r *http.Request) {

    path := "/brokers/topics"
    topics, err := c.lsChildren(path)
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
    defer client.Close()

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


func (c *Controller)lsChildren(path string) (children []string, err error) {
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

func (c *Controller) Topic(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(r.Form.Get("id")))
}