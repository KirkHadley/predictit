
package main

import (
    "fmt"
    "predictit/go_version/data_collection"
    "encoding/json"
    "os"
)



func main() {
a := data_collection.ParseData(data_collection.GetData("439", "30"))
e, _ := json.Marshal(a)
fmt.Println(string(e))
jsonFile, err := os.Create("data.json")

if err != nil {
fmt.Println(err)
}
defer jsonFile.Close()

jsonFile.Write(e)
jsonFile.Close()

}

