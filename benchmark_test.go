package main
import (
    "testing"
    "fmt"
    "net/http"
    "io/ioutil"
    "runtime"
)
var client *http.Client
const procs = 4

func init() {
    client = &http.Client{
        Transport: &http.Transport{
            // just what default client uses
            Proxy: http.ProxyFromEnvironment,
            // this leads to more stable numbers
            MaxIdleConnsPerHost: procs * runtime.GOMAXPROCS(0),
        },
    }
}

func BenchmarkHello(b *testing.B) {

    for i := 0; i < b.N; i++ {
        res, err := client.Get("http://127.0.0.1:8600/topics/123")
//        r, _ := http.NewRequest("GET", "http://127.0.0.1:8600/topics/123", nil)
        if err != nil {
            fmt.Println(err.Error())
        }
        defer res.Body.Close()
    }
}

func TestHttp(t *testing.T) {
    res, err := client.Get("http://127.0.0.1:8600/topics/123")
    if err != nil {
        t.Fatal("failed: " + err.Error())
    }
    defer res.Body.Close()

    b, err := ioutil.ReadAll(res.Body)
    if err != nil {
        t.Fatalf("ReadAll: %v", err)
    }
    if s := string(b); s != "123" {
        t.Fatalf("Got body: " + s)
    }
}