package controller
import (
    "github.com/araframework/ara"
    "time"
    "github.com/atlanssia/jackdaw/utils"
    "github.com/samuel/go-zookeeper/zk"
    kfk "github.com/Shopify/sarama"
    "path"
    "encoding/json"
    "strconv"
)

type Controller struct {
    ara.Controller
    zc *zk.Conn
    kc kfk.Client
}

type appLogger struct {}

// {"jmx_port":-1,"timestamp":"1446362470753","host":"U","version":1,"port":9092}
type Broker struct {
    Jmx_port  int `json:"jmx_port"`
    Timestamp string `json:"timestamp"`
    Host      string `json:"host"`
    Version   int `json:"version"`
    Port      int `json:"port"`
}

// close zookeeper and kafka connection while shutdown
func (c *Controller) Release() {

    if c.kc != nil {
        err := c.kc.Close()
        if err != nil {
            ara.Logger().Debug("close kafka connection failed: %s", err.Error())
        }
    }

    if c.zc != nil {
        c.zc.Close()
    }
}

// init zookeeper connection
func (c *Controller) initZc() (err error) {
    ara.Logger().Debug("init zookeeper client")
    if c.zc != nil && c.zc.State() != zk.StateHasSession {
        c.zc.Close()
    }

    c.zc, _, err = zk.Connect(utils.AppConf.Zookeepers, time.Second, c.loggerOption)
    if err != nil {
        return
    }

    ara.Logger().Debug("new zookeeper client state: %v", c.zc.State())
    return
}

// a loggerOption passed to zk.Conn
func (c *Controller) loggerOption(conn *zk.Conn) {
    logger := appLogger{}
    conn.SetLogger(logger)
}

// implement interface zk.Conn.Logger
// TODO output not right: /controller.go:69: Authenticated: id=[94863294880088115 6000], timeout=%!d(MISSING)
func (l appLogger) Printf(s string, v ...interface{})  {
    ara.Logger().Debug(s, v)
}

func (c *Controller) initKc() (err error) {
    ara.Logger().Debug("init kafka client")
    conn := c.getKafkaConn()
    c.kc, err = kfk.NewClient(conn, kfk.NewConfig())
    if err != nil {
        ara.Logger().Debug(err.Error())
        return
    }

    c.kc.Config().ClientID = "jackdaw"
    ara.Logger().Debug("got a kafka client, closed: %v", c.kc.Closed())
    return
}

// zookeeper cmd: ls /path
func (c *Controller)lsChildren(path string) (children []string, err error) {
    if c.zc == nil || c.zc.State() != zk.StateHasSession {
        err = c.initZc()
        if err != nil {
            return
        }
    }

    children, _, err = c.zc.Children(path)
    if err != nil {
        return
    }
    return
}

func (c *Controller) getChildren(path string) (value string, err error) {
    if c.zc == nil || c.zc.State() != zk.StateHasSession {
        err = c.initZc()
        if err != nil {
            return
        }
    }

    valueBytes, _, err := c.zc.Get(path)
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
            ara.Logger().Debug("error: %v", err)
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