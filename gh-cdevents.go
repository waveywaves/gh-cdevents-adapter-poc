package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    gheventFile, err := os.Open("event.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()

    b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
        fmt.Println(err)
    }

    var ghevent map[string]interface{}
    json.Unmarshal([]byte(b), &result)

    fmt.Println(ghevent["event_name"])
}
