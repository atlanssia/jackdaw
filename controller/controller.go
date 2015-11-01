package controller
import (
    "github.com/araframework/ara"
    "time"
    "github.com/atlanssia/jackdaw/utils"
    "github.com/samuel/go-zookeeper/zk"
)

type Controller struct {
    ara.Controller
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