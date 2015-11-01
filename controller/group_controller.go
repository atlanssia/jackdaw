package controller
import (
    "net/http"
    "github.com/atlanssia/jackdaw/utils"
    "encoding/json"
    "fmt"
    kfk "github.com/Shopify/sarama"
    "path"
    "strconv"
)

// ls /consumers
//[console-consumer-40820]

//ls /consumers/console-consumer-40820
//[ids, owners, offsets]

//ls /consumers/console-consumer-40820/offsets
//[mytopic]

//ls /consumers/console-consumer-40820/offsets/mytopic
//[0, 1, 2]

//get /consumers/console-consumer-40820/offsets/mytopic/0
//2

func (c *Controller) ListGroups(w http.ResponseWriter, r *http.Request) {

    // ls /consumers
    //[console-consumer-40820]
    gPath := "/consumers"
    groups, err := c.lsChildren(gPath)
    if err != nil {
        utils.WriteError(w, err)
        return
    }

    // map[groupName]map[topicName]map[partitionId]map[string]string (last map is offset, log size and lag etc.)
    resp := make(map[string]map[string]map[string]map[string]string, len(groups))
    for _, g := range groups {

        //ls /consumers/console-consumer-40820/offsets
        //[mytopic]
        topicsPath := path.Join(gPath, g, "offsets")
        topics, err := c.lsChildren(topicsPath)
        if err != nil {
            utils.WriteError(w, err)
            return
        }

        // init topic map
        topicMap := make(map[string]map[string]map[string]string)
        for _, topic := range topics {
            //ls /consumers/console-consumer-40820/offsets/mytopic
            //[0, 1, 2]
            partitionPath := path.Join(topicsPath, topic)
            partitions, err := c.lsChildren(partitionPath)
            if err != nil {
                utils.WriteError(w, err)
                return
            }

            // init partition map
            pMap := make(map[string]map[string]string)
            for _, pidStr := range partitions {
                //get /consumers/console-consumer-40820/offsets/mytopic/0
                //2
                offsetPath := path.Join(partitionPath, pidStr)
                offsetStr, err := c.getChildren(offsetPath)
                if err != nil {
                    utils.WriteError(w, err)
                    return
                }

                pid64, err := strconv.ParseInt(pidStr, 10, 32)
                if err != nil {
                    utils.WriteError(w, err)
                    return
                }
                pid := int32(pid64)
                offset, err := strconv.ParseInt(offsetStr, 10, 64)

                logSize := c.getLogSize(topic, pid)
                lag := logSize - offset

                pDataMap := make(map[string]string)
                pDataMap["offset"] = offsetStr
                pDataMap["logSize"] = strconv.FormatInt(logSize, 10)
                pDataMap["lag"] = strconv.FormatInt(lag, 10)
                pMap[pidStr] = pDataMap
            }

            topicMap[topic] = pMap
        }
        resp[g] = topicMap
    }
    encoder := json.NewEncoder(w)
    err = encoder.Encode(resp)
    if err != nil {
        fmt.Println(err.Error())
    }

    //    b, err := json.Marshal(topics)
    //    if err != nil {
    //        fmt.Println("error:", err)
    //    }
    //    w.Write(b)


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

func (c *Controller) getLogSize(topic string, pid int32) int64 {
    // from kafka
    conn := c.getKafkaConn()
    client, err := kfk.NewClient(conn, kfk.NewConfig())
    if err != nil {
        panic(err)
    }
    defer client.Close()

    client.Config().ClientID = "jackdaw"
    latestOffset, err := client.GetOffset(topic, pid, kfk.OffsetNewest)
    if err != nil {
        panic(err)
    }

    fmt.Printf("$$$$$$: %d\n", latestOffset)
    return latestOffset
}
