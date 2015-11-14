package main

import (
    "github.com/araframework/ara"
    "github.com/atlanssia/jackdaw/controller"
    "os"
    "os/signal"
)

func init() {
    // init conf here?

}

func main() {

    controller := &controller.Controller{}

    // catch Ctrl-c and kill signal
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, os.Interrupt, os.Kill)
    go func() {
        s := <-sc
        ara.Logger().Debug("[Jackdaw]Got signal %s, I will cleanup and exit now.\n", s)

        controller.Release()

        os.Exit(0)
    }()

    // 1. tell the framework what type my controller is
    router := ara.NewRouter()

    router.SetController(controller)

//    router.Handle("/", http.FileServer(http.Dir("static")))
//    router.HandleFunc("/topics", listTopics)

    ara.Start(router)
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
