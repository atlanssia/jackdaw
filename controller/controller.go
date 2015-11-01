package controller
import (
    "github.com/araframework/ara"
    "time"
    "github.com/atlanssia/jackdaw/utils"
    "github.com/samuel/go-zookeeper/zk"
    "path"
    "encoding/json"
    "fmt"
    "strconv"
)

type Controller struct {
    ara.Controller
}

// {"jmx_port":-1,"timestamp":"1446362470753","host":"U","version":1,"port":9092}
type Broker struct {
    Jmx_port  int `json:"jmx_port"`
    Timestamp string `json:"timestamp"`
    Host      string `json:"host"`
    Version   int `json:"version"`
    Port      int `json:"port"`
}

// zookeeper cmd: ls /path
func (c *Controller)lsChildren(path string) (children []string, err error) {
    zc, _, err := zk.Connect(utils.AppConf.Zookeepers, time.Second)
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

func (c *Controller) getChildren(path string) (value string, err error) {
    zc, _, err := zk.Connect(utils.AppConf.Zookeepers, time.Second)
    if err != nil {
        return
    }
    defer zc.Close()

    valueBytes, _, err := zc.Get(path)
    value = string(valueBytes)
    return
}

func (c *Controller) getKafkaConn() []string {
    brokerMap, err := c.getBrokers()
    if err != nil {
        return nil
    }

    conn := make([]string, len(brokerMap))

    for _, brokerJson := range brokerMap {

        var b Broker
        err := json.Unmarshal([]byte(brokerJson), &b)
        if err != nil {
            fmt.Println("error:", err)
            continue
        }
        conn = append(conn, b.Host + ":" + strconv.Itoa(b.Port))
    }
    return conn
}

func (c *Controller) getBrokers() (brokerMap map[string]string, err error) {
    brokersPath := "/brokers/ids"
    ids, err := c.lsChildren(brokersPath)
    if err != nil {
        return
    }

    //map[id]broker
    brokerMap = make(map[string]string, len(ids))
    for _, id := range ids {
        idPath := path.Join(brokersPath, id)
        json, err := c.getChildren(idPath)
        if err != nil {
            continue
        }
        brokerMap[id] = json
    }
    return
}