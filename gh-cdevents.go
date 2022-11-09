package main

import (
    "encoding/json"
    "fmt"
    // "io/ioutil"
    "os"
)

func main() {
    gtihubcontext := os.Getenv("GITHUB_CONTEXT")

    var ghevent map[string]interface{}
    json.Unmarshal([]byte(gtihubcontext), &ghevent)

    fmt.Println(ghevent["event_name"])
}
