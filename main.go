package main

import (
    // "fmt"
    "net/url"
    "net/http"
    // "html"
    "log"
    "encoding/json"
    "github.com/belogik/goes"
    "github.com/julienschmidt/httprouter"
)

var conn *goes.Connection = goes.NewConnection("elasticsearch", "9200")

func formatJson(w http.ResponseWriter, resp *goes.Response, err error) {
    if err != nil {
        panic(err)
    }

    json.NewEncoder(w).Encode(resp)
}

func CreateIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    mapping := map[string]interface{} {
        "settings": map[string]interface{} {
            "index.number_of_shards":   1,
            "index.number_of_replicas": 0,
        },
        "mappings": map[string]interface{} {
            "_default_": map[string]interface{} {
                "_source": map[string]interface{} {
                    "enabled": true,
                },
                "_all": map[string]interface{} {
                    "enabled": false,
                },
            },
        },
    }

    resp, err := conn.CreateIndex(ps.ByName("indexName"), mapping)
    formatJson(w, resp, err)
}

func BulkSend(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    d := goes.Document {
        Index: "joker",
        Type:  "good",
        BulkCommand: goes.BULK_COMMAND_INDEX,
        Fields: map[string]interface{}{
            "user": "james",
            "cool_level": "super",
        },
    }

    d2 := goes.Document {
        Index: "joker",
        Type:  "good",
        BulkCommand: goes.BULK_COMMAND_INDEX,
        Fields: map[string]interface{}{
            "user": "eric",
            "cool_level": "awesome",
        },
    }

    docs := []goes.Document{}
    docs = append(docs, d, d2)

    resp, err := conn.BulkSend(docs)
    formatJson(w, resp, err)
}

func Search(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    var query = map[string]interface{}{
        "from":   0,
        "size":   100,
    }

    resp, err := conn.Search(query, []string{ps.ByName("indexName")}, []string{""}, url.Values{})
    formatJson(w, resp, err)
}

func main() {

    router := httprouter.New()
    router.GET("/search/:indexName", Search)
    router.GET("/index/create/:indexName", CreateIndex)
    router.GET("/bulksend", BulkSend)

    log.Fatal(http.ListenAndServe(":8080", router))
}
