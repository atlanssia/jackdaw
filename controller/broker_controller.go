package controller
import (
    "github.com/atlanssia/jackdaw/utils"
    "net/http"
    "encoding/json"
    "fmt"
    "strings"
)

//type broker struct {
//    jmx_port  int
//    timestamp string
//    host      string
//    version   int
//    port      int
//}

// request url: /brokers
// json from zk looks like:
//get /brokers/ids/0
//{"jmx_port":-1,"timestamp":"1446347718036","host":"U","version":1,"port":9092}
func (c *Controller) ListBrokers(w http.ResponseWriter, r *http.Request) {

    path := "/brokers/ids"
    ids, err := c.lsChildren(path)
    if err != nil {
        utils.WriteError(w, err)
        return
    }

    //map[id]broker
    resp := make(map[string]string, len(ids))
    for _, id := range ids {
        idPath := strings.Join([]string{path, id}, "/")
        json, err := c.getChildren(idPath)
        if err != nil {
            continue
        }
        resp[id] = json
    }

    encoder := json.NewEncoder(w)
    err = encoder.Encode(resp)

    fmt.Printf("%v", resp)

    //    b, err := json.Marshal(topics)
    //    if err != nil {
    //        fmt.Println("error:", err)
    //    }
    //    w.Write(b)
}

func (c *Controller) Broker(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(r.Form.Get("id")))
}