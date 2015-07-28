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

func Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    resp, err := search(conn)

    if err != nil {
        panic(err)
    }

    json.NewEncoder(w).Encode(resp)
}

func createIndex(conn *goes.Connection) (*goes.Response, error) {
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

    return conn.CreateIndex("joker", mapping)
}

func count(conn *goes.Connection) (*goes.Response, error) {
    return conn.Count("asfd", []string{"joker"}, []string{"sdf"}, url.Values{})
}

func bulkSend(conn *goes.Connection) (*goes.Response, error) {
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

    return conn.BulkSend(docs)
}

func search(conn *goes.Connection) (*goes.Response, error) {
    var query = map[string]interface{}{
        "from":   0,
        "size":   100,
    }

    return conn.Search(query, []string{"joker"}, []string{""}, url.Values{})
}

func main() {

    createIndex(conn)
    bulkSend(conn)

    router := httprouter.New()
    router.GET("/", Search)

    log.Fatal(http.ListenAndServe(":8080", router))
}
